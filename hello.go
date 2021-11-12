package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	r.GET("/", root)                                      //menu principal
	r.GET("/agregar", agregar)                            //formulario de registro
	r.GET("/alumno_buscar", alumno_buscar)                //formulario para consultar promedio de un alumno
	r.GET("/alumno_listado", alumno_listado)              //pagina para ver el listado de los alumnos
	r.POST("/alumno_agregado", alumno_agregado)           //pagina de alumno agregado
	r.POST("/alumno_calificacion", alumno_calificacion)   //pagina para ver la calificacion del alumno
	r.GET("/materia_buscar", materia_buscar)              //formulario para consultar promedio de una materia
	r.POST("/materia_calificacion", materia_calificacion) //pagina para ver la calificacion de la materia

	r.GET("/general", general) //pagina para ver la calificacion

	r.Run(":3000")
}

//index.html
func root(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//Promedio General
func general(c *gin.Context) {
	c.HTML(200, "general.html", nil)
}

//Agregar
func agregar(c *gin.Context) {
	c.HTML(200, "agregar.html", nil)
}

//Alumno
func alumno_agregado(c *gin.Context) {
	nombre := c.PostForm("nombre")
	materia := c.PostForm("materia")
	calificacion := c.PostForm("calificacion")
	c.HTML(http.StatusOK, "alumno_agregado.html", gin.H{
		"nombre":       nombre,
		"materia":      materia,
		"calificacion": calificacion,
	})
}
func alumno_buscar(c *gin.Context) {
	c.HTML(200, "alumno_buscar.html", nil)
}
func alumno_calificacion(c *gin.Context) {
	nombre := c.PostForm("nombre")
	c.HTML(http.StatusOK, "alumno_calificacion.html", gin.H{
		"nombre": nombre,
	})
}
func alumno_listado(c *gin.Context) {
	//Leer datos de los alumnos en la bd
	c.HTML(200, "alumno_listado.html", nil)
}

//Materia
func materia_buscar(c *gin.Context) {
	c.HTML(200, "materia_buscar.html", nil)
}
func materia_calificacion(c *gin.Context) {
	materia := c.PostForm("materia")
	c.HTML(http.StatusOK, "materia_calificacion.html", gin.H{
		"materia": materia,
	})
}
