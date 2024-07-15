package utils

import (
	"encoding/json"
	"ki-be/configs"
	models "ki-be/models/data"
	"strings"
)

var envConfig configs.Env = configs.EnvConf()

var cloudinaryCloudName string = envConfig.CloudinaryCloudName
var cloudinaryURL string = "https://res.cloudinary.com/" + cloudinaryCloudName + "/image/upload/kompetisi-id"

/**
* function to normalize competition image from db to ki cdn / cloudinary
 */
func ImageCompetitionNormalizer(image string, imageCloudinary string) models.ImageModel {
	var imgObj models.ImageModel

	// check is available poster from cloudinary
	if imageCloudinary != "" {
		json.Unmarshal([]byte(imageCloudinary), &imgObj)
	} else {
		// check is poster from migration server
		json.Unmarshal([]byte(image), &imgObj)

		if !strings.Contains(imgObj.Original, "http") {
			imgObj.Original = strings.Replace(cloudinaryURL+imgObj.Original, "/poster", "/competition", -1)
			imgObj.Small = strings.Replace(cloudinaryURL+imgObj.Small, "/poster", "/competition", -1)
		}

	}

	return imgObj
}

/**
* function to normalize news image from db to ki cdn / cludinary
 */
func ImageNewsNormalizer(image string, imageCloudinary string) models.ImageModel {
	var imgObj models.ImageModel

	// check is available image from cloudinary

	if imageCloudinary != "" {
		json.Unmarshal([]byte(imageCloudinary), &imgObj)
	} else {
		// parsing string to object
		json.Unmarshal([]byte(image), &imgObj)

		// check is image from migrations server
		if !strings.Contains(imgObj.Original, "http") {
			imgObj.Original = cloudinaryURL + imgObj.Original
			imgObj.Small = cloudinaryURL + imgObj.Small
		}
	}

	return imgObj
}
