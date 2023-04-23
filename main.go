package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// 	return http.ErrUseLastResponse
		// },
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get("http://onet.pl/")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()
	// fmt.Println(resp.Header.Get("Host"))
	// location, err := resp.Location()
	// if err != nil {
	// 	log.Fatalf("Error: %s\n", err)
	// }
	// fmt.Printf("Location header: %s\n", location)
	fmt.Printf("Status code: %d\n", resp.StatusCode)
	// body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// fmt.Printf("StatusCode %d\n", resp.Request.Response.StatusCode)
	for i, s := range resp.Header {
		fmt.Println(i, s)
	}
}
