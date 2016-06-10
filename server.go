package main

import (
	//"io"
  "log"
	"net/http"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  //"strconv"
  "os"
  "bufio"
  "strings"
)

func main() {
  fmt.Println("Starting App01")
  readThings()

  // load wiki titles into memory

  router := httprouter.New()
  router.GET("/goapi/:query", responseHandler)
	log.Fatal(http.ListenAndServe(":9761", router))
}

func responseHandler(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
  //fmt.Println(param.ByName("query"))

  // var letters [13]string
  // for i := 0; i < 13; i++ {
	//    letters[i] = ( "" + param.ByName("query") + strconv.Itoa(i) )
	// }
  // load wiki into slice
  // string match
  // add words
  zzquery := param.ByName("query")

  zmatches := GetMatches(zzquery, 35, &globalVarListThing)

  jsonArray, _ := json.Marshal(zmatches)
  w.Write(jsonArray)
  return
  //w.Write([]byte("A response"))
  //io.WriteString(w, "Hello world!")
}

//make this trie structure
var globalVarListThing []string

func readThings() {
  zpath := "/home/ubuntu/goWorkSpace/src/github.com/EricRiney/GoFish/theData/smallWordsEn.txt"
  filePath, err := os.Open(zpath)
  if err != nil {
    panic(err)
  }

  ericScan := bufio.NewScanner(filePath)
  for ericScan.Scan() {
    //load into trie structure
    globalVarListThing = append(globalVarListThing,ericScan.Text())
  }
  fmt.Println("done doing shit")

}









//function to query trie structure

//zhao crap
func GetMatches(queryS string, rCount uint64, sliceData *[]string) *[]string {
	var tempSlice []string
	for i := 0; ( i < len(*sliceData) )&&( uint64(len(tempSlice)) < rCount ); i++ {
		if strings.HasPrefix(strings.ToLower((*sliceData)[i]), queryS) {
			tempSlice = append(tempSlice,(*sliceData)[i])
		}
	}
	return &tempSlice
}
