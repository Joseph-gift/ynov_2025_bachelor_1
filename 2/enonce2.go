package main

import "fmt"

func main() {
	var temp1 float32
	var temp2 float32
	var nom string

	fmt.Print("Quel est votre nom: ")
	fmt.Scan(&nom)
	//Enregister la temperature de la matine
	fmt.Print("Entree la temperature de la matinee: ")
	fmt.Scan(&temp1)
	// Enregister la temperature de la soire
	fmt.Print("Entrée la temperature de l'apres-midi: ")
	fmt.Scan(&temp2)
	// Afficher la temperature
	if temp1 > temp2 {
		fmt.Printf("La temperature la plus basse est: %g \n", temp2)
		fmt.Printf("La temperature la plus haute est: %g \n", temp1)
	}else if temp1 < temp2 {
		fmt.Printf("La temperature la plus haute est: %g \n", temp2)
		fmt.Printf("La temperature la plus basse est: %g \n", temp1)
	}else{
		fmt.Printf("La temperature la plus basse est: %g \n", temp2)
		fmt.Printf("La temperature la plus haute est: %g \n", temp1)
	}
}
