package main

import (
	"gocsv/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/nilaibarang", api.ExportItemReport)
	r.GET("/penjualan", api.ExportSalesReport)
	r.Run() // Listen and serve on 0.0.0.0:8080
}
