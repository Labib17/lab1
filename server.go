package main
import (
  "time"
  "encoding/json"
  "log"
  "net/http"
)

func main() {
  type Time struct{
    Time string `json:"time"`
  }

  http.HandleFunc("/time", func(w http.ResponseWriter, req *http.Request) {
    test := &Time{Time: time.Now().Format(time.RFC3339)} //current time
    res, _ := json.Marshal(test)
    w.Header().Set("content-type", "aplication/json")
    w.WriteHeader(200)
    w.Write(res)
  })

  log.Println("Starting HTTP server")
  log.Fatal(http.ListenAndServe(":8795", nil))  //http://localhost:8795/time
}
/ Client
/*
const response = await fetch('http://localhost:8795/time');
if (response.ok) {
  const res = await response.text();
  console.log(JSON.parse(res));
} else
  console.error('${respoonse.status}');
*/
