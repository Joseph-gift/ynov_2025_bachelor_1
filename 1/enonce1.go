package main

import "fmt"

func main()  {

	var age uint8
	var nom string
	var premon string

	fmt.Print("Entre votre age: ")
	fmt.Scan(&age)
	fmt.Print("Entre votre Nom: ")
	fmt.Scan(&nom)
	fmt.Print("Entre votre Prenom: ")
	fmt.Scan(&premon)
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
