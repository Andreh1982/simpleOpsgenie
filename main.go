package main

import (
	"flag"
	"fmt"
	"os"
	"simpleOpsgenie/controllers"
)

func main() {

	ListIncidentsVar := flag.String("list", "", "List Incidents by Status")
	GetIncidentVar := flag.String("get", "", "ID from Incident to Retrieve")
	ListIncidentsIDVar := flag.String("listid", "", "List Incidents ID by Status")
	CreateIncidentVar := flag.String("create", "", "Create a Incident")

	fmt.Println("Querying the API...")

	flag.Parse()

	if *ListIncidentsVar != "" {
		ListIncidents(ListIncidentsVar)
	}
	if *GetIncidentVar != "" {
		GetIncident(GetIncidentVar)
	}
	if *ListIncidentsIDVar != "" {
		ListIncidentsID(ListIncidentsIDVar)
	}
	if *CreateIncidentVar != "" {
		CreateIncident(CreateIncidentVar)
	}

}

func CreateIncident(CreateIncidentVar *string) {
	controllers.CreateIncident(CreateIncidentVar)
}

func GetIncident(GetIncidentVar *string) {
	if *GetIncidentVar != "" {
		controllers.GetOneIncident(*GetIncidentVar)
	}
}

func ListIncidents(ListIncidentsVar *string) {

	if *ListIncidentsVar == "opened" {
		controllers.GetIncidentList("opened")
	} else if *ListIncidentsVar == "resolved" {
		controllers.GetIncidentList("resolved")
	} else if *ListIncidentsVar == "closed" {
		controllers.GetIncidentList("closed")
	} else if *ListIncidentsVar == "" {
		os.Exit(0)
	}

}

func ListIncidentsID(ListIncidentsIDVar *string) {

	if *ListIncidentsIDVar == "opened" {
		controllers.GetIdFromAll("opened")
	} else if *ListIncidentsIDVar == "resolved" {
		controllers.GetIdFromAll("resolved")
	} else if *ListIncidentsIDVar == "closed" {
		controllers.GetIdFromAll("closed")
	} else if *ListIncidentsIDVar == "" {
		os.Exit(0)
	}

}
