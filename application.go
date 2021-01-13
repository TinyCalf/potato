package potato

import (
	"potato/piface"
	"potato/compnents"
	"fmt"
)

// Application ..
type Application struct {
	components map[string]piface.ICompnent
}

// NewApplication returns new Application
func NewApplication() piface.IApplication {
	app := &Application{
		components: make(map[string]piface.ICompnent),
	}

	//添加组件
	connector := compnents.NewConnector(app)
	sessionservice := compnents.NewSessionService(app)
	app.AddComponent(connector)
	app.AddComponent(sessionservice)

	return app
}

//AddComponent 增加一个组件
func (app *Application) AddComponent(compnent piface.ICompnent) {
	name := compnent.GetName()
	if _, exists := app.components[name]; exists {
		panic(fmt.Sprintf("组件名称冲突 %s", name))
	}
	app.components[name] = compnent
}

//GetComponent 获取一个组件
func (app *Application) GetComponent(name string) piface.ICompnent{
	return app.components[name]
}

// Start 启动服务
func (app *Application) Start() {
	fmt.Println("App start running")

	//启动所有组件
	for _,component := range app.components {
		component.OnAppStart()
	}
	//阻塞
	select {}
}