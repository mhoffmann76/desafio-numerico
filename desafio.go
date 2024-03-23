package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}


	println("Desafio parte 2")

	for i := 0; i <=100; i++ {

		if i%3 == 0{
			println("Pin")

		}else if i%5 == 0{
			println("Pan")
		}
	}
}
