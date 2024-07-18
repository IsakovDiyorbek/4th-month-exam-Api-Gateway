package handler

import (
	"log/slog"

	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"
	"github.com/gin-gonic/gin"
)

// AddCustomEvent handles adding a new custom event.
// @Summary Add a new custom event
// @Description Add a new custom event with provided details
// @Tags CustomEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddCustomEvent body pb.AddCustomEventRequest true "Add custom event request"
// @Success 200 {object} pb.AddCustomEventResponse "Custom event successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /customevent/add [post]
func (h *Handler) AddCustomEvent(ctx *gin.Context) {
	req := &pb.AddCustomEventRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	_, err = h.CustomEventsService.AddCustomEvent(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Created Custom Event")
}

// UpdateCustomEvent handles updating an existing custom event.
// @Summary Update a custom event
// @Description Update an existing custom event based on provided details
// @Tags CustomEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UpdateCustomEvent body pb.UpdateCustomEventsRequest true "Update custom event request"
// @Success 200 {object} pb.UpdateCustomEventsResponse "Custom event successfully updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Custom event not found"
// @Failure 500 {string} string "Internal server error"
// @Router /customevent/update [put]
func (h *Handler) UpdateCustomEvent(ctx *gin.Context) {
	req := &pb.UpdateCustomEventsRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	// Call your service method to update the custom event
	_, err = h.CustomEventsService.UpdateCustomEvent(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Update Custom Event")
}
// @Security BearerAuth
// DeleteCustomEvent handles deleting a custom event by its ID.
// @Summary Delete a custom event
// @Description Delete a custom event based on the provided ID
// @Tags CustomEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Custom event ID"
// @Success 200 {object} pb.DeleteCustomEventsResponse "Custom event successfully deleted"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Custom event not found"
// @Failure 500 {string} string "Internal server error"
// @Router /customevent/{id} [delete]
func (h *Handler) DeleteCustomEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteCustomEventsRequest{EventId: id}

	_, err := h.CustomEventsService.DeleteCustomEvent(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Delete Custom Event")
}

// GetAllCustomEvents handles fetching all custom events based on filters.
// @Summary Get all custom events
// @Description Retrieve all custom events based on provided filters
// @Tags CustomEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "User ID filter"
// @Param date query string false "Date filter"
// @Param title query string false "Title filter"
// @Success 200 {object} pb.GetAllEventsResponse "Custom events found"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /customevent/gets [get]
func (h *Handler) GetAllCustomEvents(ctx *gin.Context) {
	req := &pb.GetAllEventsRequest{
		UserId: ctx.Query("user_id"),
		Date:   ctx.Query("date"),
		Title:  ctx.Query("title"),
	}


	res, err := h.CustomEventsService.GetAllCustomEvents(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// GetByIdCustomEvent handles fetching a custom event by its ID.
// @Summary Get a custom event by ID
// @Description Retrieve a custom event based on the provided ID
// @Tags CustomEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Custom event ID"
// @Success 200 {object} pb.GetByIdEvetsResponse "Custom event found"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Custom event not found"
// @Failure 500 {string} string "Internal server error"
// @Router /customevent/get/{id} [get]
func (h *Handler) GetByIdCustomEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetByIdEvetsRequest{Id: id}

	res, err := h.CustomEventsService.GetByIdCustomEvent(ctx, req)
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
