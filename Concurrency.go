package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"encoding/json"
)

type KanyeQuote struct {
	Msg string `json:"quote"` //not an apostrophe!
}

func main() {
	//waits for two trivial background tasks and http request to complete
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		proccess("Order")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		proccess("Refund")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		response, error := http.Get("https://api.kanye.rest")
		
		if error != nil { 
			fmt.Print("Error retrieving from API"); 
			wg.Done(); return
		}

		body, error := ioutil.ReadAll(response.Body) 
		if error != nil { 
			fmt.Print("Error reading response"); 
			wg.Done(); return
		}

		var quote KanyeQuote
		json.Unmarshal(body, &quote)
		
		fmt.Println("Quote: ", quote.Msg)
		wg.Done()
	}()

	wg.Wait()
}

func proccess(item string) {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processing", item, i)
	}
}