package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/cxfcxf/nagtomaps"
	"github.com/zenazn/goji"
)

// Response ...
type Response struct {
	Hosts []*Host `json:"hosts,omitempty"`
}

// Host ...
type Host struct {
	Name     string     `json:"name,omitempty"`
	Services []*Service `json:"services,omitempty"`
}

// Service ...
type Service struct {
	Name         string `json:"name,omitempty"`
	CurrentState string `json:"current_state,omitempty"`
}

var statusFile = flag.String("s", "/data/status.dat", "Specify the location of the status file")

// GetState ...
func GetState(w http.ResponseWriter, r *http.Request) {
	// Parse the status file
	sdata := nagtomaps.ParseStatus(*statusFile)

	hostResponse := &Response{
		Hosts: nil,
	}

	for name := range sdata.Hoststatuslist {
		host := &Host{
			Name:     name,
			Services: nil,
		}

		for name2, object2 := range sdata.Servicestatuslist[name] {
			if object2["host_name"] == host.Name {
				service := &Service{
					Name:         name2,
					CurrentState: object2["current_state"],
				}
				if service.CurrentState != "0" {
				host.Services = append(host.Services, service)
				}
			}
		}

		hostResponse.Hosts = append(hostResponse.Hosts, host)
	}

	// Marshal the result
	response, err := json.Marshal(hostResponse)
	if err != nil {
		panic("Impossible to marshal")
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	// Parse the flags
	flag.Parse()

	// Routes for API
	goji.Get("/state", GetState)

	// Set the port
	flag.Set("bind", ":8080")

	// Start the server
	goji.Serve()
}
