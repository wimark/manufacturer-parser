package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/wimark/vendormap"
)

//ManufData тип данных mac префикс и производитель оборудования
type ManufData struct {
	Prefix string `bson:"_id"`
	Vendor string `bson:"vendor"`
}

//ParseData переводит данные в формат ManufData
func getAndParseData() ([]ManufData, error) {
	data, err := getMACData(Conf.DATAURL)
	if err != nil {
		return gitInitData(), err
	}
	var manuf []ManufData
	dataString := string(data)
	splitDataByLine := strings.Split(dataString, "\n")
	if len(splitDataByLine) == 0 {
		return gitInitData(), nil
	}
	for _, v := range splitDataByLine {
		var m ManufData
		line := strings.Split(v, "\t")
		if len(line) == 2 {
			m.Prefix = line[0]
			m.Vendor = line[1]
			manuf = append(manuf, m)
		}
	}
	return manuf, nil
}

func getMACData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func gitInitData() []ManufData {

	m := make([]ManufData, 0, len(vendormap.ManufacturerMap))
	for k, v := range vendormap.ManufacturerMap {
		m = append(m, ManufData{Prefix: k, Vendor: v})
	}

	return m
}
