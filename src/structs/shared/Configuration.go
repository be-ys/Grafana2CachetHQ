package shared

type Configuration struct {
	CachethqURL string                  `json:"cachethq_url"`
	Monitoring  ConfigurationMonitoring `json:"monitoring"`
	ServerPort  int                     `json:"server_port"`
	Sql         ConfigurationSQL        `json:"sql_configuration"`
}

type ConfigurationMonitoring struct {
	EnableScheduler bool   `json:"enable_scheduler"`
	Url             string `json:"url"`
	ServiceName     string `json:"service_name"`
	CachethqToken   string `json:"cachethq_token"`
}

type ConfigurationSQL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}
