package lobby

import (
	"github/beijian128/micius/frame/appframe"
	_ "net/http/pprof"
)

var userMgr *userManager

var AppInstance *appframe.Application


func InitLobbySvr(app *appframe.Application) error {

	AppInstance = app


	userMgr = initUserManager(app)

	initUserMsgHandler(app)

	// 服务注册
	//app.RegisterService(MX1.SvrTypeAuth, appframe.WithLoadBalanceRandom(app, MX1.SvrTypeAuth))
	//app.RegisterService(MX1.SvrTypeAI, appframe.WithLoadBalanceRandom(app, MX1.SvrTypeAI))
	//app.RegisterService(MX1.SvrTypeGame, appframe.WithLoadBalanceRandom(app, MX1.SvrTypeGame))
	//app.RegisterService(MX1.SvrTypeFriend, appframe.WithLoadBalanceRandom(app, MX1.SvrTypeFriend))
	//app.RegisterService(MX1.SvrTypeRank, appframe.WithLoadBalanceRandom(app, MX1.SvrTypeRank))

	app.OnExitHandler(Close)
	app.OnFiniHandler(Finish)

	return nil
}

func Close() {

}

func Finish() {}
