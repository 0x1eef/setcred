package main

import (
	"fmt"
	"github.com/0x1eef/bsd/setcred"
	"os"
)

func main() {
	err := setcred.SetCred(
		setcred.SetUid(1001),
		setcred.SetGid(1001),
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("euid: %d, egid: %d\n", os.Geteuid(), os.Getegid())
	}
}
