package api

import (
	"github.com/Exam4/4th-month-exam-Api-Gateway/api/handler"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description API documentation for Time Capsule services
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	router := r.Group("/")
	

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Use(middleware.MiddleWare())
	media := router.Group("/media")
	{
		media.POST("/add", h.AddMedia)
		media.GET("/get/:id", h.GetMediaByMemoryId)
		media.DELETE("/delete/:id", h.DeleteMedia)
		media.PUT("/update", h.UpdateMedia)
		media.GET("/gets", h.GetAllMedia)
	}

	memory := router.Group("/memory")
	{
		memory.POST("/add", h.AddMemory)
		memory.GET("/get/:id", h.GetMemory)
		memory.GET("/gets", h.GetAllMemories)
		memory.PUT("/update", h.UpdateMemory)
		memory.DELETE("/delete/:id", h.DeleteMemory)
		memory.POST("/share", h.ShareMemory)
	}

	customEvents := router.Group("/customevent")
	{
		customEvents.POST("/add", h.AddCustomEvent)
		customEvents.PUT("/update", h.UpdateCustomEvent)
		customEvents.DELETE("/delete/:id", h.DeleteCustomEvent)
		customEvents.GET("/gets", h.GetAllCustomEvents)
		customEvents.GET("/get/:id", h.GetByIdCustomEvent)
	}

	historical := router.Group("/historical")
	{
		historical.GET("/gets", h.GetAllHistoricalEvents)
		historical.POST("/add", h.AddHistoricalEvent)
		historical.DELETE("/delete/:id", h.DeleteHistoricalEvent)
	}

	milestone := router.Group("/milestone")
	{
		milestone.POST("/add", h.AddMilestone)
		milestone.GET("/gets", h.GetAllMilestones)
		milestone.PUT("/update", h.UpdateMilestone)
		milestone.DELETE("/delete/:id", h.DeleteMilestone)
	}

	timeline := router.Group("/timeline")
	{
		timeline.POST("/add", h.AddTimeLine)
		timeline.GET("/gets", h.GetEvents)			// Id si qaytyapti
		timeline.GET("/search", h.SearchTimeLine)			// bosh qaytyapti
		timeline.DELETE("/delete/:id", h.DeleteTimeLine) 
	}

	comment := router.Group("/comment")
	{
		comment.POST("/add", h.AddComment)
		comment.GET("/get/memory/:id", h.GetByMemoryId)
		comment.PUT("/update", h.UpdateComment)
		comment.DELETE("/delete/:id", h.DeleteComment)
		comment.GET("/get/:id", h.GetById)
		comment.GET("/gets", h.GetAllComments)
	}

	return r
}
