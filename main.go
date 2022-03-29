package main

import (
	"flag"
	"fmt"
	"os"
	"simpleOpsgenie/controllers"
	"simpleOpsgenie/models"
)

func main() {
	ListIncidentsVar := flag.String("list", "", "List Incidents by Status(opened, resolved, closed)")
	GetIncidentVar := flag.String("get", "", "Tiny ID from Incident to Retrieve")
	ListIncidentsIDVar := flag.String("listid", "", "List Incidents ID by Status(opened, resolved, closed")
	CreateIncidentVar := flag.String("create", "", "Create a Incident Indented by Number")
	ResolveIncidentVar := flag.String("resolve", "", "Resolve a Incident(will prompt for wich one)")
	CloseIncidentVar := flag.String("close", "", "Close a Incident(will prompt for wich one)")

	flag.Parse()

	fmt.Println("Querying the API...")

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
	if *ResolveIncidentVar != "" {
		ResolveIncident(ResolveIncidentVar)
	}
	if *CloseIncidentVar != "" {
		CloseIncident(CloseIncidentVar)
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

func ResolveIncident(ResolveIncidentVar *string) {
	controllers.ResolveIncident(models.PayloadUnitMirror{})
}

func CloseIncident(ResolveIncidentVar *string) {
	controllers.CloseIncident(models.PayloadUnitMirror{})
}
