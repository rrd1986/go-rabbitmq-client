package logs

import (
	"github.com/rrd1986/common-go-modules/log"
)

var Logger log.LoggerType

func init() {
	Logger = log.NewLogger("go-rabbitmq-client", "1.0")
}
