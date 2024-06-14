package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os/exec"
    "time"
)

func main() {
    for {
        client := http.Client{
            Timeout: 3 * time.Second,
        }
        resp, err := client.Get("http://gstatic.com/generate_204")
        
        if err != nil || resp.StatusCode != http.StatusOK {
            fmt.Println("No response or error detected. Restarting data service...")
            
            cmdDisable := exec.Command("svc", "data", "disable")
            cmdDisable.Run()
            time.Sleep(2 * time.Second)
            
            cmdEnable := exec.Command("svc", "data", "enable")
            cmdEnable.Run()
        } else {
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)
            fmt.Printf("wget: response %s\n", body)
        }
        
        time.Sleep(5 * time.Second)
    }
}
