package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func handleHealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func handleGet(c *gin.Context) {
	entity := c.Param("entity")

	routeState.RLock()
	defer routeState.RUnlock()

	entityData, exists := routeState.data[entity]
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "success",
			"data":       make([]interface{}, 0),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "success",
		"data":       entityData,
	})
}

func handleGetByID(c *gin.Context) {
	entity := c.Param("entity")
	id := c.Param("id")

	routeState.RLock()
	defer routeState.RUnlock()

	entityData, exists := routeState.data[entity]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "not found",
		})
		return
	}

	for _, data := range entityData {
		if data.(map[string]interface{})["id"] == id {
			c.JSON(200, gin.H{
				"statusCode": http.StatusOK,
				"message":    "success",
				"data":       data,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"statusCode": http.StatusNotFound,
		"message":    "not found",
	})
}

func handlePost(c *gin.Context) {
	entity := c.Param("entity")

	routeState.Lock()
	defer routeState.Unlock()

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "bad request",
			"error":      "invalid body format",
		})
	}

	if _, exists := routeState.data[entity]; !exists {
		routeState.data[entity] = make([]interface{}, 0)
	}
	id := ksuid.New()
	data["id"] = id.String()
	data["createdAt"] = id.Time().Unix()
	routeState.data[entity] = append(routeState.data[entity], data)

	c.JSON(200, gin.H{
		"statusCode": http.StatusCreated,
		"message":    "success",
		"data":       gin.H{"id": id.String()},
	})
}

func handlePut(c *gin.Context) {
	entity := c.Param("entity")
	id := c.Param("id")

	routeState.Lock()
	defer routeState.Unlock()

	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "bad request",
			"error":      "invalid body format",
		})
	}

	entityData, exists := routeState.data[entity]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "not found",
		})
		return
	}

	for i, data := range entityData {
		if data.(map[string]interface{})["id"] == id {
			tmpId := data.(map[string]interface{})["id"]
			tmpCreatedAt := data.(map[string]interface{})["createdAt"]

			data = body
			data.(map[string]interface{})["id"] = tmpId
			data.(map[string]interface{})["createdAt"] = tmpCreatedAt

			routeState.data[entity] = append(entityData[:i], entityData[i+1:]...)
			routeState.data[entity] = append(routeState.data[entity], data)

			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusOK,
				"message":    "success",
				"data":       gin.H{"id": id},
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"statusCode": http.StatusNotFound,
		"message":    "not found",
	})
}

func handleDelete(c *gin.Context) {
	entity := c.Param("entity")
	id := c.Param("id")

	routeState.Lock()
	defer routeState.Unlock()

	entityData, exists := routeState.data[entity]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "not found",
		})
		return
	}

	for i, data := range entityData {
		if data.(map[string]interface{})["id"] == id {
			routeState.data[entity] = append(entityData[:i], entityData[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusOK,
				"message":    "success",
				"data":       gin.H{"id": id},
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"statusCode": http.StatusNotFound,
		"message":    "not found",
	})
}
