package controller

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
//	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
//	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
//	"github.com/gin-gonic/gin"
//	log "github.com/sirupsen/logrus"
//	"github.com/xuri/excelize/v2"
//	"mime/multipart"
//	"strings"
//	"time"
//)
//
//func bodyModbusIOConfig(ctx *gin.Context) (dto PointWriteBody, err error) {
//	err = ctx.ShouldBind(&dto)
//	return dto, err
//}
//
//type ObjectType string
//
//const (
//	//modbus
//	ObjTypeReadCoil           ObjectType = "read_coil"
//	ObjTypeReadCoils          ObjectType = "read_coils"
//	ObjTypeReadDiscreteInput  ObjectType = "read_discrete_input"
//	ObjTypeReadDiscreteInputs ObjectType = "read_discrete_inputs"
//	ObjTypeWriteCoil          ObjectType = "write_coil"
//	ObjTypeWriteCoils         ObjectType = "write_coils"
//	ObjTypeReadRegister       ObjectType = "read_register"
//	ObjTypeReadRegisters      ObjectType = "read_registers"
//	ObjTypeReadHolding        ObjectType = "read_holding"
//	ObjTypeReadHoldings       ObjectType = "read_holdings"
//	ObjTypeWriteHolding       ObjectType = "write_holding"
//	ObjTypeWriteHoldings      ObjectType = "write_holdings"
//	ObjTypeReadInt16          ObjectType = "read_int_16"
//	ObjTypeWriteInt16         ObjectType = "write_int_16"
//	ObjTypeReadUint16         ObjectType = "read_uint_16"
//	ObjTypeWriteUint16        ObjectType = "write_uint_16"
//	ObjTypeReadInt32          ObjectType = "read_int_32"
//	ObjTypeReadUint32         ObjectType = "read_uint_32"
//	ObjTypeReadFloat32        ObjectType = "read_float_32"
//	ObjTypeWriteFloat32       ObjectType = "write_float_32"
//	ObjTypeReadFloat64        ObjectType = "read_float_64"
//	ObjTypeWriteFloat64       ObjectType = "write_float_64"
//)
//
//const (
//	voltageDC   = "0-10 vdc"
//	voltage12dc = "12 vdc"
//	dryContact  = "Dry Contact"
//	temp        = "10K2 Thermistor"
//
//	voltageDCNum   = 1
//	voltage12dcNum = 2
//	dryContactNum  = 7
//	tempNum        = 3
//
//	uI1 = "UI-1"
//	uI2 = "UI-2"
//	uI3 = "UI-3"
//	uI4 = "UI-4"
//	uI5 = "UI-5"
//	uI6 = "UI-6"
//
//	uO1 = "UO-1"
//	uO2 = "UO-2"
//	uO3 = "UO-3"
//	uO4 = "UO-4"
//	uO5 = "UO-5"
//	uO6 = "UO-6"
//
//	ui1 = 201
//	ui2 = 202
//	ui3 = 203
//	ui4 = 204
//	ui5 = 205
//	ui6 = 206
//	ui7 = 207
//
//	uo1 = 101
//	uo2 = 102
//	uo3 = 103
//	uo4 = 104
//	uo5 = 105
//	uo6 = 106
//	uo7 = 107
//)
//
//func isUO(io string) bool {
//	if strings.Contains(io, "UO") {
//		return true
//	} else {
//		return false
//	}
//}
//
//func isUI(io string) bool {
//	if strings.Contains(io, "UI") {
//		return true
//	} else {
//		return false
//	}
//}
//
//func registers(io string) int {
//	switch io {
//	case uI1:
//		return ui1
//	case uI2:
//		return ui2
//	case uI3:
//		return ui3
//	case uI4:
//		return ui4
//	case uI5:
//		return ui5
//	case uI6:
//		return ui6
//	case uO1:
//		return uo1
//	case uO2:
//		return uo2
//	case uO3:
//		return uo3
//	case uO4:
//		return uo4
//	case uO5:
//		return uo5
//	case uO6:
//		return uo6
//	}
//	return 0
//}
//
//type PointWriteBody struct {
//	File          *multipart.FileHeader `form:"file"`
//	XLSFile       string                `json:"xls_file"`
//	ReturnArray   bool                  `json:"return_array"`
//	IsSerial      bool                  `json:"is_serial"`
//	DeviceAddress int                   `json:"device_address"`
//	Client        struct {
//		SerialPort        string `json:"serial_port"`
//		BaudRate          int    `json:"baud_rate"`
//		Parity            string `json:"parity"`
//		DeviceTimeoutInMs int    `json:"device_timeout_in_ms"`
//	} `json:"client"`
//	RequestBody struct {
//		ObjectType        string  `json:"object_type"`
//		Addr              int     `json:"addr"`
//		ZeroMode          bool    `json:"zero_mode"`
//		IsHoldingRegister bool    `json:"is_holding_register"`
//		ObjectEncoding    string  `json:"object_encoding"`
//		Length            int     `json:"length"`
//		WriteValue        float64 `json:"write_value"`
//	} `json:"request_body"`
//}
//
//type PointBuilder struct {
//	pointIONum      string //UI1
//	register        int    //201
//	pointName       string
//	pointType       int //Values:  0: RAW,   1: 0-10VDC,   2: 0-12VDC,   3: Any (Virtual)
//	pointTypeImport string
//}
//
//func pointType(io string) int {
//	//	AOs   Values:  0: RAW,   1: 0-10VDC,   2: 0-12VDC,   3: Any (Virtual)
//	//UIs      Values:  0: RAW,   1: 0-10ADC,   2: 10k (resistance),   3: 10k (type 2 temp)  4: 20k,   5: 4-20MA,    6: Pulse Count,        7: DI
//	switch io {
//	case "0-10 vdc", "0-10vdc", "0-10v", "0-10dc", "0-10 dc":
//		return voltageDCNum
//	case voltage12dc:
//		return voltage12dcNum
//	case dryContact:
//		return dryContactNum
//	case temp, "10K Thermistor":
//		return tempNum
//	}
//	return 0
//}
//
//func cleanPoint(data []string) *PointBuilder {
//	pointIONum := 0 //UI1
//	pointName := 2  //AHU-1-SAT
//	_pointType := 5 //10K2 Thermistor
//	//AI = 1
//	pb := new(PointBuilder)
//	for k, v := range data {
//		if k == pointIONum {
//			pb.pointIONum = v
//			pb.register = registers(v)
//		}
//		if k == pointName {
//			pb.pointName = v
//		}
//		if k == _pointType {
//			pb.pointType = pointType(v)
//			pb.pointTypeImport = v
//		}
//	}
//	if isUO(pb.pointIONum) && pb.pointName != "" {
//		return pb
//	} else if isUI(pb.pointIONum) && pb.pointName != "" {
//		return pb
//	} else {
//		return nil
//	}
//}
//
//func (base *Controller) ModbusPoll(ctx *gin.Context) {
//	body, err := bodyModbusIOConfig(ctx)
//	po := proxyOptions{
//		ctx:          ctx,
//		refreshToken: true,
//		NonProxyReq:  true,
//	}
//	proxyReq, opt, rtn, err := base.buildReq(po)
//	if err != nil {
//		reposeHandler(nil, err, ctx)
//		return
//	}
//
//	opt = &nrest.ReqOpt{
//		Timeout:          120 * time.Second,
//		RetryCount:       0,
//		RetryWaitTime:    0 * time.Second,
//		RetryMaxWaitTime: 0,
//		Headers:          map[string]interface{}{"Authorization": rtn.Token},
//		Json:             body,
//	}
//	getPlat := proxyReq.Do(nrest.POST, FlowUrls.ModbusPollPoint, opt)
//	d, _ := getPlat.AsJson()
//	reposeHandler(d, err, ctx)
//
//}
//
//func (base *Controller) ModbusIOConfig(ctx *gin.Context) {
//	_file, err := bodyModbusIOConfig(ctx)
//	po := proxyOptions{
//		ctx:          ctx,
//		refreshToken: true,
//		NonProxyReq:  true,
//	}
//	proxyReq, opt, rtn, err := base.buildReq(po)
//	if err != nil {
//		reposeHandler(nil, err, ctx)
//		return
//	}
//
//	_body := ctx.PostForm("json")
//	var body PointWriteBody
//	err = json.Unmarshal([]byte(_body), &body)
//	if err != nil {
//		reposeHandler(nil, err, ctx)
//	}
//
//	opt = &nrest.ReqOpt{
//		Timeout:          100 * time.Second,
//		RetryCount:       0,
//		RetryWaitTime:    0 * time.Second,
//		RetryMaxWaitTime: 0,
//		Headers:          map[string]interface{}{"Authorization": rtn.Token},
//		Json:             body,
//	}
//	getFile, err := _file.File.Open()
//	if err != nil {
//		reposeHandler(nil, err, ctx)
//		return
//	}
//	f, err := excelize.OpenReader(getFile)
//	if err != nil {
//		reposeHandler(nil, err, ctx)
//		return
//	}
//	defer func() {
//		// Close the spreadsheet.
//		if err := f.Close(); err != nil {
//			reposeHandler(nil, err, ctx)
//			return
//		}
//	}()
//	doImport := "G6"          //modbus 485 address
//	controllerAddress := "G5" //modbus 485 address
//	for i, sheet := range f.GetSheetList() {
//		rows, err := f.GetRows(sheet)
//		if err != nil {
//			reposeHandler(nil, err, ctx)
//			return
//		}
//		if i > -1 {
//			for _, row := range rows {
//				dImport, err := f.GetCellValue(sheet, doImport)
//				if err != nil {
//					reposeHandler(nil, err, ctx)
//					return
//				}
//				_dImport := string(dImport)
//				cAddress, err := f.GetCellValue(sheet, controllerAddress)
//				if err != nil {
//					reposeHandler(nil, err, ctx)
//					return
//				}
//				_cAddress := string(cAddress)
//				body.DeviceAddress = types.ToInt(_cAddress)
//				cp := cleanPoint(row)
//				if _dImport == "yes" {
//					if cp != nil {
//						fmt.Println("----------", "DEVICE ADDRESS", body.DeviceAddress, "sheet", sheet)
//						body.RequestBody.ObjectType = "write_uint_16"
//						body.RequestBody.Addr = cp.register
//						body.RequestBody.WriteValue = float64(cp.pointType)
//						opt.Json = body
//						log.Println(pprint.Print(cp))
//						getPlat := proxyReq.Do(nrest.POST, FlowUrls.ModbusPollPoint, opt)
//						fmt.Println("+++++++response+++++")
//						fmt.Println(getPlat.Err)
//						fmt.Println(getPlat.StatusCode)
//						fmt.Println(getPlat.Status())
//						fmt.Println(getPlat.AsString())
//						fmt.Println("++++++++++++")
//					}
//				}
//			}
//		}
//	}
//	reposeHandler("end", err, ctx)
//	return
//
//}
