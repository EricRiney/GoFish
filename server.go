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
  // make slice
  var letters [13]string
  // for loop slice
  for i := 0; i < 13; i++ {
	   letters[i] = (   "" + param.ByName("query") + strconv.Itoa(i)   )
	}

  // populate slice with 13 entries of the query + counter
  // like: query0, query1, etc....
  // return that



  jsonArray, _ := json.Marshal(letters)
  w.Write(jsonArray)
  return
  //w.Write([]byte("A response"))
  //io.WriteString(w, "Hello world!")
}
