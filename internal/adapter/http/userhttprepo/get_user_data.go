package userhttprepo

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/model"
	"go-time-tracker/internal/utils"
)

// GetUserInfo get user info from http source
func (r Repository) GetUserInfo(passportNumber string) (*model.UserInfo, error) {
	url, err := r.BaseURL.Parse("/info")
	if err != nil {
		log.Error("Error parse base url", err)
		return nil, err
	}
	number, series, err := utils.ParsePassportNumber(passportNumber)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	url.RawQuery = "passportNumber=" + number + "&" + "passportSerie=" + series
	log.Info("Request User Info:", url)
	response, err := r.Get(url.String())
	if err != nil {
		log.Error("Error getting user info:", err)
		return nil, err
	}
	defer response.Body.Close()
	result := &model.UserInfo{}
	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Error("Decode User Info error:", err)
		return nil, err
	}
	log.Info("Response User Info:", result)
	return result, nil
}
