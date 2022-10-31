package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	log "github.com/sirupsen/logrus"
)

type Mapping struct {
	FlowNetworkUUID   string
	NetworkUUID       string
	NetworkUUIDRemote string
}

func (inst *FlowClient) MakeLocalStreams(args *Mapping) error {

	network, err := inst.GetNetworkWithPoints(args.NetworkUUID)
	if err != nil {
		log.Errorf("get network err: %s", err.Error())
		return err
	}
	for _, device := range network.Devices {
		newStream := &model.Stream{
			CommonStream: model.CommonStream{
				CommonName: model.CommonName{
					Name: fmt.Sprintf("%s:%s", network.Name, device.Name),
				},
			},
		}
		streamResp, err := inst.AddStreamToExistingFlow(args.FlowNetworkUUID, newStream, false, Remote{})
		if err != nil {
			log.Errorf("add stream err: %s", err.Error())
			return err
		}
		log.Infof("added new stream %s", streamResp.Name)

		for _, point := range device.Points {
			newProducer := &model.Producer{
				CommonNameUnique: model.CommonNameUnique{
					Name: fmt.Sprintf("%s:%s", device.Name, point.Name),
				},
				ProducerThingName:  point.Name,
				ProducerThingUUID:  point.UUID,
				ProducerThingClass: point.ThingClass,
				ProducerThingType:  point.ThingType,
				StreamUUID:         streamResp.UUID,
			}
			producer, err := inst.AddProducer(newProducer, false, Remote{})
			if err != nil {
				return err
			}
			log.Infof("added new producer %s", producer.Name)

		}

	}

	return nil

}

func (inst *FlowClient) MakeRemoteDevicePoints(args *Mapping) error {

	network, err := inst.GetNetworkWithPoints(args.NetworkUUID)
	if err != nil {
		log.Errorf("get network err: %s", err.Error())
		return err
	}

	for _, device := range network.Devices {
		newDevice := &model.Device{
			Name:        device.Name,
			NetworkUUID: args.NetworkUUIDRemote,
		}
		addDevice, err := inst.AddDevice(newDevice, true, Remote{FlowNetworkUUID: args.FlowNetworkUUID})
		if err != nil {
			log.Errorf("add device err: %s", err.Error())
			return err
		}
		log.Infof("added new device %s", addDevice.Name)
		for _, point := range device.Points {
			newPoint := &model.Point{
				Name:       point.Name,
				DeviceUUID: addDevice.UUID,
			}
			addPoint, err := inst.AddPoint(newPoint, true, Remote{FlowNetworkUUID: args.FlowNetworkUUID})
			if err != nil {
				log.Errorf("add point err: %s", err.Error())
				return err
			}
			log.Infof("added new point %s", addPoint.Name)
		}
	}
	return nil

}
