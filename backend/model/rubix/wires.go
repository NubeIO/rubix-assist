package rubix

type WiresPlat struct {
	ClientId    string `json:"client_id" get:"true" put:"true"`
	ClientName  string `json:"client_name" get:"true" put:"true"`
	CreatedOn   string `json:"created_on" get:"true" put:"true"`
	DeviceId    string `json:"device_id" get:"true" put:"true"`
	DeviceName  string `json:"device_name" get:"true" put:"true"`
	GlobalUuid  string `json:"global_uuid" get:"true" put:"true"`
	SiteAddress string `json:"site_address" get:"true" put:"true"`
	SiteCity    string `json:"site_city" get:"true" put:"true"`
	SiteCountry string `json:"site_country" get:"true" put:"true"`
	SiteId      string `json:"site_id" get:"true"  put:"true"`
	SiteLat     string `json:"site_lat" get:"true" put:"true"`
	SiteLon     string `json:"site_lon" get:"true"  put:"true"`
	SiteName    string `json:"site_name" get:"true"  put:"true"`
	SiteState   string `json:"site_state" get:"true"  put:"true"`
	SiteZip     string `json:"site_zip" get:"true"  put:"true"`
	TimeZone    string `json:"time_zone" get:"true" put:"true"`
	UpdatedOn   string `json:"updated_on" get:"true"`
}
