package delivery

import (
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/config"
	"enigmacamp.com/enigma-laundry-apps/delivery/controller"
	"enigmacamp.com/enigma-laundry-apps/manager"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	service manager.ServiceManager
	engine  *gin.Engine
	host    string
	log     *logrus.Logger
}

func (app *appServer) initController() {

	// app.engine.Use(middleware.LogRequestMiddleware(app.log))

	controller.NewEmployeeController(app.engine, app.service.EmployeeService())
	controller.NewCustomerController(app.engine, app.service.CustomerService())
	controller.NewProductController(app.engine, app.service.ProductService())
	controller.NewBillController(app.engine, app.service.BillService())
	controller.NewUserController(app.engine, app.service.UserService())
	controller.NewAuthController(app.engine, app.service.AuthService())
}

func (app *appServer) Run() {
	app.initController()

	if err := app.engine.Run(app.host); err != nil {
		panic(err)
	}
}

func Server() *appServer {
	engine := gin.Default()

	infraManager, err := manager.NewInfraManager(config.Cfg)

	if err != nil {
		panic(err)
	}

	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	return &appServer{
		service: repoManager,
		engine:  engine,
		host:    fmt.Sprintf("%s:%d", config.Cfg.Database.Host, config.Cfg.Server.Port),
		log:     logrus.New(),
	}

}
