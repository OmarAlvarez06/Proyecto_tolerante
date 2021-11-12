package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	//RUTAS
	r.GET("/", root)

	r.Run(":3000")
}

//index.html
func root(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//Alumno
//Materia
