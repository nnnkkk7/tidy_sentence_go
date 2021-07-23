package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(`./output.txt`)
	if err != nil {
		log.Fatal(err) //ファイルが開けなかったときエラー出力
	}
	defer file.Close()

	var slice []string
	slice = strings.Split(string(b), ".")
	for i := 0; i < len(slice); i++ {
		s := slice[i] + "."
		if strings.Contains(s, "-") {
			s = strings.Replace(s, "-", "", -1)
		}

		fmt.Println(s)
		s += "\n"
		file.Write([]byte(s))
	}

}
