package handler

import (
	"log/slog"
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"

	"github.com/gin-gonic/gin"
)

// AddMedia handles adding new media.
// @Summary Add new media
// @Description Add media to a specific memory
// @Tags Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddMedia body pb.AddMediaRequest true "Add media request"
// @Success 200 {object} pb.AddMediaResponse "Media successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /media/add [post]
func (h *Handler) AddMedia(ctx *gin.Context) {
	req := &pb.AddMediaRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	_, err := h.MediaService.AddMedia(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, "Success add media")
}

// GetMediaByMemoryId retrieves media associated with a memory ID.
// @Summary Get media by memory ID
// @Description Retrieve all media related to a specific memory ID
// @Tags Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Get media by memory ID request"
// @Success 200 {object} pb.GetMediaByMemoryIdResponse "Media retrieved successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /media/get/{id} [get]
func (h *Handler) GetMediaByMemoryId(ctx *gin.Context) {
	req := &pb.GetMediaByMemoryIdRequest{MemoryId: ctx.Param("id")}

	res, err := h.MediaService.GetMediaByMemoryId(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// DeleteMedia handles deleting media.
// @Summary Delete media
// @Description Delete media by memory ID and media ID
// @Tags Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param DeleteMedia body pb.DeleteMediaRequest true "Delete media request"
// @Success 200 {object} pb.DeleteMediaResponse "Media deleted successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /media/delete/{id} [post]
func (h *Handler) DeleteMedia(ctx *gin.Context) {
	req := &pb.DeleteMediaRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	_, err := h.MediaService.DeleteMedia(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, "Success delete media")
}

// UpdateMedia handles updating media information.
// @Summary Update media
// @Description Update media details such as type and URL
// @Tags Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UpdateMedia body pb.UpdateMediaRequest true "Update media request"
// @Success 200 {object} pb.UpdateMediaResponse "Media updated successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /media/update [put]
func (h *Handler) UpdateMedia(ctx *gin.Context) {
	req := &pb.UpdateMediaRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	_, err := h.MediaService.UpdateMedia(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, "Success update media")
}

// GetAllMedia retrieves all media based on type and URL filters.
// @Summary Get all media
// @Description Retrieve all media, optionally filtered by type and URL
// @Tags Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param type query string false "Type"
// @Param url query string false "URL"
// @Success 200 {object} pb.GetAllMediaResponse "Media retrieved successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /media/gets [get]
func (h *Handler) GetAllMedia(ctx *gin.Context) {
	req := &pb.GetAllMediaRequest{}
	req.Type = ctx.Query("type")
	req.Url = ctx.Query("url")

	res, err := h.MediaService.GetAllMedia(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, res)
}
