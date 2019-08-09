package main
import (
  "fmt"
  "net/http"
  "log"
  "runtime"
)

func main(){
  fmt.Println("in webserver main")
  http.HandleFunc("/",Hello)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func Hello(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"Hello, I am running on %s on %s CPU",runtime.GOOS,runtime.GOARCH)

}
