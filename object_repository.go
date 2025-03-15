package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

type Object struct {
	Oid           string  `json:"oid"`
	Ndethist      string  `json:"ndethist"`
	Ncovhist      float64 `json:"ncovhist"`
	Mjdstarthist  float64 `json:"mjdstarthist"`
	Mjdendhist    float64 `json:"mjdendhist"`
	Corrected     bool    `json:"corrected"`
	Stellar       bool    `json:"stellar"`
	Ndet          float64 `json:"ndet"`
	G_r_max       float64 `json:"g_r_max"`
	G_r_max_corr  float64 `json:"g_r_max_corr"`
	G_r_mean      float64 `json:"g_r_mean"`
	G_r_mean_corr float64 `json:"g_r_mean_corr"`
	Firstmjd      float64 `json:"firstmjd"`
	Lastmjd       float64 `json:"lastmjd"`
	Deltajd       float64 `json:"deltajd"`
	Meanra        float64 `json:"meanra"`
	Meandec       float64 `json:"meandec"`
	Sigmara       float64 `json:"sigmara"`
	Sigmadec      float64 `json:"sigmadec"`
	Step_id_corr  string  `json:"step_id_corr"`
}

type ObjectRepository interface {
	GetObject(objectId string) (Object, error)
}

type ALeRCEObjectRepository struct{}

func (a ALeRCEObjectRepository) GetObject(objectId string) (Object, error) {
	url := fmt.Sprintf("https://api.alerce.online/objects/%s", objectId)
	result, err := http.Get(url)
	if err != nil {
		return Object{}, err
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		if result.StatusCode == http.StatusNotFound {
			return Object{}, fmt.Errorf("object not found: %s", objectId)
		}
		return Object{}, fmt.Errorf("unsuccessful status code: %d. %v", result.StatusCode, result.Body)
	}

	body, err := io.ReadAll(result.Body)
	var data Object
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Object{}, err
	}
	return data, nil
}

type MockObjectRepository struct{}

func (m MockObjectRepository) GetObject(objectId string) (Object, error) {
	var object Object
	err := gofakeit.Struct(&object)
	if err != nil {
		return Object{}, err
	}
	object.Meanra = gofakeit.Latitude()
	object.Meandec = gofakeit.Longitude()
	object.G_r_mean = gofakeit.Float64Range(10, 25)
	return object, nil
}
