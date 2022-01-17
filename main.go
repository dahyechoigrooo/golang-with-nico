package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 요청에 대한 결과값을 구조체를 만들어 담는다.
type RequestResult struct {
	url    string
	status string
}

// 에러 메세지를 변수에 담아 재활용할 수 있게 만든다.
var errRequestFailed = errors.New("Request Failed")

func main() {
	// 결과 값을 map 함수로 만들어 key-value 형태의 변수를 만든다.
	results := make(map[string]string)
	// channel을 사용하여 메세지를 담을 변수를 만든다.
	c := make(chan RequestResult)

	// URL Checker에서 사용할 URL을 배열에 담는다.
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// 배열 안에 담긴 URL을 반복문으로 돌린다. 이때 URL은 goroutines를 사용하여 동시에 코드를 실행한다.
	for _, url := range urls {
		go hitURL(url, c)
	}

	// 메세지는 goroutines의 개수만큼 받아야한다. 일일이 개수를 세지 않고 URL의 개수만큼 반복문을 돌려 메세지를 받는다.
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	// URL Checking 후의 결과 값을 url, status 형태로 출력한다.
	for url, status := range results {
		fmt.Println(url, status)
	}
}

// RequestResult 라는 이름으로 메세지를 받는다(channel)
func hitURL(url string, c chan<- RequestResult) {
	resp, err := http.Get(url)
	status := "OK"
	// status 기본값이 'OK'이고 에러가 났을때만 'FAILED'를 출력한다.
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		status = "FAILED"
	}
	// status가 'OK'라면 key-value 형태의 데이터를 메시지로 메인에 전달한다.
	c <- RequestResult{url: url, status: status}
}
