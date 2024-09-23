package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetDisasterRecoveryClusterList() map[string]interface{} {
	mold, _ := ReadMoldFile()
	var baseurl = mold.MoldUrl
	params := []MoldParams{
		{"command": "getDisasterRecoveryClusterList"},
	}
	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	resp, err := http.Get(endUrl)

	if err != nil {
		log.Fatal("Failed to communicate with Mold. (getDisasterRecoveryClusterList): ", err)
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func GetListVirtualMachinesMetrics(params []MoldParams) map[string]interface{} {
	mold, _ := ReadMoldFile()
	var baseurl = mold.MoldUrl
	params1 := []MoldParams{
		{"command": "listVirtualMachinesMetrics"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	resp, err := http.Get(endUrl)

	if err != nil {
		log.Fatal("Failed to communicate with Mold. (listVirtualMachines): ", err)
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func GetListVolumesMetrics(params []MoldParams) map[string]interface{} {
	mold, _ := ReadMoldFile()
	var baseurl = mold.MoldUrl
	params1 := []MoldParams{
		{"command": "listVolumesMetrics"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	resp, err := http.Get(endUrl)

	if err != nil {
		log.Fatal("Failed to communicate with Mold. (listVolumesMetrics): ", err)
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}
