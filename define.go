// mqttJsonStruct
package HiNetEdgeDAS

type HeartbeatContentStu struct {
	Name    string `json:"Name"`
	Layer   string `json:"Layer"`
	Node    string `json:"Node"`
	Pid     int    `json:"Pid"`
	Time    int64  `json:"Ts"`
	Content string `json:"Content"`
}

type RealTimePoint struct {
	Key   string      `json:"K"`
	Value interface{} `json:"V"`
	Err   string      `json:"E"`
	Time  int64       `json:"T"`
}

type RealTimeDataStu struct {
	Layer    string          `json:"Layer"`
	Node     string          `json:"Node"`
	Time     int64           `json:"Ts"`
	Tc       int             `json:"Tc"`
	TaskNode string          `json:"TaskNode"`
	Err      int             `json:"Err"`
	Point    []RealTimePoint `json:"Point,omitempty"`
}

type ControlReqVal struct {
	Layer    string      `json:"Layer"`
	Node     string      `json:"Node"`
	TaskNode string      `json:"TaskNode"`
	Key      string      `json:"K"`
	Value    interface{} `json:"V"`
}

type ControlReqStu struct {
	Layer  string          `json:"Layer"`
	Node   string          `json:"Node"`
	Time   int64           `json:"Ts"`
	Number string          `json:"Number"`
	CmdReq []ControlReqVal `json:"CmdReq"`
}

type ControlResVal struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Key      string `json:"K"`
	Res      int    `json:"Res"`
	Desc     string `json:"Desc"`
}

type ControlResStu struct {
	Layer  string          `json:"Layer"`
	Node   string          `json:"Node"`
	Time   int64           `json:"Ts"`
	Number string          `json:"Number"`
	CmdRes []ControlResVal `json:"CmdRes"`
}

type RealStatusStu struct {
	Layer string `json:"Layer"`
	Node  string `json:"Node"`
	Time  int64  `json:"Ts"`
	Desc  string `json:"Desc"`
}

type FuncTriggerStu struct {
	Layer    string   `json:"Layer"`
	Node     string   `json:"Node"`
	TaskNode []string `json:"TaskNode,omitempty"`
}

type FuncResetCycleStu struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Cycle    int64  `json:"Cycle"`
}

type ManageFuncStu struct {
	Layer   string      `json:"Layer"`
	Node    string      `json:"Node"`
	Time    int64       `json:"Ts"`
	Func    string      `json:"Func"`
	Message interface{} `json:"Message"`
}

type EventNoticeStu struct {
	Layer    string      `json:"Layer"`
	Node     string      `json:"Node"`
	TaskNode string      `json:"TaskNode"`
	Key      string      `json:"K"`
	Time     int64       `json:"Ts"`
	Message  interface{} `json:"Message,omitempty"`
}

type CellularStatusStu struct {
	Time     int64  `json:"Ts"`
	Hardware string `json:"Hardware"`
	IMEI     string `json:"IMEI"`
	ICCID    string `json:"ICCID"`
	STAT     string `json:"STAT"`
	CSQ      int    `json:"CSQ"`
	SPNNAME  string `json:"SPNNAME"`
	SPNID    string `json:"SPNID"`
	NET      string `json:"NET"`
	MCC      string `json:"MCC"`
	MNC      string `json:"MNC"`
	LAC      string `json:"LAC"`
	CI       string `json:"CI"`
	TAC      string `json:"TAC"`
	AGPS     bool   `json:"AGPS"`
	AGPSLAT  string `json:"AGPSLAT"`
	AGPSLNG  string `json:"AGPSLNG"`
	GPS      bool   `json:"GPS"`
	GPSLAT   string `json:"GPSLAT"`
	GPSLNG   string `json:"GPSLNG"`
}

type GnssStu struct {
	Time       int64   `json:"Ts"`
	Gnss       string  `json:"Gnss"`
	Type       string  `json:"Type"`
	Invalid    bool    `json:"Invalid"`
	Numbers    int     `json:"Numbers"`
	Lat        float64 `json:"Lat"`
	NS         string  `json:"NS"`
	Lng        float64 `json:"Lng"`
	EW         string  `json:"EW"`
	SpeedOcean float64 `json:"SpeedOcean"`
	TNCourse   float64 `json:"TNCourse"`
	GnssTime   int64   `json:"GnssTime"`
}

type SystemLogStu struct {
	Service string `json:"Service"`
	Level   string `json:"Level"`
	Content string `json:"Content"`
	Ts      int64  `json:"Ts"`
}

type HardwareIoObject struct {
	Value        bool `json:"Value"`
	StartLevel   int  `json:"StartLevel"`
	LowDuration  int  `json:"LowDuration"`
	HighDuration int  `json:"HighDuration"`
}

type SystemResourceCtrlReqStu struct {
	User string      `json:"User"`
	Pid  int         `json:"Pid"`
	Ts   int64       `json:"Ts"`
	Type string      `json:"Type"`
	Obj  interface{} `json:"Obj"`
}

type SystemResourceStatusStu struct {
	Name  string      `json:"Name"`
	Value interface{} `json:"Value"`
	Ts    int64       `json:"Ts"`
}
