package HiNetEdgeDAS

type SourceIotTask struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	Tasknode string `json:"TaskNode"`
}

type SourceGnss struct {
	BDS    bool `json:"BDS"`
	GPS    bool `json:"GPS"`
	MCycle int  `json:"MCycle"`
	SCycle int  `json:"SCycle"`
}

type SourceCellular struct {
	Enable bool `json:"Enable"`
	Cycle  int  `json:"Cycle"`
}

type SourceHardware struct {
	Enable bool `json:"Enable"`
	Cycle  int  `json:"Cycle"`
}

type SouthSource struct {
	Iot      []SourceIotTask `json:"Iot,omitempty"`
	Gnss     SourceGnss      `json:"Gnss,omitempty"`
	Cellular SourceCellular  `json:"Cellular,omitempty"`
	Hardware SourceHardware  `json:"Hardware,omitempty"`
}

type CustomKey struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Key      string `json:"Key"`
	NewKey   string `json:"NewKey"`
}
