package ping

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	utils "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/ping"
	dbhandler "github.com/NubeIO/rubix-updater/pkg/handler"
	"github.com/NubeIO/rubix-updater/pkg/jobs"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func run() {

	d := dbhandler.GetDB()

	hosts, _ := d.GetHosts()

	for _, host := range hosts {
		fmt.Println(host.Name)
		if bools.IsTrue(host.PingEnable) {
			_, err, ok := utils.PingPort(host.IP, strconv.Itoa(host.RubixPort), 5, false)
			if err != nil {

			}
			fmt.Println(host.Name, ok)
			if !ok {
				host.IsOffline = bools.NewTrue()
				fmt.Println(host.OfflineCount)
				host.OfflineCount = host.OfflineCount + 1
				_, err := d.UpdateHost(host.ID, &host)
				if err != nil {
					fmt.Println(err)
					//return
				}
				res := host.OfflineCount % 10
				if res == 0 {
					fmt.Println("SEND OFFLINE")
				}

			}

		}

	}

}

func TEST() {
	j, ok := jobs.GetJobService()

	if ok {
		_, err := j.Every(30).Second().Do(run)
		if err != nil {
			log.Infof("system-plugin-schedule: error on create job %v\n", err)
		}
	}

}
