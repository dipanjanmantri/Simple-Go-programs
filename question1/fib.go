package main

import "fmt"

func main(){
for i:=1;i<=8;i++{

fmt.Println(fib(i-1))
}
}
func fib(n int) int{
if(n==0){
return 0
}else if(n==1){
return 1
}else{
return fib(n-1)+fib(n-2)
}
}