package main

func (app *App) registerRoutes() {
	// Health check
	app.router.GET("/", handleHealthCheck)

	// API routes
	api := app.router.Group("/api")

	// Dyanmic CRUD
	api.GET("/:entity", handleGet)           // Get all
	api.GET("/:entity/:id", handleGetByID)   // Get one
	api.POST("/:entity", handlePost)         // Create
	api.PUT("/:entity/:id", handlePut)       // Update
	api.DELETE("/:entity/:id", handleDelete) // Delete
	api.POST("/:entity/reset", handleReset)  // Reset
}
