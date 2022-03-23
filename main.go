package main

import (
	"fmt"
	"simpleOpsgenie/controllers"
)

func main() {

	fmt.Println("Consultando a API...")

	// controllers.IncidentList("opened")
	// controllers.IncidentList("resolved")
	// controllers.IncidentList("closed")

	// controllers.GetOneIncident("2111")

	controllers.GetIdFromAll("opened")
	// controllers.GetIdFromAll("resolved")
	// controllers.GetIdFromAll("closed")

}
