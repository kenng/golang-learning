package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
    "strconv"
	"log"
)

var out_count int

func main() {
    files, _ := ioutil.ReadDir("./")
    var img string

    fmt.Println("length of files = ", len(files))
    for idx, f := range files {
      img += f.Name() + " "
      if (idx+1) % 4 == 0 {
        //fmt.Print(idx, ". ")
        //fmt.Println(f.Name())
        //fmt.Println(img)
        merge_img(img)
        img = "";
      }
    }

    if len(files) %4 != 0 {
      //fmt.Println(img)
	  merge_img(img)
    }
}

func merge_img(fn string) {
  var output string
  cmd := "montage -mode concatenate -tile 2x2 -border 5 -bordercolor white "+fn
  //cmd := "montage -mode concatenate -tile 2x *.png "
  output = " output" + strconv.Itoa(out_count) + ".png"
  out_count++;
  fmt.Println(cmd+output)
  _, err := exec.Command("sh","-c",cmd+output).Output()
  if err != nil {
	log.Fatal(err)
  }
  //fmt.Println(out)
}
