package services

import (
	"bytes"
	"context"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"go.mongodb.org/mongo-driver/bson"
)

func GeneratePDF(htmlContent string) ([]byte, error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func DecrementAPITokenUsage(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$inc": bson.M{"remaining_requests": -1},
	}
	filter := bson.M{"token": token, "remaining_requests": bson.M{"$gt": 0}}

	result := apiTokenCollection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
