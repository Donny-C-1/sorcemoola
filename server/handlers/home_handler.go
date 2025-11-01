package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>SorceMoola Web Service</title>
		</head>
		<body>
			<h1>This is the SorceMoola Web Service</h1>
			<p>A private service for Sorcemoola crowdfunding application. Unauthorized access is prohibited.</p>
			<a href="https://sorcemoola.vercel.app">Visit SorceMoola Main Page</a>
			<img src="https://res.cloudinary.com/dgmbzqk2p/image/upload/v1757713765/circle1_kix7j5.png" width="30" height="30" alt="SorceMoola Logo">
		</body>
		</html>
	`))
}

func HealthHandler(c *gin.Context) {
	response := gin.H{
		"status":  "ok",
		"service": "Sorcemoola API",
		"version": "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}
