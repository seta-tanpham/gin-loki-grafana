package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func setupLogger() zerolog.Logger {
	// Đảm bảo thư mục tồn tại
	os.MkdirAll("logs", os.ModePerm)

	file, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	multi := zerolog.MultiLevelWriter(os.Stdout, file)

	logger := zerolog.New(multi).With().Timestamp().Logger()
	return logger
}

func main() {
	logger = setupLogger()

	r := gin.New()

	// Ghi log mọi request
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", latency).
			Msg("Incoming request")
	})

	r.GET("/ping", func(c *gin.Context) {
		logger.Info().Msg("Ping endpoint called")
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Simulated Success
	r.GET("/success", func(c *gin.Context) {
		log.Info().Msg("Success endpoint called")
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Simulated Failure
	r.GET("/fail", func(c *gin.Context) {
		log.Error().Msg("Failure endpoint called")
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
	})

	logger.Info().Msg("Starting server on :8080")
	r.Run(":8080")
}
