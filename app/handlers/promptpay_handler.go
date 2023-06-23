package handlers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strconv"

	"image/png"

	pp "github.com/Frontware/promptpay"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func GeneratePromptPayQR(c *fiber.Ctx) error {

	promptPayID := c.Query("PromptPayID")
	amountStr := c.Query("Amount")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		// Handle the error
		log.Fatal(err)
	}

	payment := pp.PromptPay{
		PromptPayID: promptPayID,
		Amount:      amount,
	}

	qrData, _ := payment.Gen()

	qrCode, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	qrImage, _, err := image.Decode(bytes.NewReader(qrCode))
	if err != nil {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, qrImage)
	if err != nil {
		log.Fatal(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	fmt.Println("QR code base64 string:")
	fmt.Println(base64Str)

	return c.SendString(base64Str)

}
