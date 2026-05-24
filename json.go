package main
import (
	"encoding/json"
	"fmt"
	"os"
)
type Fam struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
func main() {
	var members []Fam
	var x int
	fmt.Print("Enter number of family members: ")
	_, tt := fmt.Scan(&x)
	if tt != nil {
		fmt.Println("Invalid input! Please enter a number.")
		return
	}
	for i := 0; i<x; i++ {
		var member Fam
		fmt.Print("Enter name: ")
		fmt.Scan(&member.Name)
		fmt.Print("Enter age: ")
		_, tt := fmt.Scan(&member.Age)
		if tt != nil {
			fmt.Println("Invalid age! Please enter a number.")
			return
		}
		fmt.Print("Enter gender: ")
		fmt.Scan(&member.Gender)
		members = append(members, member)
	}
	file, err := os.Create("famm.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	tt = encoder.Encode(members)
	if tt != nil {
		fmt.Println("Error writing JSON:", tt)
		file.Close()
		return
	}
	file.Close()
	fmt.Println("\nData saved successfully in famm.json")
	data, tt := os.ReadFile("famm.json")
	if tt != nil {
		fmt.Println("Error reading file:", tt)
		return
	}
	var newmemb []Fam
	tt=json.Unmarshal(data,&newmemb)
	if tt!=nil{
		fmt.Println("Error parsing JSON:",tt)
		return
	}
	fmt.Println("\n Parsed struct data:")
	fmt.Printf("%+v\n",newmemb)
}