package main

import (
	"ch2/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Brrr! %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Printf("Brrr! %v\n", tempconv.KToC(0))
}
