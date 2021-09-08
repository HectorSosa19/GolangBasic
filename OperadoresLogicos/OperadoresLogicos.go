package main

import "fmt"

func main() {
	/* El Not se representa de la siguiente manera "!"
	   El And se representa de la siguiente manera "&&"
	   El And se representa de la siguiente manera "||"
	*/
	fmt.Println(!true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(3 == 1 || 2 == 2)
	fmt.Println(1 == 1 && 2 == 2)
	fmt.Println(2 == 2)
}
