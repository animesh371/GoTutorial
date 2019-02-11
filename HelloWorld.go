package main 
import ("fmt")

func add(first float32, second float32) float32 {
	return first + second
}
func main() {
	var num1 float32 = 32.1;
	var num2 float32 = 10.2;
	fmt.Println("Addition of two numbers is  ", add(num1, num2))

}