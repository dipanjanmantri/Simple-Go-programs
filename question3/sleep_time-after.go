package main

import (
"time"
"fmt"
)

func main(){
x:=4
sleep(x)
var input string
fmt.Scanln(&input)
}

func sleep(sec int) {
	select {
		case <- time.After(time.Duration(sec) * time.Second):
			return
	}
}