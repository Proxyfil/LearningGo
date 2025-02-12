package main

import (
	"fmt";
	"math/rand";
	"time";
)

func initGrille(n, m int) [][]int {
	grille := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grille[i] = append(grille[i], rand.Intn(2))
		}
	}

	return grille
}

func compterVoisins(grille [][]int, i int, j int) int {
	compteur := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if i+x >= 0 && i+x < len(grille) && j+y >= 0 && j+y < len(grille[0]) {
				compteur += grille[i+x][j+y]
			}
		}
	}
	return compteur
}

func update(grille [][]int) [][]int {
	n := len(grille)
	m := len(grille[0])
	newGrille := make([][]int, n)

	for i := 0; i < n; i++ {
		newGrille[i] = make([]int, m)
		for j := 0; j < m; j++ {
			voisins := compterVoisins(grille, i, j)
			if grille[i][j] == 1 {
				if voisins < 2 || voisins > 3 {
					newGrille[i][j] = 0
				} else {
					newGrille[i][j] = 1
				}
			} else {
				if voisins == 3 {
					newGrille[i][j] = 1
				} else {
					newGrille[i][j] = 0
				}
			}
		}
	}

	return newGrille
}

func afficherGrille(grille [][]int) {
	affichage := make([][]string, len(grille))
	for ligne := range grille {
		for colonne := range grille[ligne] {
			if(grille[ligne][colonne] == 1) {
				affichage[ligne] = append(affichage[ligne], "⬜")
			} else {
				affichage[ligne] = append(affichage[ligne], "⬛")
			}
		}
	}

	for ligne := range affichage {
		fmt.Println(affichage[ligne])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	grille := initGrille(50, 50)

	for i := 0; i < 5000; i++ {
		grille = update(grille)

		afficherGrille(grille)
		fmt.Println()

		time.Sleep(100 * time.Millisecond)
	}
}