package server

import (
	"net"
	"net/http"

	"github.com/JonaVDM/wake-the-wizzard/frontend"
	"github.com/JonaVDM/wake-the-wizzard/storage"
	"github.com/JonaVDM/wake-the-wizzard/wol"
	"github.com/gin-gonic/gin"
)

func Serve(addr string) error {
	r := initGin()
	return r.Run(addr)
}

func ServeListner(ln net.Listener) error {
	r := initGin()

	return http.Serve(ln, r.Handler())
}

func initGin() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.GET("/api/wake/:id", wakePc)
	r.GET("/api/pc", getPc)
	r.POST("/api/pc", postPc)
	r.DELETE("/api/pc/:id", deletePc)
	r.StaticFileFS("/", "app.html", http.FS(frontend.NewCustomFS()))
	r.StaticFS("/assets", http.FS(frontend.NewAssetFs()))

	return r
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
