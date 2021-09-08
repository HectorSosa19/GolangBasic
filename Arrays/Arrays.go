package main

import "fmt"

func main() {
	//Crear arreglos
	var array [5]int
	array[0] = 0
	array[1] = 1
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println(array)
	arrays := [5]string{"perro,", "gato,", "vaca,", "paloma"}
	fmt.Println(arrays)

	//Slicen
	var slice []string
	slice = append(slice, "rojo")
	slice = append(slice, "verde", "azul")
	fmt.Println(slice)
}
