package notification

import (
	context "context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationNotify(t *testing.T) {
	srsFileNotification := NewSrsFileNotification()
	ctx := context.Background()
	go srsFileNotification.Notify(ctx, "dummyModel", "dummySerialNo", "dummyFileName", 34566)
	msg := <-Messages
	assert.Equal(t, "dummyFileName", msg.Obj.Payload.FileName, "assert the error")
}
