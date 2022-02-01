package main

import (
	"fmt"
	netval "github.com/THREATINT/go-net"
	"github.com/brotherpowers/ipsubnet"
	"github.com/jordan-wright/email"
	"github.com/mcnijman/go-emailaddress"
	"github.com/xuri/excelize/v2"
	"strings"
)

const (
	voltageDC   = "0-10 vdc"
	voltage12dc = "12 vdc"
	dryContact  = "Dry Contact"
	temp        = "10K2 Thermistor"

	voltageDCNum   = 1
	voltage12dcNum = 2
	dryContactNum  = 7
	tempNum        = 3

	uI1 = "UI-1"
	uI2 = "UI-2"
	uI3 = "UI-3"
	uI4 = "UI-4"
	uI5 = "UI-5"
	uI6 = "UI-6"

	uO1 = "UO-1"
	uO2 = "UO-2"
	uO3 = "UO-3"
	uO4 = "UO-4"
	uO5 = "UO-5"
	uO6 = "UO-6"

	ui1 = 201
	ui2 = 202
	ui3 = 203
	ui4 = 204
	ui5 = 205
	ui6 = 206
	ui7 = 207

	uo1 = 101
	uo2 = 102
	uo3 = 103
	uo4 = 104
	uo5 = 105
	uo6 = 106
	uo7 = 107
)

func isUO(io string) bool {
	if strings.Contains(io, "UO") {
		return true
	} else {
		return false
	}
}

func isUI(io string) bool {
	if strings.Contains(io, "UI") {
		return true
	} else {
		return false
	}
}

func ioUOs(io string) (int, string) {
	//	AOs   Values:  0: RAW,   1: 0-10VDC,   2: 0-12VDC,   3: Any (Virtual)
	switch io {
	case voltageDC:
		return voltageDCNum, voltageDC
	case voltage12dc:
		return voltage12dcNum, voltage12dc
	}
	return 3, ""
}

func ioUIs(io string) (int, string) {
	//UIs      Values:  0: RAW,   1: 0-10ADC,   2: 10k (resistance),   3: 10k (type 2 temp)  4: 20k,   5: 4-20MA,    6: Pulse Count,        7: DI
	switch io {
	case voltageDC:
		return voltageDCNum, voltageDC
	case dryContact:
		return dryContactNum, dryContact
	case temp:
		return tempNum, temp
	}
	return 3, ""
}

func registers(io string) int {
	switch io {
	case uI1:
		return ui1
	case uI2:
		return ui2
	case uI3:
		return ui3
	case uI4:
		return ui4
	case uI5:
		return ui5
	case uI6:
		return ui6
	case uO1:
		return uo1
	case uO2:
		return uo2
	case uO3:
		return uo3
	case uO4:
		return uo4
	case uO5:
		return uo5
	case uO6:
		return uo6
	}
	return 0
}

type PointWriteBody struct {
	NetworkType             string `json:"network_type"`          //RTU
	NetworkRtuPort          string `json:"network_rtu_port"`      // /dev/ttyUSB0
	DeviceAddress           int    `json:"device_address"`        //1
	PointRegister           int    `json:"point_register"`        //1
	PointRegisterLength     int    `json:"point_register_length"` //1
	PointFunctionCode       string `json:"point_function_code"`   //READ_INPUT_REGISTERS
	PointDataType           string `json:"point_data_type"`       //FLOAT
	PointPriorityArrayWrite struct {
		Field1 int `json:"_15"`
		Field2 int `json:"_16"`
	} `json:"point_priority_array_write"`
}

func main() {
	//github.com/jordan-wright/email
	e := email.NewEmail()
	e.From = "Nube Alerts <apick1066@gmail.com>"
	e.To = []string{"ap@nube-io.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "apick1066@gmail.com", "Apphp2508!!!", "smtp.gmail.com"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	ip := "192.168.112.999"
	web := "nube-io.com.au"
	sub := ipsubnet.SubnetCalculator(ip, 24)
	fmt.Println(sub.GetIPAddress())
	fmt.Println(sub.GetSubnetMask()) // 255.255.254.0
	fmt.Println(netval.IsIPAddr(ip))
	fmt.Println(netval.IsURL(web))
	_, err := emailaddress.Parse("1a-foobar.com")
	if err != nil {

	} else {

	}
	f, err := excelize.OpenFile("/home/aidan/Downloads/test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	doImport := "G6"          //modbus 485 address
	controllerAddress := "G5" //modbus 485 address
	pointNumber := 0          //AI
	pointType := 5            //0-10dc
	fmt.Println(doImport, controllerAddress)

	for i, sheet := range f.GetSheetList() {

		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		pt := ""
		fmt.Println(pointNumber, pt)
		if i <= 1 {
			for _, row := range rows {
				fmt.Println(row)
				for i, _row := range row {

					if i == pointType {
						pt = _row
					}
					//if i == pointNumber {
					//	if isUO(_row) {
					//		num, name := ioUOs(pt)
					//		config := fmt.Sprintf("type: %s config:[name:%s num: %d]", pt, name, num)
					//		fmt.Println("pointNumber", _row, "UO-config", config, "address", registers(_row))
					//	}
					//	if isUI(_row) {
					//		num, name := ioUIs(pt)
					//		config := fmt.Sprintf("type: %s config:[name:%s num: %d]", pt, name, num)
					//		fmt.Println("pointNumber", _row, "UO-config", config, "address", registers(_row))
					//
					//	}
					//}
				}
			}
		}
	}
}
