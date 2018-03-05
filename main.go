package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Item type with SKU, Name and Total
type Item struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type Items struct {
	Items []Item `json:"items"`
}

type UntypedJson map[string][]interface{}

func main() {
	r := gin.Default()
	r.GET("/items", func(c *gin.Context) {
		// Make a get request
		rs, err := http.Get("http://0.0.0.0:3000/api/items")

		// Process response
		if err != nil {
			panic(err)
		}
		defer rs.Body.Close()

		bodyBytes, err := ioutil.ReadAll(rs.Body)
		if err != nil {
			panic(err)
		}

		bodyString := string(bodyBytes)

		var result Items
		json.Unmarshal([]byte(bodyString), &result)

		b := &bytes.Buffer{}
		w := csv.NewWriter(b)

		if err := w.Write([]string{"SKU", "Nama Item", "Jumlah Sekarang"}); err != nil {
			log.Fatalln("error writing record to csv: ", err)
		}

		for _, value := range result.Items {
			var record []string
			record = append(record, value.SKU)
			record = append(record, value.Name)
			record = append(record, strconv.Itoa(value.Total))
			if err := w.Write(record); err != nil {
				log.Fatalln("error writing record to csv: ", err)
			}
		}
		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		c.Header("Content-Description", "Catatan Jumlah Barang")
		c.Header("Content-Disposition", "attachment; filename=toko.csv")
		c.Data(http.StatusOK, "text/csv; charset=utf-8", b.Bytes())
	})
	r.Run() // Listen and serve on 0.0.0.0:8080
}
