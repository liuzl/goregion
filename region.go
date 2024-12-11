package goregion

import (
	_ "embed"
	"encoding/json"
)

//go:embed slim-2.json
var jsonData []byte

type Region struct {
	Name        string `json:"name"`
	Alpha2      string `json:"alpha-2"`
	CountryCode string `json:"country-code"`
}

var regions map[string]*Region

func init() {
	regions = make(map[string]*Region)
	var data []Region
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}
	for _, r := range data {
		regions[r.Alpha2] = &r
	}
}

func FindRegionByAlpha2(alpha2 string) *Region {
	return regions[alpha2]
}

func GetAllRegions() map[string]*Region {
	return regions
}
