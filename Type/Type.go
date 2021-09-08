package main

import (
	"fmt"
)

type entero int
type cadena string

func main() {
	fmt.Print("")
	/*var dia entero
	fmt.Println(dia)
	var nombre cadena
	fmt.Println(nombre)
	curso1 := alumno{
		nombre:    "Hector",
		numero:    10,
		materia:   "Ingles",
		habilidad: []string{"Hablar rapido", "2"},
	}



	curso2 := new(alumno)
	fmt.Println(curso2)
	curso2.nombre = "Jose"
	curso2.numero = 20
	curso2.curso = "Español"
	curso2.habilidad = []string{"Front-End"}*/

	type alumno struct {
		nombre    string
		numero    int
		curso     string
		habilidad []string
	}
}
