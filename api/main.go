package main
import(
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/dgrijalva/jwt-go"
)

var MySigingKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request){	
	fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){	
		if r.Header["Token"] != nil{
			
		}	
}

func handleRequests(){
	http.HandleFunc("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main(){
	fmt.Printf("Server started on port 9001\n")
	handleRequests()}
