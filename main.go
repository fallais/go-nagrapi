package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"strconv"

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
	Name string `json:"name,omitempty"`
	//modified to *int in order to have json correctly identify the 0 value as a real value and not nil
	CurrentState *int `json:"current_state,omitempty"`
}

var statusFile = flag.String("s", "/data/status.dat", "Specify the location of the status file")

//added the n flag to have as output only services with that level of issue
var nagiosstate = flag.Int("n", 0, "Specify the number Nagios uses to describe the status [0, 1, 2, 3]")

// GetState ...
func GetState(w http.ResponseWriter, r *http.Request) {
	// Parse the status file
	sdata := nagtomaps.ParseStatus(*statusFile)
	//needed to have a int value to check it against nagios status level
	nagstatus := *nagiosstate

	hostResponse := &Response{
		Hosts: nil,
	}

	for name := range sdata.Hoststatuslist {
		host := &Host{
			Name:     name,
			Services: nil,
		}

		for name2, object2 := range sdata.Servicestatuslist[name] {

			//if the state of the service does not match with the one in input interrupts the loop
			currstateint, _ := strconv.Atoi(object2["current_state"])
			if currstateint != nagstatus {
				continue
			}
			if object2["host_name"] == host.Name {

				service := &Service{
					Name:         name2,
					CurrentState: &currstateint,
				}

				host.Services = append(host.Services, service)

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
