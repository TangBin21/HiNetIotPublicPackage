// pointStruct
package HiNetEdgeDAS

import (
	"github.com/gopcua/opcua/ua"
)

type PointMsg struct {
	RuleIndex    int32   //规则序号
	KeyName      string  //关键词
	Tag          string  //标签
	TagType      string  //标签类型
	Reg          string  //寄存器
	AreaOrSlave  int32   //站号或区域地址
	Offset       uint32  //偏移地址
	DataType     string  //数据类型
	StrLen       uint32  //字符串长度
	BitAddr      int32   //位地址
	Ratio        float64 //倍率
	Cardinality  float64 //基数
	WrEnable     int32   //权限
	DecimalPlace int32   //小数位
	RetValue     float64 //返回数据
	RetStr       string  //返回字符串
	ErrCode      uint32  //有效标识
	Time         int64   //时间戳
	Order        string  //字节序
	Service      []byte
	OpcuaID      *ua.NodeID
}

type RuleMsg struct {
	Reg            string //寄存器
	CellFormat     string //寄存器单元格式
	StartAddr      uint32 //起始地址
	Length         uint32 //采集长度
	AreaOrSlave    int32  //站号或区域地址
	PointListStart int32  //点链表起始编号
	PointListCount int32  //点链表数量
}

type TaskMsg struct {
	TaskNode        string //任务节点
	Cycle           int32  //周期
	PointList       *List
	RuleList        *List
	ResetCycleCmdCh chan *FuncResetCycleStu
	TriggerCmdCh    chan *FuncTriggerStu
}

type ControlPoint struct {
	Layer       string
	Node        string
	TaskNode    string      //任务节点
	KeyName     string      //关键词
	Tag         string      //标签
	Reg         string      //寄存器
	AreaOrSlave int32       //站号或区域地址
	Offset      uint32      //偏移地址
	DataType    string      //数据类型
	BitAddr     int32       //位地址
	WByte       []byte      //写入字节
	Value       interface{} //写入值
	Desc        string      //结果返回
	OpcuaID     *ua.NodeID
}

type ControlReqMsg struct {
	Layer   string
	Node    string
	Numbers string //控制请求单号
	Time    int64
	Points  []ControlPoint
}

type ControlTaskStu struct {
	// Lock sync.Mutex
	Msg chan ControlReqMsg
}
