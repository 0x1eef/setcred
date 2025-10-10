package main

import (
	"fmt"
	"github.com/0x1eef/bsd/setcred"
	"os"
)

func main() {
	err := setcred.SetCred(
		setcred.SetUid(uint32(1001)),
		setcred.SetGid(uint32(1001)),
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("euid: %d, egid: %d", os.Geteuid(), os.Getegid())
	}
}
