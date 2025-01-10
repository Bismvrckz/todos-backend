package routes

import (
	"task-api/config"
	"task-api/handler"
)

func BuildRoutes(ein *config.Apps) {
	// Admin
	apiHandler := ein.Tkbai.Group(config.AppPrefix, handler.AppMiddleware)

	apiHandler.POST("/task/create", handler.CreateTask)
	//apiHandler.GET("/add/csv", handler.AdminInputView)
	//apiHandler.POST("/add/csv", handler.AdminAddStudentBulk)
	//apiHandler.POST("/add/student", handler.AdminAddStudent)
	//apiHandler.GET("/login", handler.AdminLoginView)
	//apiHandler.GET("/logout", handler.AdminLogout)
	//apiHandler.POST("/login", handler.AdminLogin)
	//apiHandler.POST("/delete/student", handler.AdminDeleteStudent)
}
