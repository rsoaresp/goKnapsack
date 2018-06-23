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
func (k *Knapsack) unboundedKP(limit int) ([]int, []int) {

	items := make([]int, limit+1)
	profit := make([]int, limit+1)

	for w := 1; w <= limit; w++ {

		temp := make([]int, w+1)

		for i, itemWeight := range k.weights {
			if itemWeight <= w {
				temp[i] = k.values[i] + profit[w - k.weights[i]]
			}
		}

		items[w], profit[w] = extrema(temp, "max")
	}

	return items, profit
}

func main() {

	w := [3]int{1, 2, 3}
	k := [3]int{7, 10, 21}

	res := Knapsack{weights: w[:], values: k[:]}

	i, m := res.unboundedKP(5)
	fmt.Println(i)
	fmt.Println(m)

}
