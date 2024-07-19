package handler

import (
	"log/slog"
	
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"
	"github.com/gin-gonic/gin"
)

// AddComment handles adding a new comment.
// @Summary Add a new comment
// @Description Add a new comment with provided details
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddComment body pb.AddCommentRequest true "Add comment request"
// @Success 200 {object} pb.AddCommentResponse "Comment successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /comment/add [post]
func (h *Handler) AddComment(ctx *gin.Context) {
	req := &pb.AddCommentRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.MemoryService.GetMemory(ctx, &pb.GetMemoryRequest{Id: req.MemoryId})
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}
	
	if res == nil  {
		ctx.JSON(404, gin.H{"error": "Memory not found"})
		return
	}

	_, err = h.CommentService.AddComment(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success create comment")
}


// GetByMemoryId handles fetching comments by memory ID.
// @Summary Get comments by memory ID
// @Description Retrieves comments associated with a specific memory ID
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Memory ID"
// @Success 200 {object} pb.GetByIdMemoryResponse "Comments found"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Comments not found"
// @Failure 500 {string} string "Internal server error"
// @Router /comment/get/memory/{id} [get]
func (h *Handler) GetByMemoryId(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetByIdMemoryRequest{MemoryId: id}

	res, err := h.CommentService.GetByMemoryId(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// UpdateComment handles updating a comment.
// @Summary Update a comment
// @Description Updates an existing comment with new content
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UpdateComment body pb.UpdateCommentRequest true "Update comment request"
// @Success 200 {object} pb.UpdateCommentResponse "Comment successfully updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /comment/update [put]
func (h *Handler) UpdateComment(ctx *gin.Context) {
	req := &pb.UpdateCommentRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	
	res, err := h.MemoryService.GetMemory(ctx, &pb.GetMemoryRequest{Id: req.MemoryId})
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}
	
	if res == nil  {
		ctx.JSON(404, gin.H{"error": "Memory not found"})
		return
	}
	_, err = h.CommentService.UpdateComment(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success update comment")
}

// DeleteComment handles deleting a comment.
// @Summary Delete a comment
// @Description Deletes a comment based on the provided ID
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param CommentID path string true "Comment ID"
// @Success 200 {object} pb.DeleteCommentResponse "Comment successfully deleted"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /comment/delete/{id} [delete]
func (h *Handler) DeleteComment(ctx *gin.Context) {
	req := &pb.DeleteCommentRequest{Id: ctx.Param("id")}

	_, err := h.CommentService.DeleteComment(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success delete comment")
}

// GetById handles fetching a comment by ID.
// @Summary Get a comment by ID
// @Description Retrieves a comment entry based on the provided ID
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Success 200 {object} pb.GetByCommentIdResponse "Comment found"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Comment not found"
// @Failure 500 {string} string "Internal server error"
// @Router /comment/get/{id} [get]
func (h *Handler) GetById(ctx *gin.Context) {
	req := &pb.GetByCommentIdRequest{Id: ctx.Param("id")}

	
	res, err := h.CommentService.GetById(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// GetAllComments handles fetching all comments.
// @Summary Get all comments
// @Description Retrieves all stored comments
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UserID query string false "User ID"
// @Param Limit query string false "Limit"
// @Param Offset query string false "Offset"
// @Success 200 {object} pb.GetCommentsResponse "List of comments"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /comment/gets [get]
func (h *Handler) GetAllComments(ctx *gin.Context) {
	req := &pb.GetCommentsRequest{
		UserId: ctx.Query("UserID"),
		Limit:  ctx.Query("Limit"),
		Ofset:  ctx.Query("Offset"),
	}

	
	res, err := h.CommentService.GetAllCommets(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}
