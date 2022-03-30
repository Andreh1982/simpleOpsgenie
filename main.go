package main

import (
	"flag"
	"fmt"
	"os"
	"simpleOpsgenie/controllers"
	"simpleOpsgenie/models"
)

func main() {

	TotalIncidentsVar := flag.String("total", "", "Incidents Totals by Status (opened, resolved, closed)")
	GetIncidentVar := flag.String("get", "", "Incident Information (TinyID)")
	ListIncidentsIDVar := flag.String("list", "", "List Incidents ID by Status (opened, resolved, closed)")
	CreateIncidentVar := flag.String("create", "", "Create an Incident (Name or Number)")
	ResolveIncidentVar := flag.Bool("resolve", false, "Resolve an Incident (will prompt for wich one)")
	CloseIncidentVar := flag.Bool("close", false, "Close an Incident (will prompt for wich one)")
	DeleteIncidentVar := flag.Bool("delete", false, "Delete an Incident (will prompt for wich one)")

	if len(os.Args) < 2 {
		fmt.Println("# SimpleOpsgenie v1.0")
		fmt.Println("# Dont known what to do, exiting.")
		os.Exit(0)
	}

	flag.Parse()

	fmt.Println("# SimpleOpsgenie v1.0")
	fmt.Println("# Querying the API...")

	if *TotalIncidentsVar != "" {
		TotalIncidents(TotalIncidentsVar)
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
	if *ResolveIncidentVar {
		ResolveIncident(ResolveIncidentVar)
	}
	if *CloseIncidentVar {
		CloseIncident(CloseIncidentVar)
	}
	if *DeleteIncidentVar {
		DeleteIncident(DeleteIncidentVar)
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

func TotalIncidents(TotalIncidentsVar *string) {
	if *TotalIncidentsVar == "opened" {
		controllers.TotalIncidentList("opened")
	} else if *TotalIncidentsVar == "resolved" {
		controllers.TotalIncidentList("resolved")
	} else if *TotalIncidentsVar == "closed" {
		controllers.TotalIncidentList("closed")
	} else if *TotalIncidentsVar == "" {
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

func ResolveIncident(ResolveIncidentVar *bool) {
	controllers.ResolveIncident(models.PayloadUnitMirror{})
}

func CloseIncident(CloseIncidentVar *bool) {
	controllers.CloseIncident(models.PayloadUnitMirror{})
}

func DeleteIncident(DeleteIncidentVar *bool) {
	controllers.DeleteIncident(models.PayloadUnitMirror{})
}
