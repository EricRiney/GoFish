package main

import (
	//"io"
  "log"
	"net/http"
  "fmt"
  "github.com/julienschmidt/httprouter"
)

func main() {
  fmt.Println("Starting App")
  
  // load wiki titles into memory

  router := httprouter.New()
  router.GET("/goapi/:query", responseHandler)
	log.Fatal(http.ListenAndServe(":9761", router))
}

func responseHandler(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
  fmt.Println(param.ByName("query"))
  
  
  //w.Write([]byte("A response"))
  //io.WriteString(w, "Hello world!")
}
