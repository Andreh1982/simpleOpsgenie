package controllers

import (
	"encoding/json"
	"fmt"
	"simpleOpsgenie/handlers"
	"simpleOpsgenie/models"
)

var apiUrl string = "https://api.opsgenie.com/v1/"
var apiUrlString string

func GetIncidentList(status string) {

	var respPayload models.PayloadListMirror
	method := "GET"

	if status == "closed" {
		apiUrlString = apiUrl + "incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"
	} else if status == "resolved" {
		apiUrlString = apiUrl + "incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"
	} else if status == "opened" {
		apiUrlString = apiUrl + "incidents?query=status%3Aopen&offset=0&limit=200&sort=createdAt&order=desc"
	}
	bodyBytes := handlers.HandlerSingle(method, apiUrlString)

	if status == "opened" {
		json.Unmarshal(bodyBytes, &respPayload)
		fmt.Println("Incidentes Abertos:", respPayload.TotalCount)
	} else if status == "resolved" {
		json.Unmarshal(bodyBytes, &respPayload)
		fmt.Println("Incidentes Resolvidos:", respPayload.TotalCount)
	} else if status == "closed" {
		json.Unmarshal(bodyBytes, &respPayload)
		fmt.Println("Incidentes Fechados:", respPayload.TotalCount)
	}
}

func GetOneIncident(incidentID string) {

	var responsePayload models.PayloadUnitMirror
	method := "GET"
	apiUrlString = apiUrl + "incidents/" + incidentID + "?identifierType=tiny"
	bodyBytes := handlers.HandlerSingle(method, apiUrlString)

	json.Unmarshal(bodyBytes, &responsePayload)
	prettyJson, _ := json.MarshalIndent(responsePayload, "", "\t")
	fmt.Println(string(prettyJson))
}

func GetIdFromAll(status string) {

	var responsePayload models.PayloadListMirror
	method := "GET"

	if status == "closed" {
		apiUrlString = apiUrl + "incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"
	} else if status == "resolved" {
		apiUrlString = apiUrl + "incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"
	} else if status == "opened" {
		apiUrlString = apiUrl + "incidents?query=status%3Aopen&offset=0&limit=200&sort=createdAt&order=desc"
	}
	bodyBytes := handlers.HandlerListID(method, apiUrlString)

	json.Unmarshal(bodyBytes, &responsePayload)
	total := len(responsePayload.Data)

	fmt.Println("Incidents "+status, total)

	for i := 0; i < total; i++ {
		idJson, _ := json.MarshalIndent(responsePayload.Data[i].TinyID, "", "\t")
		createdAtJson, _ := json.MarshalIndent(responsePayload.Data[i].CreatedAt, "", "\t")
		fmt.Println(string(idJson), string(createdAtJson))
	}
}

func CreateIncident() {

	fmt.Println("Criando incidente...")

	var apiUrl string
	var c models.CreateIncident
	var incidentNumber string
	var responders models.Responders

	method := "POST"
	apiUrl = "https://api.opsgenie.com/v1/incidents/create"

	incidentNumber = "#07"

	responders.ID = "ebc8157f-e43c-478c-ae41-4b05d0682e22"
	responders.Type = "team"
	responders.Name = "OPS"

	c.Message = "Novo incidente de teste " + incidentNumber
	c.Description = "Incidente criado via simpleOpsgenie"
	c.Responders = append(c.Responders, responders)
	c.Tags = append(c.Tags, "stg")
	c.Tags = append(c.Tags, "staging")
	c.Details.Key1 = "Detalhes key 01 lorem ipsun lorem ipsun lorem ipsun."
	c.Details.Key2 = "Mais detalhes key 02 lorem ipsun lorem ipsun lorem ipsun"
	c.Priority = "P1"
	c.ImpactedServices = append(c.ImpactedServices, "86f95b45-a110-4d1b-9982-050e7ad087e1")
	c.ImpactedServices = append(c.ImpactedServices, "04e744c8-7556-4e46-8342-f68dbb2a0c35")
	c.StatusPageEntry.Title = "Incidente " + incidentNumber
	c.StatusPageEntry.Detail = "Detalhes do Incidente " + incidentNumber
	c.NotifyStakeholders = false

	fmt.Println(c)

	handlers.HandlerCreate(c, method, apiUrl)
}
