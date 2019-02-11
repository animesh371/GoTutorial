package main 
import "fmt"

const kmh_multiple float32 = 1.615

type car struct {
	wheels int 
	seats int 
	max_speed float32


}

func (c *car) kmhr(newspeed float32) float32{
	c.max_speed = newspeed
	return c.max_speed
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

	fmt.Println(merc.max_speed)
	fmt.Println(merc.kmhr(101.0))
	fmt.Println(merc.max_speed)



}