package handler

import (
	"log/slog"
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"

	"github.com/gin-gonic/gin"
)

// AddTimeLine handles adding a new timeline entry.
// @Summary Add a new timeline entry
// @Description Add a new timeline entry with provided details
// @Tags Timeline
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AddTimeLine body pb.AddTimeLineRequest true "Add timeline request"
// @Success 200 {object} pb.AddTimeLineResponse "Timeline entry successfully added"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /timeline/add [post]
func (h *Handler) AddTimeLine(ctx *gin.Context) {
	req := &pb.AddTimeLineRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	_, err = h.TimeLineService.AddTimeLine(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Created Timeline")
}

// GetUserEvents handles fetching user events based on filters.
// @Summary Get user events
// @Description Retrieve user events based on provided filters
// @Tags Timeline
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query string false "Limit number of events"
// @Param offset query string false "Offset for pagination"
// @Param date query string false "Date filter"
// @Success 200 {object} pb.GetUserEventsResponse "User events found"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /timeline/gets [get]
func (h *Handler) GetEvents(ctx *gin.Context) {
	req := &pb.GetUserEventsRequest{
		Limit:  ctx.Query("limit"),
		Ofset: ctx.Query("offset"),
		Date:   ctx.Query("date"),
	}

	res, err := h.TimeLineService.GetEvent(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// SearchTimeLine handles searching timeline events within a date range.
// @Summary Search timeline events
// @Description Search timeline events within a specified date range
// @Tags Timeline
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param start_date query string false "Start date of search"
// @Param end_date query string false "End date of search"
// @Success 200 {object} pb.SearchTimeLineResponse "Timeline events found"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /timeline/search [get]
func (h *Handler) SearchTimeLine(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	if startDate == "" || endDate == "" {
		ctx.JSON(400, gin.H{"error": "Start date and end date are required"})
		return
	}

	req := &pb.SearchTimeLineRequest{
		StartDate: startDate,
		EndDate:   endDate,
	}

	res, err := h.TimeLineService.SearchTimeLine(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, res)
}

// DeleteTimeLine handles deleting a timeline entry by its ID.
// @Summary Delete a timeline entry
// @Description Delete a timeline entry based on the provided ID
// @Tags Timeline
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Timeline entry ID"
// @Success 200 {object} pb.DeleteTimeLineResponse "Timeline entry successfully deleted"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Timeline entry not found"
// @Failure 500 {string} string "Internal server error"
// @Router /timeline/delete/{id} [delete]
func (h *Handler) DeleteTimeLine(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteTimeLineRequest{Id: id}

	_, err := h.TimeLineService.DeleteTimeLine(ctx, req)
	if err != nil {
		statusCode := 500
		if err.Error() == "not found" {
			statusCode = 404
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		slog.Info(err.Error())
		return
	}

	ctx.JSON(200, "Success Delete Timeline")
}
