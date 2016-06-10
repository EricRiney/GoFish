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
  //"trie"
)

func main() {
  fmt.Println("Starting App")
  readThings()
  router := httprouter.New()
  router.GET("/goapi/:query", responseHandler)
	log.Fatal(http.ListenAndServe(":9761", router))
}

func responseHandler(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
  query := param.ByName("query")
  matches := GetMatches(query, 35, &globalVarListThing)
  jsonArray, _ := json.Marshal(matches)
  w.Write(jsonArray)
  return
}

//make this trie structure
type Trie struct {}
var globalVarListThing []string

func readThings() {
  path := "/home/ubuntu/goWorkSpace/src/github.com/EricRiney/GoFish/theData/smallWordsEn.txt"
  filePath, err := os.Open(path)
  if err != nil {
    panic(err)
  }

  ericScan := bufio.NewScanner(filePath)
  for ericScan.Scan() {
    //load into trie structure
    globalVarListThing = append(globalVarListThing,ericScan.Text())
  }
  fmt.Println("done reading maybe")
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
