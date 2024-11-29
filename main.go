package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ahallhognason/MaskedSanta/handlers"
)

func main() {
	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", handlers.ShowIndex)
	r.POST("/room/create", handlers.CreateRoom)
	r.GET("/room/:roomID", handlers.ShowRoom)
	r.POST("/room/:roomID/add-participant", handlers.AddParticipant)
	r.GET("/room/:roomID/assign", handlers.AssignGifts)
	r.GET("/participant/:urlID", handlers.ViewAssignment)

	// Start server
	r.Run(":8080")
}
