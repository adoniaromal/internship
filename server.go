package main
import(
	"fmt"
	"net/http"
)
func server(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello")
}
func main(){
	http.HandleFunc("/",server)
	fmt.Println("server running")
	http.ListenAndServe(":8080",nil)
}