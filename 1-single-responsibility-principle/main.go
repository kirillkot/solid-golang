// go run main.go
// Stupid example for realize principle of Single responsibility
// bad part
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReportBad ...
type ReportBad struct {
	Title string
	Date  string
}

func (r *ReportBad) getTitle() string {
	return r.Title
}

func (r *ReportBad) getDate() string {
	return r.Date
}

func (r *ReportBad) getContents() (string, string) {
	return r.Title, r.Date
}

func (r *ReportBad) formatJSON() ([]byte, error) {
	answer, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("Method: formatJSON, err: %v", err)
	}

	return answer, nil
}

// good part

// ReportGood ...
type ReportGood struct {
	Title string
	Date  string
}

func (r *ReportGood) getTitle() string {
	return r.Title
}

func (r *ReportGood) getDate() string {
	return r.Date
}

func (r *ReportGood) getContents() (string, string) {
	return r.Title, r.Date
}

// ReportFormat ...
type ReportFormat struct {
}

func (rf *ReportFormat) formatJSON(r *ReportGood) ([]byte, error) {
	answer, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("Method: formatJSON, err: %v", err)
	}

	return answer, nil
}

func main() {
	repBad := &ReportBad{Date: "123", Title: "Bad"}
	ansBad, err := repBad.formatJSON()
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(ansBad)
	fmt.Println()

	repGood := &ReportGood{Date: "321", Title: "Good"}
	ansGood, err := (&ReportFormat{}).formatJSON(repGood)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(ansGood)
	fmt.Println()
}
