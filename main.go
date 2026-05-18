package main
import(
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	file,tt :=os.Create("fam.csv")
	if tt!=nil{
		fmt.Println("Error creating file:",tt)
		return
	}
	wr:=csv.NewWriter(file)
	wr.Write([]string{"Name","Age","Gender"})
	var x int
	fmt.Print("Enter number of family members: ")
	fmt.Scan(&x)
	for i :=0;i<x;i++{
		var name string
		var age int
		var gender string
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Age: ")
		_,tt :=fmt.Scan(&age)
		if tt!=nil{
			fmt.Println("Invalid!.Please enter a number.")
			return
		}
		fmt.Print("Gender: ")
		fmt.Scan(&gender)
		wr.Write([]string{name,strconv.Itoa(age),gender})
	}
	wr.Flush()
	file.Close()
	fmt.Println("\nData saved successfully in fam.csv")
	readFile,tt :=os.Open("fam.csv")
	if tt!= nil {
		fmt.Println("Error opening file:",tt)
		return
	}
	defer readFile.Close()
	rd :=csv.NewReader(readFile)
	records,tt :=rd.ReadAll()
	if tt != nil {
		fmt.Println("Error reading file:",tt)
		return
	}
	for _,record:=range records{
		fmt.Println(strings.Join(record,","))
	}
}
