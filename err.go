// err
package HiNetEdgeDAS

import (
	// "encoding/binary"
	// "fmt"
	// "math"
	"strings"
)

const (
	GetConnTimeout     = 0x00010000 //获取连接超时
	GetConnErr         = 0x00020000 //获取连接错误
	ConnReadTimeout    = 0x00030000 //读取超时
	ConnFault          = 0x00040000 //连接异常
	IllegalReturnData  = 0x00050000 //无效的响应
	UnknownProtocolErr = 0x00060000 //未知协议错误
	UnMatchDataTypeErr = 0x00070000 //数据类型不匹配
)

type ControlResContent struct {
	Res  int
	Desc string
}

const (
	ControlOk             string = "success"
	ControlDisable        string = "control disable"
	ControlValIsNotExist  string = "variable does not exist"
	ControlValNotSupp     string = "variable not supported"
	ControlTypeNotSupp    string = "type not supported"
	ControlInvalidValue   string = "invalid value"
	ControlValueLenBeyond string = "beyond limit length"
	ControlCmdFailure     string = "control cmd failure"
	DeviceIsOffline       string = "device is offline"
	DriverAckTimeout      string = "driver ack timeout"
)

var DriverControlResDef = [10]ControlResContent{
	{0, ControlOk},
	{1, ControlDisable},
	{2, ControlValIsNotExist},
	{3, ControlValNotSupp},
	{4, ControlTypeNotSupp},
	{5, ControlInvalidValue},
	{6, ControlValueLenBeyond},
	{7, ControlCmdFailure},
	{8, DeviceIsOffline},
	{9, DriverAckTimeout}}

var DriverContorlResCnMap = map[int]string{
	0: "成功",
	1: "控制使能禁用",
	2: "变量不存在",
	3: "变量不支持写入",
	4: "写入值类型不支持",
	5: "写入值无效",
	6: "写入值长度超出",
	7: "写入设备失败",
	8: "设备不在线",
	9: "设备无响应"}

func SearchDescRes(str string) int {
	for _, val := range DriverControlResDef {
		if strings.Index(str, val.Desc) != -1 {
			return val.Res
		}
	}
	return 0x00
}
