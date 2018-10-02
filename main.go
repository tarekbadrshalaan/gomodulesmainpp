package main

import (
	"fmt"

	"github.com/tarekbadrshalaan/gomodulesdepone"
	"github.com/tarekbadrshalaan/gomodulesdeptwo"
)

func main() {
	fmt.Println("Hi from Main app")

	fmt.Println(gomodulesdepone.GetData())

	//using new feature in gomodulesdepone
	fmt.Println(gomodulesdepone.GetExtraData())

	//using gomodulesdeptwo with calling gomodulesdepone/v2
	fmt.Println(gomodulesdeptwo.GetDatadepone())
}
