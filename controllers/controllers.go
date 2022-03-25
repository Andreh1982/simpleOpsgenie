package controllers

import (
	"encoding/json"
	"fmt"
	"simpleOpsgenie/models"
	"simpleOpsgenie/routes"
)

func GetIncidentList(status string) {

	var url string
	method := "GET"

	if status == "closed" {
		url = "https://api.opsgenie.com/v1/incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"

	} else if status == "resolved" {
		url = "https://api.opsgenie.com/v1/incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"

	} else if status == "opened" {
		url = "https://api.opsgenie.com/v1/incidents?query=status%3Aopen&offset=0&limit=200&sort=createdAt&order=desc"
	}

	bodyBytes := routes.Handler(method, url)

	var respPayload models.PayloadListMirror

	if status == "opened" {

		json.Unmarshal(bodyBytes, &respPayload)
		fmt.Println("Incidentes Abertos:", respPayload.TotalCount)
		// prettyJson01, _ := json.MarshalIndent(respPayload, "", "\t")
		// fmt.Println(string(prettyJson01))

	} else if status == "resolved" {

		json.Unmarshal(bodyBytes, &respPayload)
		fmt.Println("Incidentes Resolvidos:", respPayload.TotalCount)
		// prettyJson01, _ := json.MarshalIndent(respPayload, "", "\t")
		// fmt.Println(string(prettyJson01))

	} else if status == "closed" {

		json.Unmarshal(bodyBytes, &respPayload)
		fmt.Println("Incidentes Fechados:", respPayload.TotalCount)
		// prettyJson01, _ := json.MarshalIndent(respPayload, "", "\t")
		// fmt.Println(string(prettyJson01))

		// fmt.Println("Incidentes Resolvidos:", respPayload.Data[0].CreatedAt)

	}

}

func GetOneIncident(incidentID string) {

	var responsePayload models.PayloadUnitMirror

	method := "GET"
	baseUrl := "https://api.opsgenie.com/v1/incidents/" + incidentID + "?identifierType=tiny"
	fmt.Println(baseUrl)

	bodyBytes := routes.Handler(method, baseUrl)

	json.Unmarshal(bodyBytes, &responsePayload)
	prettyJson, _ := json.MarshalIndent(responsePayload, "", "\t")
	fmt.Println(string(prettyJson))

}

func GetIdFromAll(status string) {

	var responsePayload models.PayloadListMirror
	var url string

	method := "GET"

	if status == "closed" {
		url = "https://api.opsgenie.com/v1/incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"

	} else if status == "resolved" {
		url = "https://api.opsgenie.com/v1/incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"

	} else if status == "opened" {
		url = "https://api.opsgenie.com/v1/incidents?query=status%3Aopen&offset=0&limit=200&sort=createdAt&order=desc"
	}

	bodyBytes := routes.HandlerListID(method, url)

	json.Unmarshal(bodyBytes, &responsePayload)
	total := len(responsePayload.Data)

	fmt.Println("Incidentes Abertos:", total)

	for i := 0; i < total; i++ {

		prettyJson, _ := json.MarshalIndent(responsePayload.Data[i].TinyID, "", "\t")
		fmt.Println(string(prettyJson))

	}

}

func CreateIncident() {
	fmt.Println("Criando incidente...")

	var apiUrl string
	var c models.CreateIncident

	method := "POST"
	apiUrl = "https://api.opsgenie.com/v1/incidents/create"

	c.Message = "Incidente de teste 'Message'"
	c.Description = "Incidente criado via simpleOpsgenie"
	// c.Responders[0].ID = "ID do Responder"
	// c.Responders[0].Type = "Tipo de Responde?"
	// c.Responders[0].Name = "Fulano da Silva"
	// c.Tags[0] = "Test"
	c.Details.Key1 = "Key1"
	c.Details.Key2 = "Key2"
	c.Priority = "P1"
	// c.ImpactedServices[0] = "Meu Serviço Impactado"
	c.StatusPageEntry.Title = "Incidente #1"
	c.StatusPageEntry.Detail = "Detalhes do Incidente #1"
	// createPayload.StatusPageEntry.NotifyStakeholders = "true"

	routes.CreateIncidentHandler(c, method, apiUrl)

}
