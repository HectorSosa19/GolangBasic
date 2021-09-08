package main

import (
	"fmt"
)

func main() {
	//Fuciona como arreglo map
	mapa := map[string]int{
		"Uno":  1,
		"Dos":  2,
		"Tres": 3,
	}

	fmt.Println(mapa)
	mapa["Cuatro"] = 4
	mapa["Cinco"] = 5
	fmt.Println(mapa)
	fmt.Println(mapa["Cinco"])
	delete(mapa, "Cinco")
	fmt.Println(mapa)
	mapa2 := make(map[int]string)
	mapa2[1] = "Vaca"
	mapa2[2] = "Perro"
	mapa2[3] = "Lagarto"
	fmt.Println(mapa2)
}
