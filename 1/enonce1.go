package main

import "fmt"

func main()  {

	var age uint8

	fmt.Print("Entre votre age: ")
	fmt.Scan(&age)
	// si la personne est majeur
	if age > 18{
		// print  majeur
		fmt.Println("majeur")
		// si non
	}else{
		// pritn minueur
		fmt.Println("mineur")
	}
}
