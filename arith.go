package HiNetEdgeDAS

type ArithmeticInterfacer interface {
	About() string
	RegMatch(interface{})
	Process(interface{}) interface{}
	Engine()
}

type ArithInputVal struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Key      string `json:"Key"`
}

type ArithFeedbackVal struct {
	Layer    string `json:"Layer"`
	Node     string `json:"Node"`
	TaskNode string `json:"TaskNode"`
	Key      string `json:"Key"`
	DataType string `json:"DataType"`
}

type ArithOutputVal struct {
	Layer       string `json:"Layer"`
	Node        string `json:"Node"`
	TaskNode    string `json:"TaskNode"`
	Key         string `json:"Key"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Unit        string `json:"Unit"`
	DataType    string `json:"DataType"`
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
