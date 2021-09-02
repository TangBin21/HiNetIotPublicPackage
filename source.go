package HiNetEdgeDAS

//懵逼啦

type SouthSource struct {
	Iot   SourceIot   `json:"Iot,omitempty"`
	Sys   SourceSys   `json:"Sys,omitempty"`
	Event SourceEvent `json:"Event,omitempty"`
}

type SourceIot struct {
	Enable bool      `json:"Enable"`
	Tasks  []MinUnit `json:"Tasks"`
}

type SourceSys struct {
	Enable bool      `json:"Enable"`
	Cycle  int       `json:"Cycle"`
	Tasks  []MinUnit `json:"Tasks"`
}

type MinUnit struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
}

type SourceEvent struct {
	GnssEnable     bool `json:"GnssEnable"`
	GnssCycleM     int  `json:"GnssCycleM"`
	GnssCycleS     int  `json:"GnssCycleS"`
	CellularEnable bool `json:"CellularEnable"`
	CellularCycle  int  `json:"CellularCycle"`
}

type CustomKey struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Key      string `json:"Key"`
	NewKey   string `json:"NewKey"`
	DatType  string `json:"DatType"`
}
