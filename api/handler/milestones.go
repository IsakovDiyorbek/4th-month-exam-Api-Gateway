package handler

import (
	"log/slog"
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"

	"github.com/gin-gonic/gin"
)

// AddMilestone handles adding a new milestone.
// @Summary Add a new milestone
// @Description Add a new milestone with provided details
// @Tags Milestones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddMilestone body pb.AddMilestonesRequest true "Add milestone request"
// @Success 200 {object} pb.AddMilestonesResponse "Milestone successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /milestone/add [post]
func (h *Handler) AddMilestone(ctx *gin.Context) {
	req := &pb.AddMilestonesRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	_, err = h.MilestoneService.AddMilestone(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Created Milestone")
}

// GetAllMilestones handles fetching all milestones based on filters.
// @Summary Get all milestones
// @Description Retrieve all milestones based on provided filters
// @Tags Milestones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param date query string false "Date filter"
// @Param title query string false "Title filter"
// @Param category query string false "Category filter"
// @Param user_id query string false "User ID filter"
// @Success 200 {object} pb.GetAllMilestonesResponse "Milestones found"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /milestone/gets [get]
func (h *Handler) GetAllMilestones(ctx *gin.Context) {
	req := &pb.GetAllMilestonesRequest{
		Date:     ctx.Query("date"),
		Title:    ctx.Query("title"),
		Category: ctx.Query("category"),
		UserId:   ctx.Query("user_id"),
	}

	// Call your service method to get all milestones
	res, err := h.MilestoneService.GetAllMilestone(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// UpdateMilestone handles updating an existing milestone.
// @Summary Update a milestone
// @Description Update an existing milestone based on provided details
// @Tags Milestones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UpdateMilestone body pb.UpdateMilestonesRequest true "Update milestone request"
// @Success 200 {object} pb.UpdateMilestonesResponse "Milestone successfully updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Milestone not found"
// @Failure 500 {string} string "Internal server error"
// @Router /milestone/update [put]
func (h *Handler) UpdateMilestone(ctx *gin.Context) {
	req := &pb.UpdateMilestonesRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	
	_, err = h.MilestoneService.UpdateMilestone(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success updated milestone")
}

// DeleteMilestone handles deleting a milestone by its ID.
// @Summary Delete a milestone
// @Description Delete a milestone based on the provided ID
// @Tags Milestones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Milestone ID"
// @Success 200 {object} pb.DeleteMilestonesResponse "Milestone successfully deleted"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Milestone not found"
// @Failure 500 {string} string "Internal server error"
// @Router /milestone/delete/{id} [delete]
func (h *Handler) DeleteMilestone(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteMilestonesRequest{Id: id}

	_, err := h.MilestoneService.DeleteMilestone(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Delete Milestone")
}
