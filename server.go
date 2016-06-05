package main

import (
	//"io"
  "log"
	"net/http"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  "strconv"
)

func main() {
  fmt.Println("Starting App")

  // load wiki titles into memory

  router := httprouter.New()
  router.GET("/goapi/:query", responseHandler)
	log.Fatal(http.ListenAndServe(":9761", router))
}

func responseHandler(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
  //fmt.Println(param.ByName("query"))

  var letters [13]string
  for i := 0; i < 13; i++ {
	   letters[i] = ( "" + param.ByName("query") + strconv.Itoa(i) )
	}
  // load wiki into slice
  // string match
  // add words
  // var responce = param.ByName("query")



  jsonArray, _ := json.Marshal(letters)
  w.Write(jsonArray)
  return
  //w.Write([]byte("A response"))
  //io.WriteString(w, "Hello world!")
}
