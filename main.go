package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/gin-gonic/gin"
	"image/png"
	"os"
	"regexp"
)

func main() {
	r := gin.Default()

	r.GET("/generate", func(c *gin.Context) {
		isValid, _ := regexp.MatchString("^[A-Za-z0-9]+$", c.Query("content"))
		if !isValid {
			c.AbortWithStatus(422)
			return
		}
		filepath := "./barcodes/" + c.Query("content") + ".png"
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			bc, err := code128.Encode(c.Query("content"))
			if err != nil {
				panic(err)
			}
			bc, _ = barcode.Scale(bc, 1000, 200)
			file, _ := os.Create(filepath)
			defer file.Close()
			png.Encode(file, bc)
		}
		c.File(filepath)
	})

	r.Run()
}
