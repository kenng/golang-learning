package main

import (
  "fmt"
  "io/ioutil"
)

func main(){
  files, err = ioutil.ReadDir("./")
  for idx, f := range files {
    if ( 0 == idx % 4 ) {
      exe_cmd()
    }
  }

}
