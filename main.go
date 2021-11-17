package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var materiasMap map[string]map[string]float64
var alumnosMap map[string]map[string]float64

type Alumnos struct {
	Name    string
	Subject string
	Grade   string
}

//-------------------------------------------------------| CÓDIGO BASE DE DATOS |-----------------------------------------------------------
type AlumnoBD struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Grade   string `json:"grade"`
}

func getAlumnos() []*AlumnoBD {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM alumnos")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var alumnos []*AlumnoBD
	for results.Next() {
		var u AlumnoBD
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name, &u.Subject, &u.Grade)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		alumnos = append(alumnos, &u)
	}

	return alumnos
}
func alumnoPage(w http.ResponseWriter, r *http.Request) {
	alumnos := getAlumnos()

	fmt.Println("Endpoint Hit: alumnoPage")
	json.NewEncoder(w).Encode(alumnos)
}

//-------------------------------------------------------| FIN CÓDIGO BASE DE DATOS |-------------------------------------------------------

type infoServer struct {
	Info []Alumnos
}

func (info *infoServer) Add(data Alumnos) {
	info.Info = append(info.Info, data)
}

var auxAlumnos infoServer

func registroCalif(data []string, reply *string) error {
	if _, err := alumnosMap[data[0]][data[1]]; err {
		*reply = "ERROR. Ya se ha registrado la calificacion"
		return nil
	}
	materiaAux := make(map[string]float64)
	alumnoAux := make(map[string]float64)
	calificacionAux, _ := strconv.ParseFloat(data[2], 64)
	materiaAux[data[1]] = calificacionAux
	alumnoAux[data[0]] = calificacionAux

	if _, err := alumnosMap[data[0]]; err {
		alumnosMap[data[0]][data[1]] = calificacionAux
		if _, err := materiasMap[data[1]]; err {
			materiasMap[data[1]][data[0]] = calificacionAux
		} else {
			materiasMap[data[1]] = alumnoAux
		}
	} else {
		alumnosMap[data[0]] = materiaAux
		if _, err := materiasMap[data[1]]; err {
			materiasMap[data[1]][data[0]] = calificacionAux
		} else {
			materiasMap[data[1]] = alumnoAux
		}
	}
	*reply = "Se ha registrado la calificacion de " + data[0]
	return nil
}

func obtenerPromedioAlumno(name string, reply *string) error {
	if _, err := alumnosMap[name]; err {
		var sum float64
		var subject int
		for _, grade := range alumnosMap[name] {
			sum = sum + grade
			subject = subject + 1
		}
		finalGrade := sum / float64(subject)
		*reply = "Promedio del alumno " + name + ": " + strconv.FormatFloat(finalGrade, 'f', 2, 64)
		return nil
	} else {
		*reply = "ERROR. No se han registrado calificaciones para este alumno"
		return nil
	}
}

func obtenerPromedioMateria(subject string, reply *string) error {
	if _, err := materiasMap[subject]; err {
		var sum float64
		var counter int
		for _, grade := range materiasMap[subject] {
			sum = sum + grade
			counter = counter + 1
		}
		finalGrade := sum / float64(counter)
		*reply = "Promedio general de la materia " + subject + ": " + strconv.FormatFloat(finalGrade, 'f', 2, 64)
		return nil
	} else {
		*reply = "ERROR. No se han registrado calificaciones para esta materia"
		return nil
	}
}

func obtenerPromedioGeneral(all int, reply *string) error {
	var sum float64
	var cont int
	for student := range alumnosMap {
		for _, grade := range alumnosMap[student] {
			sum = sum + grade
			cont = cont + 1
		}
	}
	finalGrade := sum / float64(cont)
	*reply = "Promedio general: " + strconv.FormatFloat(finalGrade, 'f', 0, 64)
	return nil
}

func alumnos(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		if req.RequestURI == "/alumnos" {
			fmt.Println(req.PostForm)
			data := Alumnos{Name: req.FormValue("nombre"), Subject: req.FormValue("materia"), Grade: req.FormValue("calificacion")}
			var info []string
			info = append(info, data.Name)
			info = append(info, data.Subject)
			info = append(info, data.Grade)
			var result string
			err := registroCalif(info, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				if result != "ERROR. Ya se ha registrado la calificacion" {
					auxAlumnos.Add(data)
				}
			}
			fmt.Println(auxAlumnos)
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				result,
			)
		}
		if req.RequestURI == "/promedioAlumno" {
			fmt.Println(req.PostForm)
			data := req.FormValue("nombre")
			var result string
			err := obtenerPromedioAlumno(data, &result)
			if err != nil {
				fmt.Println(err)
			}
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				result,
			)
		}
		if req.RequestURI == "/promedioMateria" {
			fmt.Println(req.PostForm)
			data := req.FormValue("nombre")
			var result string
			err := obtenerPromedioMateria(data, &result)
			if err != nil {
				fmt.Println(err)
			}
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				result,
			)
		}
	case "GET":
		var all int
		var result string
		err := obtenerPromedioGeneral(all, &result)
		if err != nil {
			fmt.Println(err)
		}
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioGeneral.html"),
			result,
		)
	}
}

func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)
	return string(html)
}

func main() {
	materiasMap = make(map[string]map[string]float64)
	alumnosMap = make(map[string]map[string]float64)
	http.HandleFunc("/form", form)
	http.HandleFunc("/alumnos", alumnos)
	http.HandleFunc("/promedioAlumno", alumnos)
	http.HandleFunc("/promedioMateria", alumnos)
	http.HandleFunc("/promedioGeneral", alumnos)
	http.HandleFunc("/alumnosbd", alumnoPage)
	fmt.Println("Corriendo servidor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
