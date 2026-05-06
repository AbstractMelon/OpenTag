package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// API Routes
	api := r.Group("/api")
	{
		api.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"system": "The OpenTag API is running",
			})
		})

		api.POST("/connect", func(c *gin.Context) {
			var req struct {
				ClientName string `json:"client_name" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}

			log.Printf("Client connected: %s", req.ClientName)
			c.JSON(http.StatusOK, gin.H{"message": "Welcome, " + req.ClientName})
		})

		api.POST("/hit", func(c *gin.Context) {
			var req struct {
				ClientName string `json:"client_name" binding:"required"`
				TargetID   string `json:"target_id" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}

			log.Printf("Hit registered: %s hit %s", req.ClientName, req.TargetID)
			c.JSON(http.StatusOK, gin.H{
				"message": "Hit recorded",
				"client":  req.ClientName,
				"target":  req.TargetID,
			})
		})

		api.POST("/health", func(c *gin.Context) {
			var req struct {
				ClientName string `json:"client_name" binding:"required"`
				Health     int    `json:"health" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}

			log.Printf("Health update: %s = %d", req.ClientName, req.Health)
			c.JSON(http.StatusOK, gin.H{"message": "Health updated"})
		})

		api.POST("/score", func(c *gin.Context) {
			var req struct {
				ClientName string `json:"client_name" binding:"required"`
				Score      int    `json:"score" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}

			log.Printf("Score update: %s = %d", req.ClientName, req.Score)
			c.JSON(http.StatusOK, gin.H{"message": "Score updated"})
		})

		api.POST("/event", func(c *gin.Context) {
			var req struct {
				ClientName string `json:"client_name" binding:"required"`
				Event      string `json:"event" binding:"required"`
				Data       string `json:"data"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}

			log.Printf("Event: %s - %s: %s", req.ClientName, req.Event, req.Data)
			c.JSON(http.StatusOK, gin.H{"message": "Event recorded"})
		})

		api.POST("/heartbeat", func(c *gin.Context) {
			var req struct {
				ClientName string `json:"client_name" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}

			log.Printf("Heartbeat: %s", req.ClientName)
			c.JSON(http.StatusOK, gin.H{"status": "alive"})
		})
	}

	// Serve Svelte frontend if build folder exists
	if _, err := os.Stat("../web/build"); !os.IsNotExist(err) {
		r.StaticFS("/_app", http.Dir("../web/build/_app"))

		// Fallback for SPA routing
		r.NoRoute(func(c *gin.Context) {
			c.File("../web/build/index.html")
		})
	} else {
		log.Println("Frontend build directory '../web/build' not found. API mode only.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed:", err)
	}
}
