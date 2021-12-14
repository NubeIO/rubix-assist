package rubix

type DiscoveredSlaves struct {
	Slaves []Slaves `json:"slaves"`
}

type Slaves struct {
	GlobalUUID   string `json:"global_uuid" get:"true" post:"true" delete:"true"`
	CreatedOn    string `json:"created_on,omitempty"  get:"true"`
	UpdatedOn    string `json:"updated_on,omitempty"  get:"true"`
	ClientID     string `json:"client_id,omitempty"  get:"true"`
	ClientName   string `json:"client_name,omitempty" get:"true"`
	SiteID       string `json:"site_id,omitempty"  get:"true"`
	SiteName     string `json:"site_name,omitempty"  get:"true"`
	DeviceID     string `json:"device_id,omitempty"  get:"true"`
	DeviceName   string `json:"device_name,omitempty" get:"true"`
	SiteAddress  string `json:"site_address,omitempty" get:"true"`
	SiteCity     string `json:"site_city,omitempty" get:"true"`
	SiteState    string `json:"site_state,omitempty" get:"true"`
	SiteZip      string `json:"site_zip,omitempty" get:"true"`
	SiteCountry  string `json:"site_country,omitempty" get:"true"`
	SiteLat      string `json:"site_lat,omitempty" get:"true"`
	SiteLon      string `json:"site_lon,omitempty" get:"true"`
	TimeZone     string `json:"time_zone,omitempty" get:"true"`
	IsMaster     bool   `json:"is_master,omitempty" get:"true"`
	Count        int    `json:"count,omitempty" get:"true"`
	IsOnline     bool   `json:"is_online,omitempty" get:"true"`
	TotalCount   int    `json:"total_count,omitempty" get:"true"`
	FailureCount int    `json:"failure_count,omitempty" get:"true"`
}
