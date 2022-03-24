package main

import (
	"fmt"
	"simpleOpsgenie/controllers"
)

func main() {

	fmt.Println("Consultando a API...")

	controllers.GetIncidentList("opened")
	controllers.GetIncidentList("resolved")
	controllers.GetIncidentList("closed")

	// controllers.GetOneIncident("2111")

	// controllers.GetIdFromAll("opened")
	// controllers.GetIdFromAll("resolved")
	// controllers.GetIdFromAll("closed")

}
