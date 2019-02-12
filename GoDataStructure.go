package main
import "fmt"

func main() {
	// var x [5]int
 // 	x[3] = 1000
 //	fmt.Println(x)

 	var name [50]string //array of 50 string elements
	 fmt.Println(name)
	 name[43] = "Welcome to Go Programming"
	 fmt.Println(name)
	 var x [100]string
	 fmt.Println("Length of arr x:", len(x)) //calculate length of array
	 z := [5]int{43, 78, 54, 23, 22}
	 fmt.Println(z) 
	 for i := 65; i <= 122; i++ {
	 x[i-65] = string(i)
	}
	fmt.Println(x)

	// HashTables or Map 
	var dictionary map[string]int //map using var
	dictionary[“Zero”]=1
    fmt.Println(dictionary[“Zero”]) //Accessing value using key
    var details = map[int]string{101: "Jordan", 102: "Roger", 103: "Rafel"} //using var 
	fmt.Println(details)
	 fmt.Println(details[102])
	 countries := make(map[string]string) //Shorthand and make
	 countries["US"] = "United States"
	 countries["IND"] = "India"
	fmt.Println(countries)

}