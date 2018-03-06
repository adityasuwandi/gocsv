package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Laporan Nilai Barang
type itemReport struct {
	SKU    string  `json:"sku"`
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
	Value  float64 `json:"value"`
}

type summary struct {
	PrintDate   string  `json:"printdate"`
	TotalSKU    int     `json:"totalsku"`
	TotalAmount int     `json:"totalamount"`
	TotalValue  float64 `json:"totalvalue"`
}

type report struct {
	Items   []itemReport `json:"items"`
	Summary summary      `json:"summary"`
}

// ExportItemReport exports Nilai Barang in CSV format
func ExportItemReport(c *gin.Context) {
	// Make a get request
	rs, err := http.Get("http://0.0.0.0:3000/api/nilaibarang")

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

	var result report
	json.Unmarshal([]byte(bodyString), &result)

	b := &bytes.Buffer{}
	W = csv.NewWriter(b)

	// Summary
	WriteCSV(5, []string{"LAPORAN NILAI BARANG"})
	WriteCSV(5, []string{})
	WriteCSV(5, []string{"Tanggal Cetak", result.Summary.PrintDate})
	WriteCSV(5, []string{"Jumlah SKU", strconv.Itoa(result.Summary.TotalSKU)})
	WriteCSV(5, []string{"Jumlah Total Barang", strconv.Itoa(result.Summary.TotalAmount)})
	WriteCSV(5, []string{"Total Nilai", "Rp" + strconv.Itoa(int(math.Round(result.Summary.TotalValue)))})
	WriteCSV(5, []string{})

	// Header
	WriteCSV(5, []string{
		"SKU",
		"Nama Item",
		"Jumlah",
		"Rata-Rata Harga Beli",
		"Total",
	})

	// Content
	for _, value := range result.Items {
		var record []string
		record = append(record, value.SKU)
		record = append(record, value.Name)
		record = append(record, strconv.Itoa(value.Amount))
		record = append(record, strconv.Itoa(int(math.Round(value.Price))))
		record = append(record, strconv.Itoa(int(math.Round(value.Value))))
		WriteCSV(5, record)
	}

	W.Flush()
	if err := W.Error(); err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Description", "Laporan Nilai Barang")
	c.Header("Content-Disposition", "attachment; filename=nilai_barang.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}
