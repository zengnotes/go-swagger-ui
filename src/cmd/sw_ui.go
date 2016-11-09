package main

import (
	"net/http"
	"swaggerui"
	"log"
	"os"
	"io/ioutil"
	"bytes"
)

func main() {
	//文件读取

	fi, err := os.Open("swagger.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	http.Handle("/", swaggerui.Handler("/", bytes.NewReader(fd)))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
