package main

import "fmt"

func main(){
	website := map[string]map[string]string {
		"Google":  {
			"name":"Google",
			"type":"Search",
		},
		"YouTube":  {
			"name":"YouTube",
			"type":"video",
		},
	}

	fmt.Println(website["Google"]["name"])
	fmt.Println(website["Google"]["type"])
}