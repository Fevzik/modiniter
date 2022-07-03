package modiniter

import (
	finder "github.com/Fevzik/finder"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ModuleConfig struct {
	InitMode       string
	RouterInitMode string
	Server         *fiber.App
	Store          *sqlx.DB
	GRPCPort       *string
	AMQPDsn        *string
	Etcd           *[]string
	ELK           *[]string
}

type Module interface {
	GetModuleLabel() string
	GetModuleCode() string
	GetConfig() *ModuleConfig
}

type ActionBarModule interface {
	GetActions() ActionBarItemsList
}

type WidgetModule interface {
	GetWidgets() WidgetsList
}

type MenuModule interface {
	GetMenu() MenuItemList
}

type PermissionModule interface {
	GetPermissions() finder.PermissionList
}

type ReportModule interface {
	GetReports() ReportsList
}

type AdminModule interface {
	GetAdminMenu() AdminMenuList
}

type StoreModule interface {
	GetStore() *sqlx.DB
}

type HttpModule interface {
	ImportRoutes(app *fiber.App, mode string)
}

type GRPCModule interface {
	Module
	GStart() error
	GStop()
}

type AMQPModule interface {
	SetAmqpConnector(string)
	ConsumeTasks()
	CreateQueue(dsn string, queueName string) error
	GetQueues() []string
}

type EtcdModule interface {
	SetEtcdConnector(hosts []string)
}

type SearchModule interface {
	SetSearchConnector(hosts []string)
}
