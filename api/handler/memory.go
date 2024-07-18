package handler

import (
	"log/slog"
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"

	"github.com/gin-gonic/gin"
)

// AddMemory handles adding a new memory.
// @Summary Add a new memory
// @Description Add a new memory with provided details
// @Tags Memory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddMemory body pb.AddMemoryRequest true "Add memory request"
// @Success 200 {object} pb.AddMemoryResponse "Memory successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /memory/add [post]
func (h *Handler) AddMemory(ctx *gin.Context) {
	req := &pb.AddMemoryRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	_, err = h.MemoryService.AddMemory(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Created Memory")
}

// GetMemory handles fetching a memory by its ID.
// @Summary Get a memory by ID
// @Description Retrieve a memory based on the provided ID
// @Tags Memory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Memory ID"
// @Success 200 {object} pb.GetMemoryResponse "Memory found"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Memory not found"
// @Failure 500 {string} string "Internal server error"
// @Router /memory/get/{id} [get]
func (h *Handler) GetMemory(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetMemoryRequest{Id: id}


	res, err := h.MemoryService.GetMemory(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// GetAllMemories handles fetching all memories based on filters.
// @Summary Get all memories
// @Description Retrieve all memories based on provided filters
// @Tags Memory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query string false "Page number"
// @Param limit query string false "Number of items per page"
// @Param start_date query string false "Start date filter"
// @Param end_date query string false "End date filter"
// @Param tags query string false "Tags filter"
// @Success 200 {object} pb.GetAllMemoriesResponse "Memories found"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /memory/gets [get]
func (h *Handler) GetAllMemories(ctx *gin.Context) {
	req := &pb.GetAllMemoriesRequest{
		Page:      ctx.Query("page"),
		Limit:     ctx.Query("limit"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		Tags:      ctx.Query("tags"),
	}

	// Call your service method to get all memories
	res, err := h.MemoryService.GetAllMemories(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// UpdateMemory handles updating an existing memory.
// @Summary Update a memory
// @Description Update an existing memory based on provided details
// @Tags Memory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UpdateMemory body pb.UpdateMemoryRequest true "Update memory request"
// @Success 200 {object} pb.UpdateMemoryResponse "Memory successfully updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Memory not found"
// @Failure 500 {string} string "Internal server error"
// @Router /memory/update [put]
func (h *Handler) UpdateMemory(ctx *gin.Context) {
	req := &pb.UpdateMemoryRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	// Call your service method to update the memory
	_, err = h.MemoryService.UpdateMemory(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success updated memory")
}

// DeleteMemory handles deleting a memory by its ID.
// @Summary Delete a memory
// @Description Delete a memory based on the provided ID
// @Tags Memory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Memory ID"
// @Success 200 {object} pb.DeleteMemoryResponse "Memory successfully deleted"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Memory not found"
// @Failure 500 {string} string "Internal server error"
// @Router /memory/delete/{id} [delete]
func (h *Handler) DeleteMemory(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteMemoryRequest{Id: id}

	// Call your service method to delete the memory
	res, err := h.MemoryService.DeleteMemory(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// ShareMemory handles sharing a memory with other users.
// @Summary Share a memory
// @Description Share a memory with specified users
// @Tags Memory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ShareMemory body pb.ShareMemoryRequest true "Share memory request"
// @Success 200 {object} pb.ShareMemoryResponse "Memory successfully shared"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Memory not found"
// @Failure 500 {string} string "Internal server error"
// @Router /memory/share [post]
func (h *Handler) ShareMemory(ctx *gin.Context) {
	req := &pb.ShareMemoryRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	_, err = h.MemoryService.ShareMemory(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success add tags")
}
