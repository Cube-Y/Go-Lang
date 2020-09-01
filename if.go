package main
import "fmt"

func main() {
	a,b := "10","100"
	Ab := "a is larger than b"
	aB := "a is smaller than b"
	AB := "a equals b"
	
	if a > b {
		fmt.Println(Ab)
	}else if a < b {
		fmt.Println(aB)
	}else {
		fmt.Println(AB)
	}
}