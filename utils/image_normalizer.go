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
* function to normalize image from db to ki cdn / cloudinary
 */
func ImageNormalizer(image string, imageCloudinary string) models.ImageModel {
	var posterObj models.ImageModel

	// check is poster from cloudinary
	if imageCloudinary != "" {
		json.Unmarshal([]byte(imageCloudinary), &posterObj)
	} else {
		// check is poster from migration server
		json.Unmarshal([]byte(image), &posterObj)

		if !strings.Contains(posterObj.Original, "http") {
			posterObj.Original = strings.Replace(cloudinaryURL+posterObj.Original, "/poster", "/competition", -1)
			posterObj.Small = strings.Replace(cloudinaryURL+posterObj.Small, "/poster", "/competition", -1)
		}

	}

	return posterObj
}
