package main
import (
	"fmt"
	"net/http"
)
func main(){
	x:=http.FileServer(http.Dir("./static"))
	http.Handle("/static/",http.StripPrefix("/static/",x))
	fmt.Println("Server running on http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}