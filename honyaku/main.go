package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	var slice []string
	slice = strings.Split(string(b), ".")
	var auth string = "mytoken"
	client := new(http.Client)
	for i := 0; i < len(slice); i++ {
		s := slice[i] + "."
		if strings.Contains(s, "-") {
			s = strings.Replace(s, "-", "", -1)
		}

		// 短く分割する
		ss := strings.Split(s, ".")
		for j := 0; j < len(ss); j++ {
			url := "https://api-free.deepl.com/v2/translate?auth_key=" + auth + "&text=" + ss[j] + "." + "&source_lang=IT&target_lang=JA"
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				log.Fatal(err)
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			fmt.Println(resp)
		}
	}

}
