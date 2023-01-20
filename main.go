package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	company  string
	category string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=프론트엔드&recruitPage="

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPagesNumber(baseURL)
	// totalPages := 2

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, pageC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := baseURL + strconv.Itoa(page) + "&recruitSort=relation&recruitPageCount=100"
	fmt.Println(pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	pageC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".job_tit>a").Text())
	company := cleanString(card.Find(".corp_name>a").Text())
	category := ""

	card.Find(".job_sector>a").Each(func(i int, card *goquery.Selection) {
		if i != 0 {
			category += ", "
		}
		category += card.Text()
	})

	category = cleanString(category)

	c <- extractedJob{
		id: id, title: title, company: company, category: category,
	}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Company", "Category"}

	wErr := w.Write(headers)
	checkErr(wErr)

	jobBaseUrl := "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx="

	for _, job := range jobs {
		jobSlice := []string{jobBaseUrl + job.id, job.title, job.company, job.category}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPagesNumber(url string) int {
	pages := 0
	res, err := http.Get(url)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(_ int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
