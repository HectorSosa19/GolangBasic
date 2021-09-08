package main

import (
	"fmt"
)

func main() {
	//Diferentes formas de usar el For y Range como un ciclo en un arreglo
	array1 := []int{1, 2, 3, 4}

	for i, num := range array1 {
		fmt.Println(i, "==>", num)
	}
	//Total de la suma con un ciclo
	suma := 0
	for _, num := range array1 {
		suma += num

	}
	fmt.Println("La suma total es", suma)

	colores := map[string]string{
		"Yellow": "Amarillo",
		"Blue":   "Azul",
		"Green":  "Verde",
	}
	for k, v := range colores {
		fmt.Println(k, "==>", v)

	}

}
