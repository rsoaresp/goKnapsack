package main

import "fmt"

// A structure to hold the properties the knapsack
// problem, the values and weights of the products.
type Knapsack struct {
	weights []int
	values  []int
}


// The maximum or minimum element of a slice, returning
// its position and the corresponding value.
func extrema(array []int, op string) (int, int) {

	var pos int = 0
	var ext int = array[0]

	var opString string = op[0:3]

	if opString == "min" {
		for i, value := range array {
			if value <= ext {
				ext = value
				pos = i
			}
		}
	} else if opString == "max" {
		for i, value := range array {
			if value >= ext {
				ext = value
				pos = i
			}
		}
	} else {
		panic("Operation unknow")
	}

	return pos, ext
}


// The unbounded knapsack problem using a dynamic programming algorithm.
// More information at https://en.wikipedia.org/wiki/Knapsack_problem
func (k *Knapsack) unboundedKP(limit int) ([][]int, []int) {

	profit := make([]int, limit+1)

	// items is a 2D slice that will contain the set of items that must
	// be included in the knapsack to optimize the profit at weight w.
	var pos int = 0
	items := make([][]int, limit+1)
	items[0] = make([]int, 1)

	for w := 1; w <= limit; w++ {

		temp := make([]int, w+1)

		for i, itemWeight := range k.weights {
			if itemWeight <= w {
				temp[i] = k.values[i] + profit[w-k.weights[i]]
			}
		}

		// We look for the combination that optimizes the profit
		// for a knapsack with weight w. The list of items is the
		// one just found plus the optimal set with total capacity
		// reduced by the weight of the chosen item.
		pos, profit[w] = extrema(temp, "max")
		if w - k.weights[pos] >= 0 {		
			items[w] = append(items[w - k.weights[pos]], k.weights[pos])
		}
	}

	return items, profit
}


func main() {
	// An example of how to call the knapsack problem.
	// We create the structure and call the method.
	w := [3]int{1, 2, 3}
	k := [3]int{7, 10, 21}

	res := Knapsack{weights: w[:], values: k[:]}

	i, m := res.unboundedKP(5)
	fmt.Println(i)
	fmt.Println(m)

}
