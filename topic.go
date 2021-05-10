// topic
package HiNetEdgeDAS

type TopicStu struct {
	Cont     string
	Qos      byte
	Retained bool
}

type MqttTopic struct {
	DriverRtData   TopicStu
	DriverRtStat   TopicStu
	DriverCtrlReq  TopicStu
	DriverCtrlRes  TopicStu
	DriverManage   TopicStu
	HardwareRtData TopicStu

	EdgeRtData  TopicStu
	EdgeRtStat  TopicStu
	EdgeCtrlReq TopicStu
	EdgeCtrlRes TopicStu
	EdgeTrigger TopicStu

	Heartbeat TopicStu
}
