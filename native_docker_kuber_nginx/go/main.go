package main

import (
	"fmt"
	"shell/db"
)

func main() {
	me := db.Storage{}
	fmt.Println(me.GetMe())
}
