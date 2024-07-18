package handler

import (
	"log/slog"
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"

	"github.com/gin-gonic/gin"
)

// GetAllHistoricalEvents handles fetching all historical events based on filters.
// @Summary Get all historical events
// @Description Retrieve all historical events based on provided filters
// @Tags HistoricalEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query string false "Limit number of events"
// @Param offset query string false "Offset for pagination"
// @Param date query string false "Date filter"
// @Success 200 {object} pb.GetAllHistoricalResponse "Historical events found"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /historical/gets [get]
func (h *Handler) GetAllHistoricalEvents(ctx *gin.Context) {
	req := &pb.GetAllHistoricalRequest{
		Limit:  ctx.Query("limit"),
		Ofset: ctx.Query("offset"),
		Date:   ctx.Query("date"),
	}

	res, err := h.HistoricalService.GetAllHistoricalEvents(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// AddHistoricalEvent handles adding a new historical event.
// @Summary Add a new historical event
// @Description Add a new historical event with provided details
// @Tags HistoricalEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddHistoricalEvent body pb.AddHistoricalEventRequest true "Add historical event request"
// @Success 200 {object} pb.AddHistoricalEventResponse "Historical event successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /historical/add [post]
func (h *Handler) AddHistoricalEvent(ctx *gin.Context) {
	req := &pb.AddHistoricalEventRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	_, err = h.HistoricalService.AddHistoricalEvent(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Created Historical Event")
}

// DeleteHistoricalEvent handles deleting a historical event by its ID.
// @Summary Delete a historical event
// @Description Delete a historical event based on the provided ID
// @Tags HistoricalEvents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Historical event ID"
// @Success 200 {object} pb.DeleteHistoricalEventResponse "Historical event successfully deleted"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Historical event not found"
// @Failure 500 {string} string "Internal server error"
// @Router /historical/delete/{id} [delete]
func (h *Handler) DeleteHistoricalEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteHistoricalEventRequest{Id: id}

	// Call your service method to delete the historical event
	_, err := h.HistoricalService.DeleteHistoricalEvent(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Delete Historical Event")
}
