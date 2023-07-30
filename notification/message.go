package notification

// make a gobal LzNotificationMessage channel variable to feed the notification to be read from the main thread
var Messages = make(chan LzNotificationMessage)

type Payload struct {
	MftID           string `json:"mftId"`
	FileName        string `json:"fileName"`
	SerialNumber    string `json:"serialNumber"`
	Model           string `json:"model"`
	FileSize        string `json:"fileSize"`
	SessionID       string `json:"sessionID"`
	FilelocationURI string `json:"filelocationURI"`
	TransferType    string `json:"transferType"`
	GwSerialNumber  string `json:"gwSerialNumber"`
	SiteID          string `json:"siteId"`
}

type Obj struct {
	NotificationID      string `json:"notificationId"`
	Type                string `json:"type"`
	Category            string `json:"category"`
	Payload             `json:"payload"`
	Status              string `json:"status"`
	NotificationTime    string `json:"notificationTime"`
	AcknowledgementTime string `json:"acknowledgementTime"`
	RoutingKey          string `json:"routingKey"`
	Subscriber          string `json:"subscriber"`
	Fin                 string `json:"fin"`
	FileProcessed       bool   `json:"fileProcessed"`
}

type LzNotificationMessage struct {
	Source    string `json:"source"`
	Timestamp int64  `json:"timestamp"`
	ObjType   string `json:"objType"`
	Obj       `json:"obj"`
}
