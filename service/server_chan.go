package service

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/spf13/viper"
)

func ServerChanPushAll(text string, desp string) {
	serverChans := viper.GetStringSlice("serverchan")
	for _, serverChan := range serverChans {
		serverChanPush(serverChan, text, desp)
	}
}

func serverChanPush(serverChan string, text string, desp string) {
	url := fmt.Sprintf("https://sc.ftqq.com/%v.send?text=%v&desp=%v", serverChan, text, desp)
	_, _ = requests.Get(url)
}
