package netlogger

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

var logtText = `asdadasd`

func Test1(t *testing.T) {

	for i := 0; i < 1; i++ {
		agent := ZapLoggerAgent{}
		logger := agent.Init(&LogAgentConf{
			ServerName: fmt.Sprintf("server%d", i),
			AgentAddr:  "127.0.0.1:8899",
		}).Conn().Daemon().Logger()
		go func(l *zap.Logger) {
			for {
				l.Sugar().Debug(logtText)
				l.Sugar().Error(logtText)
				l.Sugar().Info(logtText)
				l.Sugar().Warn(logtText)
				//logger.Sugar().Panic(logtText)
				time.Sleep(time.Millisecond * 10)
			}
		}(logger)
	}

	select {}
}
