package utils

import (
	"context"
	models "ki-be/models/data"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// ref: https://cloudinary.com/documentation/go_image_and_video_upload
func UploadCloudinary(dir string, imageFile string) (error, models.ImageModel) {
	var ctx = context.Background()

	imageResult := models.ImageModel{}

	cld, _ := cloudinary.NewFromParams(envConfig.CloudinaryCloudName, envConfig.CloudinaryApiKey, envConfig.CloudinarySecretKey)
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: dir})

	if err != nil {
		return err, imageResult
	}

	// small image transformation
	i, err := cld.Image(strings.Replace(resp.URL, "http://res.cloudinary.com/dhjkktmal/image/upload/", "", -1))
	if err != nil {
		return err, imageResult
	}

	i.Transformation = "w_400"
	transformationUrl, err := i.String()

	imageResult = models.ImageModel{
		Original: resp.URL,
		Small:    transformationUrl,
	}

	// fmt.Println(resp, err)
	return err, imageResult
}
