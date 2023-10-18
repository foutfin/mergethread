package main

import (
	"fmt"
)

func main() {
  // Change these variable to alter smaple size
  step := 100
  upto := 100000

  fmt.Println("Generating Graph .....")
  benchMark(upto,step)
  fmt.Println("Graph generated successfully")
}
