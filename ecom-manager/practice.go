package main

import (
	"fmt"
)

func main() {
	// var numbers [10]int
	// integers := [4]int{10, 20, 30, 40}
	// integerSlices := []int{3, 4, 5, 6}
	// fmt.Println(integerSlices[2:4])

	// countryNames := []string{"India", "Spain", "France", "Germany"}
	// countryNames = append(countryNames, "Japan")
	// // fmt.Println(countryNames)
	// countryNames = append(countryNames, "Australia")
	// // fmt.Println(countryNames)
	// countryNames = append(countryNames, "Greece", "Dombivli")
	// // fmt.Println(countryNames)
	// for index, country := range countryNames {
	// 	fmt.Printf("Index = %d, Country Name = %s\n", index, country)
	// }
	// // fmt.Println(countryNames)

	// fmt.Println(test.WordCount("the quick brown dog jumped over the lazy dog"))
	// fmt.Println(test.Greeting())

	byteStr := []byte("Iis")
	fmt.Println(len(byteStr), byteStr[0])
	fmt.Println(string(byteStr))
}
