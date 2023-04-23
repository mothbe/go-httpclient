package lib

import (
	"log"
	"net/http"
	"time"
)

func GetStatusCode(url string) int {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(url)
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
	// log.Printf("Function Request: status code: %d for %s\n", resp.StatusCode, url)
	// body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// fmt.Printf("StatusCode %d\n", resp.Request.Response.StatusCode)
	// for i, s := range resp.Header {
	// 	fmt.Println(i, s)
	// }
	return resp.StatusCode

}
