package main

import (
  "sync"
)

func main(){
    wg := new(sync.WaitGroup)
    wg.Add(3)

    x := []string{"echo newline >> foo.o", "echo newline >> f1.o", "ls ./"}
    go exe_cmd(x[0], wg)
    go exe_cmd(x[1], wg)
    go exe_cmd(x[2], wg)

    wg.Wait()
}
