package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

func bodyFFNetworkBody(ctx *gin.Context) (dto model.FFNetwork, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func bodyFFFlowNetworkBody(ctx *gin.Context) (dto model.FFFlowNetwork, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

var FlowUrls = struct {
	FlowNetworks     string
	Streams          string
	Writers          string
	Networks         string
	Devices          string
	Points           string
	Plat             string
	MasterConnection string
}{
	FlowNetworks:     "/ff/api/flow_networks",
	Streams:          "/ff/api/streams",
	Writers:          "/ff/api/points",
	Networks:         "/ff/api/networks",
	Devices:          "/ff/api/devices",
	Points:           "/ff/api/points",
	Plat:             "/api/wires/plat",
	MasterConnection: "/ff/api/localstorage_flow_network",
}

// FFNetworkWizard wizard for adding a new network, device and points
func (base *Controller) FFNetworkWizard(ctx *gin.Context) {
	body, err := bodyFFNetworkBody(ctx)
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
		NonProxyReq:  true,
	}
	proxyReq, opt, rtn, err := base.buildReq(po)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}

	b := map[string]interface{}{
		"name":        body.NetworkName,
		"plugin_name": body.PluginName,
	}

	opt = &nrest.ReqOpt{
		Timeout:          2 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    0 * time.Second,
		RetryMaxWaitTime: 0,
		Headers:          map[string]interface{}{"Authorization": rtn.Token},
		Json:             b,
	}

	makeNetwork := proxyReq.Do(nrest.POST, FlowUrls.Networks, opt)
	netUUID := gjson.Get(string(makeNetwork.Body), "uuid")
	log.Println("netUUID ", netUUID)
	log.Println("makeNetwork status ", makeNetwork.Status())
	log.Println("makeNetwork body ", makeNetwork.AsString())

	b = map[string]interface{}{
		"name":         body.DeviceName,
		"network_uuid": netUUID.String(),
	}
	opt.Json = b
	log.Println("opt.Json ", opt.Json)
	makeDev := proxyReq.Do(nrest.POST, FlowUrls.Devices, opt)
	devUUID := gjson.Get(string(makeDev.Body), "uuid")
	log.Println("makeDev status ", makeDev.Status())
	log.Println("makeDev body ", makeDev.AsString())
	log.Println("devUUID ", devUUID)

	for _, pnt := range body.Points {
		b = map[string]interface{}{
			"name":        pnt,
			"device_uuid": devUUID.String(),
		}
		opt.Json = b
		log.Println("opt.Json ", opt.Json)
		makePnt := proxyReq.Do(nrest.POST, FlowUrls.Points, opt)
		pntUUID := gjson.Get(string(makePnt.Body), "uuid")
		log.Println("makePnt status ", makePnt.Status())
		log.Println("makePnt body ", makePnt.AsString())
		log.Println("pntUUID ", pntUUID)

	}

}

func (base *Controller) FFFlowNetworkWizard(ctx *gin.Context) {
	body, err := bodyFFFlowNetworkBody(ctx)
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
		NonProxyReq:  true,
	}
	proxyReq, opt, rtn, err := base.buildReq(po)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}

	opt = &nrest.ReqOpt{
		Timeout:          2 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    0 * time.Second,
		RetryMaxWaitTime: 0,
		Headers:          map[string]interface{}{"Authorization": rtn.Token},
	}

	getPlat := proxyReq.Do(nrest.GET, FlowUrls.Plat, opt)
	siteName := gjson.Get(string(getPlat.Body), "site_name")
	log.Println("getPlat status ", getPlat.Status())
	log.Println("getPlat body ", getPlat.AsString())
	log.Println("siteName ", siteName.String())

	getMaster := proxyReq.Do(nrest.GET, FlowUrls.MasterConnection, opt)
	flowToken := gjson.Get(string(getPlat.Body), "flow_token")
	log.Println("getMaster status ", getMaster.Status())
	log.Println("getMaster body ", getMaster.AsString())
	log.Println("flowToken ", flowToken.String())

	b := map[string]interface{}{
		"name":                      strings.ToUpper(siteName.String()),
		"fetch_histories":           false,
		"fetch_hist_frequency":      0,
		"delete_histories_on_fetch": false,
		"is_master_slave":           body.IsMasterSlave,
		"is_mqtt":                   body.IsMqtt,
		"flow_https":                body.FlowHttps,
		"flow_ip":                   body.FlowIp,
		"flow_port":                 body.FlowPort,
		"flow_username":             body.FlowUsername,
		"flow_password":             body.FlowPassword,
		"flow_token":                flowToken.String(),
	}

	opt.Json = b
	log.Println("BODY ", opt.Json)
	log.Println("URL ", FlowUrls.FlowNetworks)
	makeNetwork := proxyReq.Do(nrest.POST, FlowUrls.FlowNetworks, opt)
	netUUID := gjson.Get(string(makeNetwork.Body), "uuid")
	log.Println("netUUID ", netUUID)
	log.Println("makeNetwork status ", makeNetwork.Status())
	log.Println("makeNetwork body ", makeNetwork.AsString())

	fn := make([]interface{}, 0)
	fn = append(fn, map[string]string{"uuid": netUUID.String()})

	b = map[string]interface{}{
		"name":          body.StreamName,
		"flow_networks": fn,
	}

	opt.Json = b
	log.Println("stream BODY ", opt.Json)
	makeStream := proxyReq.Do(nrest.POST, FlowUrls.Streams, opt)
	streamUUID := gjson.Get(string(makeStream.Body), "uuid")
	log.Println("streamUUID ", streamUUID)
	log.Println("makeStream status ", makeStream.Status())
	log.Println("makeStream body ", makeStream.AsString())

}
