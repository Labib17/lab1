package main

import (
  "encoding/json"
  "html/template"
  "net/http"
  "path"
  "time"
)

//Time ...
type Time struct {
  Hour    int `json:"hour"`
  Minutes int `json:"min"`
  Second  int `json:"sec"`
}

func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  http.HandleFunc("/", mainPage)
  http.HandleFunc("/online", online)
  http.HandleFunc("/time", timePage)
  http.ListenAndServe(":8795", nil)
}

func timePage(w http.ResponseWriter, r *http.Request) {
  p := time.Now()
  time := Time{p.Hour(), p.Minute(), p.Second()}
  JSON, _ := json.Marshal(time)
  fp := path.Join("server-time.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if err := tmpl.Execute(w, time); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  w.Write([]byte("<ul class=" + "list-group list-group-horizontal" + ">"))
  w.Write([]byte("<li class=" + "list-group-item" + ">"))
  w.Write([]byte(JSON))
  w.Write([]byte("</li>"))
  w.Write([]byte("</ul>"))
}
func mainPage(w http.ResponseWriter, r *http.Request) {
  fp := path.Join("main.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func online(w http.ResponseWriter, r *http.Request) {
  p := time.Now()
  time := Time{p.Hour(), p.Minute(), p.Second()}
  fp := path.Join("online-time.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if err := tmpl.Execute(w, time); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
