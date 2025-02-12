package main

import (
	"fmt";
	"math/rand";
	"time"
)

func estBissextile(annee int) bool {
	return annee%4 == 0 && (annee%100 != 0 || annee%400 == 0)
}

func estPremier(nombre int) bool {
	if nombre < 2 {
		return false
	}
	for i := 2; i*i <= nombre; i++ {
		if nombre%i == 0 {
			return false
		}
	}
	return true
}

func premiersNombresPremiers(taille int) []int {
	var liste []int
	for i := 0; len(liste) < taille; i++ {
		if estPremier(i) {
			liste = append(liste, i)
		}
	}
	return liste
}

func genererTableauAleatoire(taille int) []int {
	var liste []int
	for i := 0; i < taille; i++ {
		liste = append(liste, rand.Intn(100))
	}
	return liste
}

func triBulles(tableau []int) []int {
	for i := 0; i < len(tableau); i++ {
		for j := 0; j < len(tableau)-1; j++ {
			if tableau[j] > tableau[j+1] {
				tableau[j], tableau[j+1] = tableau[j+1], tableau[j]
			}
		}
	}
	return tableau
}

func rechercheDichotomique(tableau []int, nombre int) (bool, int) {
	subtableau := tableau
	for i := 0; i*i < len(tableau); i++ {
		if subtableau[len(subtableau)/2] == nombre {
			return true, len(subtableau)/2
		} else if subtableau[len(subtableau)/2] < nombre {
			subtableau = subtableau[len(subtableau)/2:]
		} else {
			subtableau = subtableau[:len(subtableau)/2]
		}
	}
	return false, -1
}

func organiserParTaille(tableau []string) map[int][]string {
	m := make(map[int][]string)
	for _, element := range tableau {
		m[len(element)] = append(m[len(element)], element)
	}
	return m
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tableau := genererTableauAleatoire(10)
	tableauTriee := triBulles(genererTableauAleatoire(10))

	fmt.Println(estBissextile(2020))
	fmt.Println(estPremier(7))
	fmt.Println(premiersNombresPremiers(10))
	fmt.Println(tableau)
	fmt.Println(tableauTriee)
	fmt.Println(rechercheDichotomique(tableauTriee, 51))
	fmt.Println(organiserParTaille([]string{"un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix"}))
}