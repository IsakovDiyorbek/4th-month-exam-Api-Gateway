package handler

import (
	pb "github.com/Exam4/4th-month-exam-Api-Gateway/genproto"
)

type Handler struct {
	CommentService pb.CommentServiceClient
    MediaService   pb.MediaServiceClient
    MemoryService pb.MemoryServiceClient
    CustomEventsService pb.CustomEventsServiceClient
    HistoricalService  pb.HistoricalServiceClient
    MilestoneService   pb.MilestoneServiceClient
    TimeLineService    pb.TimeLineServiceClient

}

func NewHandler(comment pb.CommentServiceClient, media pb.MediaServiceClient, 
    memory pb.MemoryServiceClient, custum pb.CustomEventsServiceClient,
    histori pb.HistoricalServiceClient, mile pb.MilestoneServiceClient,
    time pb.TimeLineServiceClient) *Handler {
	return &Handler{
		CommentService: comment,
        MediaService: media,
        MemoryService: memory,
        CustomEventsService: custum,
        HistoricalService: histori,
        MilestoneService: mile,
        TimeLineService: time,
	}
}
