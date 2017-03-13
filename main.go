package main

import (
	"fmt"

	"github.com/spenczar/retool_test_lib"
)

const version = "v2"

func main() {
	_ = retool_test_lib.Exported
	fmt.Println(version)
}
