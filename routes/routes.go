package routes

import (
	"task-api/config"
	"task-api/handler"
)

func BuildRoutes(ein *config.Apps) {
	// Admin
	apiHandler := ein.Tkbai.Group(config.AppPrefix, handler.AppMiddleware)

	apiHandler.POST("/task", handler.CreateTask)
	apiHandler.GET("/task", handler.GetAllTask)
	apiHandler.PUT("/task", handler.UpdateTask)
	apiHandler.DELETE("/task", handler.DeleteTask)
}
