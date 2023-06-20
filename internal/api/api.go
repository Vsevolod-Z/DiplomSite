package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FindById(input string) models.AppMainData {
	fmt.Println(input)
	url := fmt.Sprintf("http://109.254.9.58:8080/api/apps/findByIds?appids=%s", input)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to send request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var apps []models.AppMainData

	err = json.Unmarshal(body, &apps)
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Println(apps[0])
	return apps[0]
}
