package netlogger

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

var logtText = `asdadasd %+v`

type ADVideoAwardCallBackReq struct {
	Sign         string `json:"sign" form:"sign"` //签名 = sha256(appSecurityKey:TransId)
	UserId       string `json:"user_id" form:"user_id"`
	TransId      string `json:"trans_id" form:"trans_id"`
	RewardAmount int    `json:"reward_amount" form:"reward_amount"`
	RewardName   string `json:"reward_name" form:"reward_name"`
	Extra        string `json:"extra" form:"extra"`
}

func Test1(t *testing.T) {

	for i := 0; i < 1; i++ {
		agent := ZapLoggerAgent{}
		logger := agent.Init(&LogAgentConf{
			ServerName:  fmt.Sprintf("server%d", i),
			AgentAddr:   "127.0.0.1:8899",
			EncoderType: ConsoleEncoder,
		}).Conn().Daemon().Logger()

		go func(l *zap.Logger) {
			for {
				l.Sugar().Debug(logtText)
				l.Sugar().Error(logtText)
				l.Sugar().Infof(logtText, ADVideoAwardCallBackReq{})
				l.Sugar().Warn(logtText)
				//logger.Sugar().Panic(logtText)
				time.Sleep(time.Millisecond * 10)
			}
		}(logger)
	}

	select {}
}
