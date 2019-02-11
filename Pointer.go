package main 
import ("fmt")

func main() {
	x := 15
	a := &x
	fmt.Println(x, a)
	fmt.Println(x, *a)
	*a = *a * *a
	fmt.Println(*a, x)

}