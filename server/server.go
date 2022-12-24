package main

import (
	"net/http"
	"strings"

	"github.com/JonaVDM/wake-the-wizzard/storage"
	"github.com/JonaVDM/wake-the-wizzard/wol"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/wake/:id", wakePc)
	r.GET("/api/pc", getPc)
	r.POST("/api/pc", postPc)
	r.DELETE("/api/pc/:id", deletePc)

	r.Run(":3080")
}

func wakePc(c *gin.Context) {
	id := c.Param("id")
	items, err := storage.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	for _, item := range items {
		if item.Id != id {
			continue
		}

		mac := strings.ReplaceAll(item.Mac, ":", "")
		wol.SendWol(mac)
		c.String(http.StatusOK, "ok")
		return
	}

	c.String(http.StatusBadRequest, "pc not found")
}

func getPc(c *gin.Context) {
	items, err := storage.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, items)
}

func postPc(c *gin.Context) {
	var data struct {
		Name string `json:"name"`
		Mac  string `json:"mac"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := storage.Add(data.Name, data.Mac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "ok")
}

func deletePc(c *gin.Context) {
	id := c.Param("id")
	if err := storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "ok")
}
