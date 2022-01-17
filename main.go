package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractJobs := getPage(i)
		jobs = append(jobs, extractJobs...)
		fmt.Println(jobs)
	}
}

// 각각의 페이지의 URL주소9
func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting" + pageURL)
	resp, err := http.Get(pageURL)
	checkErr(err)
	checkCode(resp)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs

}

// 각 취업 공고의 정보를 스크래핑하는 함수(struct 반환)
func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".salary-snippet span").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

// 공백을 제거해주는 함수
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

// HTML 파싱해오는 함수
func getPages() int {
	pages := 0
	resp, err := http.Get(baseURL)
	checkErr(err)
	checkCode(resp)

	// 함수가 종료된 후에 닫아줘서 메모리가 새어나가는 것을 막는다.
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	fmt.Println(doc)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

// Goquery 도큐먼트에 대한 에러 검사
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err) // 프로그램 종료
	}
}

// Get에 대한 에러 검사
func checkCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed withStatus : ", resp.StatusCode)
	}
}
