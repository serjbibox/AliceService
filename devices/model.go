package devices

type Parameter struct {
	Instance string `json:"instance" example:"temperature"`
	Unit     string `json:"unit" example:"unit.temperature.celsius"`
}

type Property struct {
	Type        string    `json:"type" example:"123"`
	Retrievable bool      `json:"retrievable"`
	Reportable  bool      `json:"reportable"`
	Parameters  Parameter `json:"parameters" example:"123"`
}

type Capability struct {
	Type        string    `json:"type" example:"123"`
	Retrievable bool      `json:"retrievable"`
	Reportable  bool      `json:"reportable"`
	Parameters  Parameter `json:"parameters" example:"123"`
}

type DeviceInfo struct {
	Manufacturer string `json:"manufacturer" example:"Serj"`
	Model        string `json:"model" example:"S-01"`
	HwVersion    string `json:"hw_version" example:"V1.0"`
	SwVersion    string `json:"sw_version" example:"V1.1"`
}

type Device struct {
	ID           string            `json:"id" example:"123"`
	Name         string            `json:"name" example:"lamp"`
	Description  string            `json:"description" example:"123"`
	Room         string            `json:"room" example:"123"`
	Type         string            `json:"type" example:"devices.properties.float"`
	Custom_data  map[string]string `json:"custom_data"`
	Capabilities []Capability      `json:"Capabilities"`
	Properties   []Property        `json:"properties"`
}

type PayLoad struct {
	UserID  string   `json:"user_id" example:"123"`
	Devices []Device `json:"devices"`
}
