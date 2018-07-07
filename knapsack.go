package main

import "fmt"

// Knapsack structure holds the properties the knapsack
// problem, the values and weights of the products.
type Knapsack struct {
	weights []int
	values  []int
}


// Size returns the number of elements of the problem.
func (k *Knapsack) Size() int {
	return len(k.weights)
}


// Unbounded solves the knapsack problem when there is no limit
// on the numbers of a particular item to be carried.
// It does so using the well know algorithm of dynamic programming
// stated in https://en.wikipedia.org/wiki/Knapsack_problem
func (k *Knapsack) Unbounded(limit int) ([][]int, []int) {

	var temp int = -1

	profit := make([]int, limit+1)

	items := make([][]int, limit+1)
	for i := 1; i <= limit; i++ {
		items[i] = make([]int, 0)
	}

	for w := 1; w <= limit; w++ {
		for i, itemWeight := range k.weights {
			if itemWeight <= w {
				temp = k.values[i] + profit[w-itemWeight]
				if temp > profit[w] {
					profit[w] = temp
					items[w] = append(items[w-itemWeight], itemWeight)
				}
			}
		}
	}

	return items, profit
}


// Bounded solves the knapsack problem when only one instance
// if each item can be carried.
func (k *Knapsack) Bounded(limit int) ([][][]int, [][]int) {

	var temp int = -1

	profit := make([][]int, k.Size()+1)
	for i := range profit {
		profit[i] = make([]int, limit+1)
	}

	// In order to know which items must be carried for
	// the combinations of total weight and value we must
	// have a list for each pair.
	items := make([][][]int, k.Size()+1)
	for i := 0; i < k.Size()+1; i++ {
		items[i] = make([][]int, limit+1)
		for j:=0; j < limit+1; j++ {
			items[i][j] = make([]int, 0)
		}
	}


	for i := 1; i < k.Size()+1; i++ {

		value := k.values[i-1]		
		weight := k.weights[i-1]

		for w := 0; w <= limit; w ++ {

			items[i][w] = items[i-1][w]

			if weight > w {
				profit[i][w] = profit[i-1][w]
			} else {
				temp = profit[i-1][w-weight] + value
				if temp > profit[i-1][w] {
					profit[i][w] = temp
					items[i][w] = append(items[i-1][w-weight], weight)
				} else {
					profit[i][w] = profit[i-1][w]
				}
			}
		}
	}

    	return items, profit
}


func main() {
	// An example of how to call the knapsack problem.
	// We create the structure and call the method.
	w := [4]int{5, 4, 6, 3}
	k := [4]int{10, 40, 30, 50}

	res := Knapsack{weights: w[:], values: k[:]}

	a,b := res.Bounded(10)
	fmt.Println("--------------")
	fmt.Println(a,b)

}
