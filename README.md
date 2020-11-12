# Grafana2CachetHQ

Grafana2CachetHQ is a simple tool, written in Golang, which allow you to update CachetHQ status page from Grafana alerting (through webhooks).

# Introduction

## Functionalities
* Authentication through CachetHQ Token
* Flexibility, reconciliating Grafana graphs & CachetHQ services through tags
* Monitoring-callback
* Lightweight
* Easy configuration
* Docker-native

## How does it works ?
* Grafana receives data/metrics from various sources (Prometheus, influxDB, etc).
* These metrics are monitored through some graphs and alerts
* A Grafana Notification Channel send all alerts and resolutions to our tool
* The tool asks CachetHQ for services list, and retrieve concerned-ones
* Then, the tool send to CachetHQ the update.


# Setting things up

## Prerequisites
* A CachetHQ server
* A Grafana server
* A MariaDB server

## Configuration file
The config file is pretty simple. You have to put it at the root base of the project / in the same folder as the compiled project. A copy of this file is also provided when you clone the repo.

```json
{
	"cachethq_url": "http://status.my_awesome_company.net/",
	"monitoring": {
		"enable_scheduler": true,
		"url": "http://monitoring.private.my_awesome_company.local/",
		"service_name": "monitoring",
		"cachethq_token": "Your cachet token goes here"
	},
	"server_port": 8080,
	"sql_configuration": {
		"username": "root",
		"password": "root",
		"host": "localhost",
		"port": 3306,
		"database": "cachet"
	}
}
```

Detail of the config file :
* `cachethq_url` : The URL of your cachet server.
* `monitoring` : This part is for monitoring-callback : ask Grafana2CachetHQ to ensure that Grafana is running.
  * `enable_scheduler` : Enable or disable monitoring-callback. If you want to disable it, **please keep the other variables on the file (just empty them)**
  * `url` : The URL of your Grafana server
  * `service_name` : the name of the corresponding service in CachetHQ
  * `cachethq_token` : Your CachetHQ token.
* `server_port` : the port where Grafana2CachetHQ start. If you are deploying through Docker, you'll better not touch to this.
* `sql_configuration` : The SQL configuration. I think this part is obvious to understand... :)

## Configuring your Grafana
* First of all, you have to create a notification Channel :
  * Kind : `webhook`
  * Url : `Your Grafana2CachetHQ URL`
  * Optional Webhook settings :
     * HTTP Method : `POST̀`
     * Username : `Whatever you want here, we are not using it`
     * Password : `Your CachetHQ Token` 
  * Notification settings :
     * Default : `Yes`
     * Include image : `No`
     * Disable resolve message : `no`
     * Send reminders : `No`

* Then, we have to update our grafana rules (only ones that must update CachetHQ) :
  * In `Alert` tab, add some tags : 
    * Tag Name : the name of the CachetHQ impacted service (tags)
    * Tag value : The criticality level :
      * `critical` => Service will appear with "Major disruption" on CachetHQ
      * `normal` => Service will appear with "Partial disruption" on CachetHQ
      * ̀`minimal` => Service will appear with "Performance issue" on CachetHQ

## CachetHQ Configuration
* Create all your wanted components, and add keywords on them. The keywords must be the same as the Tag Name in Grafana Alert.

## Run your Grafana2CachetHQ server 

### From command line :
We are assuming that golang is installed on your system.

```sh
export GOPATH="$PWD"
go get github.com/gorrila/mux
go get github.com/go-sql-driver/mysql
go get github.com/zhashkevych/scheduler

go run ./src
```

### From Docker :
```sh
docker build -t "cachetconnector:latest" .
docker run -d --name="cachetconnector" -p 8080:8080 -v your_config_file_location:/app/config.json cachetconnector:latest
```
(Think to change `your_config_file_location` with the path of your config.json).

# Good to know

## Upcoming functionalities
* It will be great to use "incident" functionality on CachetHQ instead of updating the component. We are not doing this because we currently cannot create incident with more than one component impacted (see https://github.com/CachetHQ/Cachet/issues/1844). We are still looking on CachetHQ updates to add this functionality later.

## Dependencies used & acknowledgement
* [https://github.com/gorilla/mux](Gorilla/Mux) : A useful Golang lib to set up a Web server
* [https://github.com/go-sql-driver/mysql](go-sql-driver/mysql) : MySQL/MariaDB driver for Golang
* [https://github.com/zhashkevych/scheduler](zhashkevych/scheduler) : Goland scheduler (used for monitoring-callback).

* A special thanks to all people working on [https://github.com/CachetHQ/Cachet](CachetHQ) !

## Authors
Copyright (c) 2020 be ys group.

Crafted with ❤️ by [https://github.com/Artheriom](@Artheriom).

## License
This software is provided under MIT license. 


