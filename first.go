package main

import "fmt"


func main()  {
	var Nums [5]int
		Nums[0] = 0
		Nums[1] = 45
		Nums[2] = 65
		Nums[3] = 80
		Nums[4] = 90

	for  _,i := range Nums {
		fmt.Println(i)
	}
}