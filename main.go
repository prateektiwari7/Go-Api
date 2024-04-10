package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://api.restful-api.dev/objects"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
