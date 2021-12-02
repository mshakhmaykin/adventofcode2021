package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

func get_day_inputs(year, day int) {
	session := "53616c7465645f5f7268814467ff246cf61295c6dd09368558bf74acaeb16ec583034d53643cdaec47ae21c231960973"
	var day_url string = fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	path_dir := "/Users/mshakhmaykin/work/test/GOLANG/adventofcode2021/inputs"
	os.MkdirAll(path_dir, 0755)
	var day_file string = fmt.Sprintf("%s/%d_%d.in", path_dir, year, day)

	f, err := os.Open(day_file)
	if err == nil || os.IsExist(err) {
		f.Close()
		fmt.Println("Skipping inputs from", day_url, "as file", day_file, "already exists")
		return
	}

	fmt.Println("Getting inputs from", day_url)
	request, err := http.NewRequest("GET", day_url, nil)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	request.AddCookie(&http.Cookie{Name: "session", Value: session})

	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	resp, err := client.Do(request)

	//processerror(err)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	output_file, err := os.Create(day_file)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(output_file, string(body))
}

func main() {
	today_number := int(math.Min(float64(time.Now().Day()), 25))
	year := 2021
	for day := 1; day <= today_number; day++ { // getting all inputs to date
		get_day_inputs(year, day)
	}
}
