package main

import (
  "encoding/json"
  "net/http"
  "fmt"
)

func main() {
    http.HandleFunc("/", hello)

    http.HandleFunc("/job/", func(w http.ResponseWriter, r *http.Request) {

      data, err := queryJob()
      if err != nil {
        fmt.Println("This point broke it")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      w.Header().Set("Content-Type", "application/json; charset=utf-8")
      json.NewEncoder(w).Encode(data)
    })

    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}


type jobDatum struct {
    Query string `json:"query"`
    Total float64 `json:"totalResults"`
}


func queryJob() (jobDatum, error) {
    resp, err := http.Get("http://api.indeed.com/ads/apisearch?publisher=1439892112925001&q=ruby&l=boston%2C+ma&sort=&radius=&st=&jt=&start=&limit=&fromage=&filter=&latlong=1&co=us&chnl=&userip=1.2.3.4&useragent=Mozilla/%2F4.0%28Firefox%29&v=2&format=json")
    if err != nil {
      return jobDatum{}, err
    }

    defer resp.Body.Close()

    var d jobDatum

    if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
      fmt.Println("This point")
      return jobDatum{}, err
    }

    return d, nil
}
