package main

func (app *App) registerRoutes() {
	api := app.router.Group("/api")

	// Health Check
	api.GET("/health", handleHealthCheck)

	// Dyanmic CRUD
	api.GET("/:entity", handleGet)           // Get all
	api.GET("/:entity/:id", handleGetByID)   // Get one
	api.POST("/:entity", handlePost)         // Create
	api.PUT("/:entity/:id", handlePut)       // Update
	api.DELETE("/:entity/:id", handleDelete) // Delete
}
