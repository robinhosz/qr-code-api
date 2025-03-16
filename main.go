package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	router := gin.Default()

	router.GET("/generate-qr", func(c *gin.Context) {
		text := c.Query("text")
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro 'text' é obrigatório"})
			return
		}

		qr, err := qrcode.Encode(text, qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar QR code"})
			return
		}

		qrBase64 := base64.StdEncoding.EncodeToString(qr)

		c.JSON(http.StatusOK, gin.H{
			"qr_code": qrBase64,
		})
	})

	router.GET("/download-qr", func(c *gin.Context) {
		text := c.Query("text")
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro 'text' é obrigatório"})
			return
		}

		qr, err := qrcode.Encode(text, qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar QR code"})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename=qrcode.png")
		c.Data(http.StatusOK, "image/png", qr)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(router.Run(":" + port))
}