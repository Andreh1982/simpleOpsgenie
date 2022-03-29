package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"simpleOpsgenie/models"
)

func HandlerSingle(method string, url string) []byte {
	_, genieKey := InitEnv()
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", genieKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

func HandlerListID(method string, url string) []byte {
	_, genieKey := InitEnv()
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", genieKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

func HandlerCreate(c models.CreateIncident, method string, apiUrl string) {
	_, genieKey := InitEnv()
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, apiUrl, reader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", genieKey)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

func HandlerResolve(c models.PayloadUnitMirror, apiUrlString string) {
	_, genieKey := InitEnv()
	method := "PATCH"
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, apiUrlString, reader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", genieKey)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

func HandlerClose() {

}

func InitEnv() (apiUrl string, genieKey string) {
	file, err := os.Open("config.json")
	if err != nil {
		os.Exit(0)
	}
	decoder := json.NewDecoder(file)
	configuration := models.Configuration{}
	err1 := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err1)
	}
	apiUrl = configuration.ApiUrl
	genieKey = configuration.GenieKey
	return apiUrl, genieKey
}
