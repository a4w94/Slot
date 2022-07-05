package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port      = ":8000"
	htmlroute = "../chart/*"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(htmlroute)
	r.GET("/", LoadChart)
	r.Run(port)

}

func LoadChart(c *gin.Context) {

	c.HTML(http.StatusOK, "MainGame倍率區間.html", nil)
	c.HTML(http.StatusOK, "MainGame倍率區間RTP.html", nil)

	c.HTML(http.StatusOK, "FreeGame倍率區間.html", nil)
	c.HTML(http.StatusOK, "FreeGame倍率區間RTP.html", nil)

}
