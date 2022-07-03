package modiniter

import (
	"errors"
	finder "github.com/Fevzik/finder"
)

var GModules []GRPCModule
var Mods []Module

func InitModule(m interface{}, config ModuleConfig) error {
	if config.InitMode == InitModeSA {
		if config.Store == nil || config.Server == nil {
			return errors.New("in sa mode server and store must be given")
		}
	}
	if config.InitMode == InitModeStorelessService {
		if config.Server == nil {
			return errors.New("in storeless service mode server must be given")
		}
	}
	mod, ok := m.(Module)
	if !ok {
		return errors.New("first parameter must be valid module interface")
	}
	modConfig := mod.GetConfig()
	modConfig.InitMode = config.InitMode
	modConfig.RouterInitMode = config.RouterInitMode
	modConfig.Server = config.Server
	modConfig.Store = config.Store
	modConfig.GRPCPort = config.GRPCPort
	modConfig.AMQPDsn = config.AMQPDsn

	adm, ok := m.(AdminModule)
	if ok {
		AdminMenu = adm.GetAdminMenu()
	}

	ab, ok := m.(ActionBarModule)
	if ok {
		ActionBar = append(ActionBar, ab.GetActions()...)
	}

	menuMod, ok := m.(MenuModule)
	if ok {
		Menu = append(Menu, menuMod.GetMenu()...)
	}

	reportMod, ok := m.(ReportModule)
	if ok {
		Reports = append(Reports, reportMod.GetReports()...)
	}

	permMod, ok := m.(PermissionModule)
	if ok {
		permissions := permMod.GetPermissions()
		ModulePermissionMap = append(ModulePermissionMap, finder.DiscoveryModule{
			Code:        mod.GetModuleCode(),
			Label:       mod.GetModuleLabel(),
			Permissions: permissions,
		})
	}

	widgetMod, ok := m.(WidgetModule)
	if ok {
		Widgets = append(Widgets, widgetMod.GetWidgets()...)
	}

	httpMod, ok := m.(HttpModule)
	if ok && config.Server != nil && (config.InitMode == InitModeSA || config.InitMode == InitModeStorelessService) {
		httpMod.ImportRoutes(config.Server, config.RouterInitMode)
		grpcMod, ok := m.(GRPCModule)
		if ok && config.GRPCPort != nil {
			GModules = append(GModules, grpcMod)
		}
	}

	amqpMod, ok := m.(AMQPModule)
	if ok && config.AMQPDsn != nil {
		amqpMod.SetAmqpConnector(*config.AMQPDsn)
		for _, v := range amqpMod.GetQueues() {
			err := amqpMod.CreateQueue(*config.AMQPDsn, v)
			if err != nil {
				return err
			}
		}
		amqpMod.ConsumeTasks()
	}

	etcdMod, ok := m.(EtcdModule)
	if ok && config.Etcd != nil {
		etcdMod.SetEtcdConnector(*config.Etcd)
	}

	elkMod, ok := m.(SearchModule)
	if ok && config.ELK != nil {
		elkMod.SetSearchConnector(*config.ELK)
	}

	Mods = append(Mods, mod)
	return nil
}
