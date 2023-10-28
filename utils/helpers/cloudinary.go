package helpers

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(fileHeader *multipart.FileHeader) (string, error) {

	file, _ := fileHeader.Open()

	ctx := context.Background()
	CLOUDINARY_URL := "cloudinary://747566867141386:C5PfCtdMbGhX-KUmtJyIL6KNG8Y@dvu15ohox"
	cldService, _ := cloudinary.NewFromURL(CLOUDINARY_URL)
	resp, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
