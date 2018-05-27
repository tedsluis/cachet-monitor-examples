/*
* Http (curl) request in golang
*/

package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "bytes"
)

type Status struct {
        status string `json: string`
}

func update_status(status string, component_id string){
        url := "http://cachet-node:80/api/v1/components/"
        req, _ := http.NewRequest("PUT", url + component_id, bytes.NewBuffer(bytes("{status: 3}")))
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("X-Cachet-Token", "ijHkUxeOPK9jO9k6eosr")
        req.Header.Add("cache-control", "no-cache")
        res, _ := http.DefaultClient.Do(req)
        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)
        fmt.Println(string(body))
        fmt.Println(res)
        fmt.Println(req)

}


func main() {
        update_status("testgroup", "bla") 

}
