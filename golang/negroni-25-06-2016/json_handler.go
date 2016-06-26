package main

import (
    "net/http"
    "encoding/json"
    "fmt"
)

func jsonHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    
    myItems := []string{"item1", "item2", "item3", "&"}
    a, _ := json.Marshal(myItems)
    w.Write(a)
    fmt.Println(string(a))
    return
}


func main() {
    http.HandleFunc("/", jsonHandler)
    //nil uses the default multiplexer
    http.ListenAndServe(":8080", nil)
}
