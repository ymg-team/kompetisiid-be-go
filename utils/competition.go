package utils

import (
	b64 "encoding/base64"
	"strconv"
	"strings"
)

/**
 * function to encrypt id kompetisi
 * double base64 encode, and remove =
 */
func EncCompetitionId(competition_id int) string {
	competition_id_string := strconv.Itoa(competition_id)
	enc_competition_id := b64.StdEncoding.EncodeToString([]byte(competition_id_string))
	enc_competition_id = b64.StdEncoding.EncodeToString([]byte(enc_competition_id))
	enc_competition_id = strings.Replace(enc_competition_id, "=", "", 3)
	return enc_competition_id
}

/**
 * function to decrypt id kompetisi
 * double base64 decode
 */
func DecCompetitionId(enc_competition_id string) int {
	competition_id, _ := b64.StdEncoding.DecodeString(enc_competition_id)
	competition_id, _ = b64.StdEncoding.DecodeString(string(competition_id))
	competition_id_number, _ := strconv.ParseInt(string(competition_id), 10, 64)

	return int(competition_id_number)
}
