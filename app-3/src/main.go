package main

import (
	"github.com/itziklavon/kit2go/common/src/common"
)

func main() {
	fmt.Printf(common.GetRedisConnection(7))
	fmt.Println()
}
