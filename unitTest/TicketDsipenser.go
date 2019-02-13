package main
import "fmt"

type StationPriceData struct {
	station_name string
	price int
}

var station_price_map map[string]StationPriceData

func main() {
	station_price_map = make(map[string]StationPriceData)
	station_price_map["Churchgate"] = StationPriceData{
		"Andheri", 20,
	}
	station_price_map["Santacruz"] = StationPriceData{
		"Dadar", 10,
	}
	station_price_map["Dadar"] = StationPriceData{
		"MarineLines", 5,
	}
	fmt.Println(station_price_map["Dadar"].station_name, " - station name, ", station_price_map["Dadar"].price, " - price")
}