package basics

import (
	"fmt"
	"net/http"
)

func print_import() {
	fmt.Println("Hello, Go Library")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
	fmt.Println(resp.Body)
}
