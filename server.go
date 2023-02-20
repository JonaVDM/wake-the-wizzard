package main

import (
	"embed"
	"net/http"

	"github.com/JonaVDM/wake-the-wizzard/storage"
	"github.com/JonaVDM/wake-the-wizzard/wol"
	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var static embed.FS

func main() {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.GET("/api/wake/:id", wakePc)
	r.GET("/api/pc", getPc)
	r.POST("/api/pc", postPc)
	r.DELETE("/api/pc/:id", deletePc)
	r.StaticFileFS("/", "app.html", http.FS(CustomFS{static}))
	r.StaticFS("/assets", http.FS(CustomAssetsFS{static}))

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

		if err := wol.SendWol(item.Mac); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

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

	id, err := storage.Add(data.Name, data.Mac)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func deletePc(c *gin.Context) {
	id := c.Param("id")
	if err := storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
