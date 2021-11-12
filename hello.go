package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	//ROOT
	r.GET("/", root)
	//Agregar
	r.GET("/agregar", agregar)
	//ALUMNOS
	r.POST("/alumno_agregar", alumno_agregar)
	//MATERIAS

	r.Run(":3000")
}

//index.html
func root(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//Agregar
func agregar(c *gin.Context) {
	c.HTML(200, "agregar.html", nil)
}

//Alumno
func alumno_agregar(c *gin.Context) {
	//nombre
	nombre := c.PostForm("nombre")
	//materia
	materia := c.PostForm("materia")
	//calificacion
	calificacion := c.PostForm("calificacion")
	c.HTML(http.StatusOK, "alumno_agregado.html", gin.H{
		"nombre":       nombre,
		"materia":      materia,
		"calificacion": calificacion,
	})
}

//Materia
