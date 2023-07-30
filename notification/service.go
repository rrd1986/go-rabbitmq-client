package notification

import (
	"context"
	"strconv"
	"time"

	"github.com/rrd1986/go-rabbitmq-client/logs"

	"github.com/rrd1986/go-rabbitmq-client/utils"
)

const MftID = "36767782"
const SessionID = "c4de24f829ec42a59facae8a08fab1e2"
const TransferType = "OPERATIONS-GATEWAY"
const GwSerialNumber = "OPCCMOCK001"
const SiteID = "Durham"
const Type = "MESSAGING"
const Category = "UI_DATA"
const Status = "NOTIFIED"
const RoutingKey = "small.opcc.platformstatus"
const Subscriber = "OPERATIONS-GATEWAY"
const Fin = "975c1700-497b-11e8-864b-5f874762ea72"
const Source = "srs-mock-service"
const ObjType = "EsrsNotificationMessage"

type FileNotification interface {
	Notify(ctx context.Context, model string, serialNumber string, fileName string, fileSize int64)
}

type srsFileNotification struct{}

func NewSrsFileNotification() FileNotification {
	return srsFileNotification{}
}

func (n srsFileNotification) Notify(ctx context.Context, model string, serialNumber string, fileName string, fileSize int64) {
	var notificationID string

	// define the message
	payload := Payload{
		MftID:           MftID,
		FileName:        fileName,
		SerialNumber:    serialNumber,
		Model:           model,
		FileSize:        strconv.Itoa(int(fileSize)),
		SessionID:       SessionID,
		FilelocationURI: utils.MockSrsUrlPathPrefix + model + "/" + serialNumber + "/" + fileName,
		TransferType:    TransferType,
		GwSerialNumber:  GwSerialNumber,
		SiteID:          SiteID,
	}

	obj := Obj{
		NotificationID:      notificationID,
		Type:                Type,
		Category:            Category,
		Payload:             payload,
		Status:              Status,
		NotificationTime:    time.Now().String(),
		AcknowledgementTime: time.Now().String(),
		RoutingKey:          RoutingKey,
		Subscriber:          Subscriber,
		Fin:                 Fin,
		FileProcessed:       false,
	}

	lzNotificationMsg := LzNotificationMessage{
		Source:    Source,
		Timestamp: time.Now().UnixNano(),
		ObjType:   ObjType,
		Obj:       obj,
	}

	Messages <- lzNotificationMsg
	logs.Logger.Info(ctx, "File upload notification to mock srs for opcc serial no %s sent", serialNumber)
}
