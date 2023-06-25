package utils

import (
	"encoding/json"
	"ki-be/configs"
	models "ki-be/models/data"
)

var envConfig configs.Env = configs.EnvConf()

/**
* function to normalize image from db to ki cdn / cloudinary
 */
func ImageNormalizer(image string) models.ImageModel {
	var posterObj models.ImageModel
	json.Unmarshal([]byte(image), &posterObj)

	posterObj.Small = posterObj.Original

	return posterObj
}
