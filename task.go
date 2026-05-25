package main
import(
	"encoding/json"
	"fmt"
	"net/http"
)
type Product struct{
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
}
func fetch(url string) (*Product, error) {
	response,tt:=http.Get(url)
	if tt!= nil {
		return nil,tt
	}
	defer response.Body.Close()
	var product Product
	tt = json.NewDecoder(response.Body).Decode(&product)
	if tt!= nil {
		return nil, tt
	}

	return &product, nil
}

func main() {

	url := "https://dummyjson.com/products/1"

	product, tt := fetch(url)
	if tt!= nil {
		fmt.Println("Error:",tt)
		return
	}

	
	fmt.Printf("%+v\n", *product)
}

