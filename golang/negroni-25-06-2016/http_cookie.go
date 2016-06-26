package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func main(){
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
        cookie, err := req.Cookie("my-cookie")
        if err != nil {
            http.SetCookie(res, &http.Cookie{
                Name: "my-cookie",
                Value: "1",
            })
        } else {
            numval, _ := strconv.ParseInt(cookie.Value, 10, 0)
            
            //? 
                    
            http.SetCookie(res, &http.Cookie{
                Name: "my-cookie",
                Value: fmt.Sprintf("%v", numval+1),
            })  

            fmt.Fprintf(res, "Read cookie - %v", cookie.String() )      
        }
    })
    http.ListenAndServe(":9000", nil)
}