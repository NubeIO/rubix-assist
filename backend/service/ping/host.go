package ping

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	utils "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/ping"
	"github.com/NubeIO/rubix-updater/model"
	dbhandler "github.com/NubeIO/rubix-updater/pkg/handler"
	"github.com/NubeIO/rubix-updater/pkg/jobs"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var alertTpe = model.CommonAlertTypes.HostPing

//createAlert make if not exist
func createAlert(host *model.Host) {
	d := dbhandler.GetDB()
	alert, _ := d.GetAlertByType(host.UUID, alertTpe)
	//make alert if not exist
	if alert == nil {
		var a model.Alert
		a.HostUUID = host.UUID
		a.AlertType = alertTpe
		a.Date = time.Now().UTC()
		alert, _ = d.CreateAlert(&a)
	}
	//add a message log to the alert
	var msg model.Message
	msg.AlertUUID = alert.UUID
	msg.Date = time.Now().UTC()
	_, _ = d.CreateMessage(&msg)

}

func run() {
	d := dbhandler.GetDB()
	hosts, _ := d.GetHosts()
	for _, host := range hosts {
		if bools.IsTrue(host.PingEnable) {
			_, _, ok := utils.PingPort(host.IP, strconv.Itoa(host.RubixPort), 5, false)
			fmt.Println(host.Name, ok)
			if !ok {
				host.IsOffline = bools.NewTrue()
				host.OfflineCount = host.OfflineCount + 1
				host, err := d.UpdateHost(host.UUID, &host)
				if err != nil {
					fmt.Println(err)
					//return
				}
				res := host.OfflineCount % 10
				if res == 0 {
					fmt.Println("SEND OFFLINE")
				}
				createAlert(host)
			}
		}
	}

}

func TEST() {
	j, ok := jobs.GetJobService()
	if ok {
		_, err := j.Every(120).Second().Do(run)
		if err != nil {
			log.Infof("system-plugin-schedule: error on create job %v\n", err)
		}
	}
}
