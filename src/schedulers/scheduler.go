package schedulers

import (
	"context"
	"helpers"
	"net/http"
	"services"
	"structs/shared"
	"time"
)

func Scheduler(ctx context.Context){
	configuration := helpers.GetConfiguration()

	//Check if Monitoring callback is up : GET on URL, check if status code is == 200
	res, _ := http.NewRequest("GET", configuration.Monitoring.Url, nil)
	client := &http.Client{Timeout: time.Second * 20}
	resp, err := client.Do(res)

	state := "ok"
	if err != nil || resp.StatusCode != 200 {
		state = "alerting"
	}

	services.UpdateService(configuration.Monitoring.CachethqToken, shared.UpdateRequest{State: state, RuleId: -1, Tags: map[string]string{ configuration.Monitoring.ServiceName: "critical"}})
}
