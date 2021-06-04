package HiNetEdgeDAS

type SourceIotTask struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	Tasknode string `json:"TaskNode"`
}

//type SourceGnss struct {
//	BDS    bool `json:"BDS"`
//	GPS    bool `json:"GPS"`
//	MCycle int  `json:"MCycle"`
//	SCycle int  `json:"SCycle"`
//}
//
//type SourceCellular struct {
//	Enable bool `json:"Enable"`
//	Cycle  int  `json:"Cycle"`
//}

type SourceSysObject struct {
	Layer    string
	Node     string
	TaskNode string
	Key      string
}

type SourceSys struct {
	Enable  bool              `json:"Enable"`
	Cycle   int               `json:"Cycle"`
	Objects []SourceSysObject `json:"Objects"`
}

type SouthSource struct {
	Iot []SourceIotTask `json:"Iot,omitempty"`
	//Gnss     SourceGnss      `json:"Gnss,omitempty"`
	//Cellular SourceCellular  `json:"Cellular,omitempty"`
	Sys SourceSys `json:"Sys,omitempty"`
}

type CustomKey struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Key      string `json:"Key"`
	NewKey   string `json:"NewKey"`
}
