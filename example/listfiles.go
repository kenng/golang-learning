package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    files, _ := ioutil.ReadDir("./")
    for _, f := range files {
            fmt.Println(f.Name())
    }

    for idx, f := range files {
            fmt.Print(idx, ". ")
            fmt.Println(f.Name())
    }
}
