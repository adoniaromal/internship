package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
}
type keyy struct {
	Products []Product `json:"products"`
}
func fetch(url string) (*keyy, error) {
	response, tt := http.Get(url)
	if tt != nil {
		return nil, tt
	}
	defer response.Body.Close()
	var result keyy
	tt = json.NewDecoder(response.Body).Decode(&result)
	if tt != nil {
		return nil, tt
	}
	return &result, nil
}
func main() {
	url := "https://dummyjson.com/products"
	result, tt := fetch(url)
	if tt != nil {
		fmt.Println("Error:", tt)
		return
	}
	for _, product := range result.Products {
		fmt.Println("Title:", product.Title)
		fmt.Println("Description:", product.Description)
		fmt.Println("Price:", product.Price)
		fmt.Println()
	}
}