package main

import (
	"fmt"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/test"
	"github.com/larksuite/oapi-sdk-go/core/tools"
	eventhttpserver "github.com/larksuite/oapi-sdk-go/event/http/native"
	application "github.com/larksuite/oapi-sdk-go/service/application/v1"
	"net/http"
	"path"
)

func main() {

	var conf = test.GetISVConf("online")

	application.SetAppOpenEventHandler(conf, func(coreCtx *core.Context, appOpenEvent *application.AppOpenEvent) error {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(appOpenEvent)
		fmt.Println(tools.Prettify(appOpenEvent))
		return nil
	})

	application.SetAppStatusChangeEventHandler(conf, func(coreCtx *core.Context, appStatusChangeEvent *application.AppStatusChangeEvent) error {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(appStatusChangeEvent.Event.AppId)
		fmt.Println(appStatusChangeEvent.Event.Status)
		fmt.Println(tools.Prettify(appStatusChangeEvent))
		return nil
	})

	eventhttpserver.Register(path.Join("/", conf.GetAppSettings().AppID, "webhook/event"), conf)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		panic(err)
	}

}
