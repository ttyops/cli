package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

type Service struct {
	Name *string `json:"name,omitempty"`
	ID   *string `json:"id,omitempty"`
}

type Services struct {
	Services []Service `json:"services"`
}

func services() Services {
	var services Services
	_, b := get("/service")
	json.Unmarshal(b, &services)
	return services
}

func listServices() {
	services := services()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)

	fmt.Fprintf(w, "%s\n", "NAME")

	for _, service := range services.Services {
		fmt.Fprintf(w, "%s\n",
			*service.Name)
	}

	w.Flush()
}

func service() {
	if len(os.Args) < 3 {
		usageService()
		os.Exit(0)
	}

	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listServices()
	} else {
		usageService()
	}
}
