package main

import (
	"fmt";
	"math/rand";
	"time";
	"math";
)

func initVec() [2]int {
	return [2]int{rand.Intn(2), rand.Intn(2)}
}

func addVec(vec1 [2]int, vec2 [2]int) [2]int {
	return [2]int{vec1[0] + vec2[0], vec1[1] + vec2[1]}
}

func subVec(vec1 [2]int, vec2 [2]int) [2]int {
	return [2]int{vec1[0] - vec2[0], vec1[1] - vec2[1]}
}

func multVec(vec1 [2]int, vec2 [2]int) [2]int {
	return [2]int{vec1[0] * vec2[0], vec1[1] * vec2[1]}
}

func normVec(vec [2]int) float64 {
	return math.Sqrt(math.Pow(float64(vec[0]), 2) + math.Pow(float64(vec[1]), 2))
}

func normalizeVec(vec [2]int) [2]float64 {
	norm := normVec(vec)
	return [2]float64{float64(vec[0]) / norm, float64(vec[1]) / norm}
}

func scalarVec(vec [2]int, scalar int) [2]int {
	return [2]int{vec[0] * scalar, vec[1] * scalar}
}

func dotVec(vec1 [2]int, vec2 [2]int) int {
	return vec1[0] * vec2[0] + vec1[1] * vec2[1]
}

func main() {

	rand.Seed(time.Now().UnixNano())

	vec1 := initVec()
	vec2 := initVec()

	fmt.Println("vec1:", vec1)
	fmt.Println("vec2:", vec2)
	fmt.Println("vec1 + vec2:", addVec(vec1, vec2))
	fmt.Println("vec1 - vec2:", subVec(vec1, vec2))
	fmt.Println("vec1 * vec2:", multVec(vec1, vec2))
	fmt.Println("norm(vec1):", normVec(vec1))
	fmt.Println("normalize(vec1):", normalizeVec(vec1))
	fmt.Println("vec1 * 3:", scalarVec(vec1, 3))
	fmt.Println("dot(vec1, vec2):", dotVec(vec1, vec2))

}