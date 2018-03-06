package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

type transactionReport struct {
	OrderID   string  `json:"orderid"`
	TimeStamp string  `json:"timestamp"`
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Amount    int     `json:"amount"`
	Price     float64 `json:"price"`
	Purchase  float64 `json:"purchase"`
	Omzet     float64 `json:"omzet"`  // Amount * Price
	Profit    float64 `json:"profit"` // Omzet - Harga Beli * Jumlah
}

type salesSummary struct {
	PrintDate   string  `json:"printdate"`
	StartDate   string  `json:"startdate"`
	EndDate     string  `json:"enddate"`
	TotalSales  int     `json:"totalsales"`
	TotalAmount int     `json:"totalamount"`
	TotalOmzet  float64 `json:"totalomzet"`
	TotalProfit float64 `json:"totalprofit"`
}

type salesReport struct {
	Items   []transactionReport `json:"items"`
	Summary salesSummary        `json:"summary"`
}

// ExportSalesReport exports Laporan Penjualan in CSV format
func ExportSalesReport(c *gin.Context) {
	// Get params
	startdate := c.Query("startdate")
	enddate := c.Query("enddate")

	// Validate
	if !validateParams(c, startdate, "startdate") {
		return
	}
	if !validateParams(c, enddate, "enddate") {
		return
	}

	// Construct URL
	url := BaseURL + "/penjualan?startdate=" +
		startdate + "&enddate=" + enddate

	// Make a get request
	rs, err := http.Get(url)

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

	if rs.StatusCode != http.StatusOK {
		c.JSON(rs.StatusCode, gin.H{
			"message": bodyString,
		})
		return
	}

	fmt.Println(bodyString)

	var result salesReport
	json.Unmarshal([]byte(bodyString), &result)

	b := &bytes.Buffer{}
	W = csv.NewWriter(b)

	// Summary
	WriteCSV(9, []string{"LAPORAN PENJUALAN"})
	WriteCSV(9, []string{})
	WriteCSV(9, []string{"Tanggal Cetak", result.Summary.PrintDate})
	WriteCSV(9, []string{
		"Tanggal", result.Summary.StartDate + " - " + result.Summary.EndDate,
	})
	p := message.NewPrinter(message.MatchLanguage("en"))
	WriteCSV(9, []string{
		"Total Omzet",
		"Rp" + p.Sprint(int(math.Round(result.Summary.TotalOmzet))),
	})
	WriteCSV(9, []string{
		"Laba Kotor",
		"Rp" + p.Sprint(int(math.Round(result.Summary.TotalProfit))),
	})
	WriteCSV(9, []string{
		"Penjualan", strconv.Itoa(result.Summary.TotalSales),
	})
	WriteCSV(9, []string{
		"Total Barang", strconv.Itoa(result.Summary.TotalAmount),
	})
	WriteCSV(9, []string{})

	// Header
	WriteCSV(9, []string{
		"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah",
		"Harga Jual", "Total", "Harga Beli", "Laba",
	})

	// Content
	for _, value := range result.Items {
		var record []string
		record = append(record, value.OrderID)
		record = append(record, convertDate(value.TimeStamp))
		record = append(record, value.SKU)
		record = append(record, value.Name)
		record = append(record, strconv.Itoa(value.Amount))
		record = append(record, strconv.Itoa(int(math.Round(value.Price))))
		record = append(record, strconv.Itoa(int(math.Round(value.Omzet))))
		record = append(record, strconv.Itoa(int(math.Round(value.Purchase))))
		record = append(record, strconv.Itoa(int(math.Round(value.Profit))))
		WriteCSV(9, record)
	}

	W.Flush()
	if err := W.Error(); err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Description", "Laporan Penjualan")
	c.Header("Content-Disposition", "attachment; filename=penjualan.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}

func validateParams(c *gin.Context, param string, msg string) bool {
	if len(param) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing " + msg + " params",
		})
		return false
	}
	return true
}

func convertDate(timeStamp string) string {
	dt, _ := time.Parse("2006-01-02T15:04:05Z", timeStamp)
	return dt.Format("2006-01-02 15:04:05")
}
