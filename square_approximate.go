package main

/// my attempt to make a fast square root appromimator
/// method is based on a youtube video i stumbled upon
/// I may as well credit i guess https://www.youtube.com/watch?v=PJHtqMjrStk&t=303s&ab_channel=tecmath

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func generate_sqr_list(n int) []int { //makes an array of full squares limited by users input
	var intrange []int
	for i := 0; i <= n; i++ {
		intrange = append(intrange, i*i)
	}
	return intrange
}

func binarySearch(a []int, search int) (result int) { //searches numbers interval in square list
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		return
	case a[mid] > search:
		result = binarySearch(a[:mid], search)
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}

	return
}

func get_constant(n float64, array []int) float64 { //i calculate a constant used in a method for given number

	clst_rt := binarySearch(array, int(n)) - 1
	clst_sqrt := array[clst_rt]
	act_rt := math.Sqrt(n)
	delta := n - float64(clst_sqrt)
	return delta / ((act_rt - float64(clst_rt)) * float64(clst_rt))

}

//var index_storage []int
var const_storage []float64

//var avg float64

func make_lot_const(t []int, limit int) { //makes a constant list

	for i := 1; i < limit; i++ {
		number := i
		//index_storage = append(index_storage, number)
		const_storage = append(const_storage, get_constant(float64(number), t))
	}

}

func eval_root(t []int, number float64, constant_use bool) float64 {
	upperindex := binarySearch(t, int(number))
	var coefficient float64
	switch constant_use {
	case true:
		coefficient = const_storage[int(number)-1]
	default:
		coefficient = float64(2) + math.Pow(math.E, math.E-1)/number //thats my coeffitient approximation, using it avoids creating constant list
	}

	clst_sqrt := t[upperindex-1]
	clst_rt := float64(upperindex) - 1

	return clst_rt + ((number - float64(clst_sqrt)) / (coefficient * clst_rt))

}

//func calc_error(t []int) float64 { //// calculates error using approximated coefficient
//var sqrt_err float64
//for i := 2; i < 1000002; i++ {
//	rnd_flt := 2 + rand.Float64()*(1000000-2)
//	sqrt_err += math.Abs(math.Sqrt(rnd_flt) - eval_root(t, rnd_flt))
//}
//return sqrt_err / 1000000
//}

func main() {
	//fmt.Println(math.Sqrt(14455))
	var limit float64
	var tmp float64
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Enter n in 10^n as upper Integer limit for root calculation")

	_, err := fmt.Scanf("%f\n", &limit)
	limit = math.Pow10(int(limit))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Initializing square lists...")
	t := generate_sqr_list(int(limit))
	fmt.Println("If you want to initialize constants (10^-8 error) input 1, or input any other numeber to approximate (10^-4 error).")
	fmt.Println("Constant initialization will take more time and RAM.")
	var decisionint float64
	fmt.Scanf("%f\n", &decisionint)
	var decision bool
	switch decisionint {
	case 1:
		decision = true
		fmt.Println("Initializing constants...")
		make_lot_const(t, int(limit))
	default:
		decision = false
	}

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		fmt.Println("Proceed with a number to sqrt()")
		fmt.Scanf("%f\n", &tmp)
		if tmp < limit {
			fmt.Println(eval_root(t, tmp, decision))
		} else {
			fmt.Println("Your input is higher than inputted limit.")
		}

		fmt.Println()
		fmt.Println()
	}
}
