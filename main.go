package main

import (
	"log"
	"net/http"

	"github.com/aqwervinker/GO-distributed-cache/cache"
	"github.com/aqwervinker/GO-distributed-cache/metrics"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Инициализация кэша
	distributedCache := cache.NewDistributedCache()

	// Маршрут для метрик
	router.GET("/metrics", gin.WrapH(metrics.MetricsHandler()))

	// Маршруты для операций с кэшем
	router.GET("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, err := distributedCache.Get(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	router.POST("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		var body struct {
			Value string `json:"value"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if exists := distributedCache.Exists(key); exists {
			if err := distributedCache.Set(key, body.Value); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Status(http.StatusOK)
			return
		}
		if err := distributedCache.Set(key, body.Value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusCreated)
	})

	router.DELETE("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		if err := distributedCache.Delete(key); err != nil {
			if err == cache.ErrKeyNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
