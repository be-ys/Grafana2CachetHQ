package services

import (
	"encoding/json"
	"helpers"
	"io/ioutil"
	"net/http"
	"sqlconnector"
	"strconv"
	"strings"
	"structs/shared"
	"time"
)

func UpdateService(token string, request shared.UpdateRequest) (int,string){

	//Retrieving configuration
	configuration := helpers.GetConfiguration()

	//Check if CachetHQ token provided is valid.
	//CachetHQ does not have /introspect or /check endpoint, so we are creating a bad request.
	res, _ := http.NewRequest("POST", configuration.CachethqURL+"/api/v1/components/", nil)
	res.Header.Add("Content-Type", "application/json")
	res.Header.Set("x-cachet-application", "Lunokhod")
	res.Header.Set("x-cachet-token", token)

	client := &http.Client{Timeout: time.Second * 20}
	resp, err := client.Do(res)
	if err != nil || resp.StatusCode != 400 {
		return http.StatusUnauthorized, "Invalid token provided"
	}



	//Retrieving list of components from CachetHQ.
	var cachetComponents shared.CachetPageable
	res, _ = http.NewRequest("GET", configuration.CachethqURL+"/api/v1/components?per_page=40000", nil)
	res.Header.Add("Content-Type", "application/json")

	client = &http.Client{Timeout: time.Second * 20}
	resp, err = client.Do(res)
	if err != nil || resp.StatusCode != 200 {
		return http.StatusBadRequest, "Error while updating component on CachetHQ"
	}

	body, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	_ = json.Unmarshal(body, &cachetComponents)


	//For each tag in Grafana request, select concerned services on CachetHQ to queue them for update
	for key, value := range request.Tags {
		var concernedServices []int
		for _, data := range cachetComponents.Data {
			for _, data2 := range data.Tags {
				if key == data2 {
					concernedServices = append(concernedServices, data.Id)
				}
			}
		}

		//Parsing criticality and state from Grafana to CachetHQ format
		status := 0
		criticality := 3 //Default criticality

		switch value {
			case "critical":
				criticality = 4
			case "normal":
				criticality = 3
			case "minimal":
				criticality = 2
		}

		switch request.State {
			case "ok":
				status = 1
			case "paused":
				status = criticality
			case "alerting":
				status = criticality
			case "no_data":
				status = criticality
			case "pending":
				status = 2
			default:
				return http.StatusBadRequest, "unknown status"
		}

		definitiveStatus := status
		var marshalled []byte

		if status==1 {
			//If received status from grafana is OK, delete alert from database and retrieve new criticality for components

			sqlconnector.DeleteAlert(request.RuleId)
			nb := sqlconnector.CountAlerts(key)

			if nb != 0 {
				definitiveStatus = sqlconnector.ReturnMaxCriticality(key)
			} else {
				definitiveStatus = 1
			}

			cachetBody := shared.CachetUpdate{Status: definitiveStatus}
			marshalled, _ = json.Marshal(cachetBody)

		} else {
			//If received status is not OK, insert alert into database and send new criticality for components
			sqlconnector.InsertAlert(key, request.RuleId, criticality)

			definitiveStatus := sqlconnector.ReturnMaxCriticality(key)

			cachetBody := shared.CachetUpdate{Status: definitiveStatus}
			marshalled, _ = json.Marshal(cachetBody)
		}

		//For each component, send update to CachetHQ
		for _, c := range concernedServices {
			readyToSend := strings.NewReader(string(marshalled))

			res, _ = http.NewRequest("PUT", configuration.CachethqURL+"/api/v1/components/"+strconv.Itoa(c), readyToSend)
			res.Header.Add("Content-Type", "application/json")
			res.Header.Set("x-cachet-application", "Lunokhod")
			res.Header.Set("x-cachet-token", token)

			client = &http.Client{Timeout: time.Second * 20}
			resp, err = client.Do(res)
			if err != nil || resp.StatusCode != 200 {
				return http.StatusBadRequest, "Error while updating component on CachetHQ"
			}
		}
	}

	return 0, ""
}
