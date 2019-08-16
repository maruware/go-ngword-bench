package main

import (
	"encoding/csv"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const sampleCsvURL = "http://monoroch.net/kinshi/kinshi.csv"

func fetchWords() ([]string, error) {
	c := http.Client{}
	res, err := c.Get(sampleCsvURL)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(transform.NewReader(res.Body, japanese.ShiftJIS.NewDecoder()))

	words := []string{}
	for {
		record, err := reader.Read() // 1行読み出す
		// fmt.Printf("%v\n", record)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		words = append(words, record[0])
	}
	return words, nil
}

func BenchmarkRegex(b *testing.B) {
	words, err := fetchWords()

	if err != nil {
		panic(err)
	}

	pattern := strings.Join(words, "|")
	r, _ := regexp.Compile(pattern)

	b.ResetTimer()

	text := "あいうえお"

	for i := 0; i < b.N; i++ {
		r.MatchString(text)
	}
}

func BenchmarkContains(b *testing.B) {
	words, err := fetchWords()

	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	text := "あいうえお"

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			strings.Contains(text, word)
		}
	}
}
