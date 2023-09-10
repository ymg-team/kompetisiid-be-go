package repositories

import (
	tableModels "ki-be/models/tables"
	storageDb "ki-be/storages/db"
)

/**
* function to get user data by userkey
 */
func GetUserByUserKey(userKey string) (error, tableModels.User) {
	db := storageDb.ConnectDB()
	resultData := tableModels.User{}
	query := db.Select(`user.id_user, user.username, user.level`).Where("user.user_key = ?", userKey)
	query.First(&resultData)
	return nil, resultData
}
