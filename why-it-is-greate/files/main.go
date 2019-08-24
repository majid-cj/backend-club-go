package main
import (
	"os"
	"fmt"
	 "io/ioutil"
	)

func main() { 

	file, err := os.Create("why-it-is-greate/files/test.txt") 
	if err != nil {
		fmt.Printf("%v", err)
	} 
	defer file.Close()
	file.WriteString("test")

	bs, err := ioutil.ReadFile("why-it-is-greate/files/test.txt") 
	if err != nil { 
		fmt.Printf("%v", err)
	} 
	str := string(bs) 
	fmt.Println(str)
}
