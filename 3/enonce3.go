package main

import "fmt"

func main(){
	var x float32
	var v float32
	couleur:= "Jaune"

	fmt.Printf("Prix du produit: ")
	fmt.Scan(&v)
	fmt.Printf("Somme remise par le client: ")
	fmt.Scan(&x)
	//Comparer l'argent du client x et la valeur du produit v
	// Si x > v
	if x > v {
		// il  a donner asser et imprimer la monnaie
		fmt.Printf("Le client a donner assez. Sa monnaie est de %g \n",x-v )
		fmt.Printf("La couleur du ticket imprimer est %s \n", couleur)
	}else if x < v{
	// si v < x 
	// il a pas donner assez 
	fmt.Println("Le client n'as pas donner assez")
	}else{
		// si x == v il a donner la bonne somme
		fmt.Println("Il a donné exactement le bonne somme")
	}
}