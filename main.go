package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type data struct {
	NumField            float64        `json:"numField"`
	ObjField            map[string]int `json:"objField"`
	ArrayField          []string       `json:"arrayField"`
	IntField            int64          `json:"intField"`
	BoolField           bool           `json:"boolField"`
	StringField         string         `json:"stringField"`
	StringDateField     string         `json:"stringDateField"`     // 1999-01-08 is January 8, 1999
	StringDateTimeField string         `json:"stringDateTimeField"` // 1985-04-12T23:20:50.52Z
	StringTimeField     string         `json:"stringTimeField"`     // [hour]:[minute]:[second].[subsecond]Z
}

func main() {
	count := flag.Int("count", 1, "number of documents to generate")
	flag.Parse()
	gofakeit.Seed(0)

	for idx := 1; idx <= *count; idx++ {
		var d data
		if err := gofakeit.Struct(&d); err != nil {
			log.Fatal(err)
		}

		d.IntField = int64(gofakeit.Number(-1000, 1000))
		d.NumField = gofakeit.Float64Range(-1000, 1000)
		d.StringDateField = gofakeit.Date().Format("2006-01-02")
		d.StringDateTimeField = gofakeit.Date().Format(time.RFC3339)
		d.StringTimeField = gofakeit.Date().Format("15:04:05") + "." + fmt.Sprintf("%d", gofakeit.Number(10, 50)) + "Z"

		b, err := json.Marshal(&d)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}
}