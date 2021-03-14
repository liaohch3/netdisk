package main

import (
	"fmt"
	"net/http"
	"netdisk/meta"
)

func main() {
	initial()

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(fmt.Sprint("server start err: %v", err))
	}
}

func initial() {
	meta.InitFileMetas()
	initHandler()
}
