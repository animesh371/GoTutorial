package main 
import "fmt"

const kmh_multiple float32 = 1.615

type car struct {
	wheels int 
	seats int 
	max_speed float32


}

func (c car) kmhr() float32{
	return c.max_speed * kmh_multiple
}

func main() {
	// bmw := car{
	// 	wheels: 4, 
	// 	seats:5, 
	// 	max_speed: 200}

	merc := car{
		wheels: 4, 
		seats:4, 
		max_speed: 130.0}	

	fmt.Println(merc.max_speed, merc.kmhr())


}