package helpers

import (
	"context"
	"mime/multipart"
	"mini_project/app/configs"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(fileHeader *multipart.FileHeader) (string, error) {

	file, _ := fileHeader.Open()

	ctx := context.Background()
	cldService, _ := cloudinary.NewFromURL(configs.CLOUDINARY_URL)
	resp, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
