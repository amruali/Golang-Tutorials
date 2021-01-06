package main

import "fmt"

func CalculateSumFrom1ToN(n int)(result int){
	result = n * ( n + 1 ) / 2
	return result
}


func main(){
	res := CalculateSumFrom1ToN(50)
	fmt.Println(res)
}
