package HiNetEdgeDAS

type ArithmeticInterfacer interface {
	About() string
	RegMatch(interface{})
	Process(interface{}) interface{}
	Engine()
}

type ArithInputVal struct {
	Layer    string `json:"Layer"`
	DevNode  string `json:"DevNode"`
	DevName  string `json:"DevName"`
	TaskNode string `json:"TaskNode"`
	TaskName string `json:"TaskName"`
	KeyNode  string `json:"KeyNode"`
	KeyName  string `json:"KeyName"`
}

type ArithFeedbackVal struct {
	Layer    string `json:"Layer"`
	DevNode  string `json:"DevNode"`
	DevName  string `json:"DevName"`
	TaskNode string `json:"TaskNode"`
	TaskName string `json:"TaskName"`
	KeyNode  string `json:"KeyNode"`
	KeyName  string `json:"KeyName"`
	DataType string `json:"DataType"`
}

type ArithOutputVal struct {
	Layer       string `json:"Layer"`
	DevNode     string `json:"DevNode"`
	DevName     string `json:"DevName"`
	TaskNode    string `json:"TaskNode"`
	TaskName    string `json:"TaskName"`
	KeyNode     string `json:"KeyNode"`
	KeyName     string `json:"KeyName"`
	Description string `json:"Description"`
	Unit        string `json:"Unit"`
	DataType    string `json:"DataType"`
	WrEnable    int    `json:"WrEnable"`
}

type ArithmeticParamStu struct {
	TaskNode   string        `json:"TaskNode"`
	Arithmetic string        `json:"Arithmetic"`
	Params     interface{}   `json:"Params,omitempty"`
	Input      []interface{} `json:"Input,omitempty"`
	Feedback   []interface{} `json:"Feedback,omitempty"`
	Output     []interface{} `json:"Output,omitempty"`
	Script     string        `json:"Script,omitempty"`
}
