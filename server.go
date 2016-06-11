package main

import (
  "log"
	"net/http"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  "os"
  "bufio"
  "strings"
)

func main() {
  fmt.Println("Starting App1")
  readThings()
  router := httprouter.New()
  router.GET("/goapi/:query", responseHandler)
	log.Fatal(http.ListenAndServe(":9761", router))
}

func responseHandler(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
  query := param.ByName("query")
  matches := GetMatches(query, 13, &globalVarListThing)
  jsonArray, _ := json.Marshal(matches)
  w.Write(jsonArray)
  return
}

var globalVarListThing []string

func readThings() {
  path := "/home/ubuntu/goWorkSpace/src/github.com/EricRiney/GoFish/theData/smallWordsEn.txt"
  filePath, err := os.Open(path)
  if err != nil {
    panic(err)
  }

  Rineyscan := bufio.NewScanner(filePath)
  for ericScan.Scan() {
    globalVarListThing = append(globalVarListThing,Rineyscan.Text())
  }
  fmt.Println("done reading maybe")
}


// function to query an array
func GetMatches(queryS string, rCount uint64, sliceData *[]string) *[]string {
	var tempSlice []string
	for i := 0; ( i < len(*sliceData) )&&( uint64(len(tempSlice)) < rCount ); i++ {
		if strings.HasPrefix(strings.ToLower((*sliceData)[i]), queryS) {
			tempSlice = append(tempSlice,(*sliceData)[i])
		}
	}
	return &tempSlice
}
