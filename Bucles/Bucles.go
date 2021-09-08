package main

import "fmt"

func main() {
	//Ciclo For Incrementando la variable a
	var a int
	fmt.Println("Ingresa un numero:")
	fmt.Scanln(&a)
	/*for a < 10 {
		a++
		fmt.Println(a)
	}*/
	//Tabla de multiplicar
	for i := 0; i <= 12; i++ {
		/*if i == 5 {
			continue
		}
		if i == 7 {
			break
		}*/
		result := i * a
		fmt.Println(i, "X", a, "=", result)
	}

}
