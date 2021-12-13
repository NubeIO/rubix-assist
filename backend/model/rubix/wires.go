package rubix

type WiresPlat struct {
	ClientId    string `json:"client_id"`
	ClientName  string `json:"client_name"`
	CreatedOn   string `json:"created_on"`
	DeviceId    string `json:"device_id"`
	DeviceName  string `json:"device_name"`
	GlobalUuid  string `json:"global_uuid"`
	SiteAddress string `json:"site_address"`
	SiteCity    string `json:"site_city"`
	SiteCountry string `json:"site_country"`
	SiteId      string `json:"site_id"`
	SiteLat     string `json:"site_lat"`
	SiteLon     string `json:"site_lon"`
	SiteName    string `json:"site_name"`
	SiteState   string `json:"site_state"`
	SiteZip     string `json:"site_zip"`
	TimeZone    string `json:"time_zone"`
	UpdatedOn   string `json:"updated_on"`
}
