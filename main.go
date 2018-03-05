package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Email string
	Open  int64
	Link  int64
}

type Contacts []Contact

func main() {
	r := gin.Default()
	r.POST("/ping", func(c *gin.Context) {
		contactsData := c.PostForm("data")

		var contacts Contacts
		json.Unmarshal([]byte(contactsData), &contacts)

		b := &bytes.Buffer{}
		w := csv.NewWriter(b)

		if err := w.Write([]string{"email", "opens"}); err != nil {
			log.Fatalln("error writing record to csv: ", err)
		}

		for _, contact := range contacts {
			var record []string
			record = append(record, contact.Email)
			record = append(record, strconv.FormatInt(contact.Open, 10))
			if err := w.Write(record); err != nil {
				log.Fatalln("error writing record to csv: ", err)
			}
		}
		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename=contacts.csv")
		c.Data(http.StatusOK, "text/csv", b.Bytes())
	})
	r.Run() // Listen and serve on 0.0.0.0:8080
}
