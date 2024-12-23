package main

import (
	"github.com/ahallhognason/MaskedSanta/handlers"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

func main() {
	r := gin.Default()

	funcMap := template.FuncMap{
		"joinStrings": strings.Join,
	}

	// Load templates with the custom FuncMap
	r.SetFuncMap(funcMap)
	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", handlers.ShowIndex)
	r.POST("/room/create", handlers.CreateRoom)
	r.GET("/room/:roomID", handlers.ShowRoom)
	r.POST("/room/:roomID/add-participant", handlers.AddParticipant)
	r.DELETE("/room/:roomID/participant/:urlID", handlers.RemoveParticipant)
	r.GET("/room/:roomID/assign", handlers.AssignGifts)
	r.GET("/participant/:urlID", handlers.ViewAssignment)
	//TO DO: Guard against changing participants after an assignment call is made

	// Start server
	r.Run(":8080")
}
