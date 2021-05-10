package swolf

import (
	"database/sql"

	"github.com/aoyako/swolf/connector"
	_ "github.com/lib/pq"
)

type Dealer struct {
	masterController masterController
	tenantController tenantController
}

func Setup(cfg Config) *Dealer {
	global, err := sql.Open(cfg.Driver, cfg.Connection)
	if err != nil {
		panic("Cannot open sql environment")
	}

	master, err := connector.Connect(cfg.Driver, cfg.Connection, cfg.MasterDatabase)
	if err != nil {
		panic("Cannot open master database")
	}

	cfg.Mapper.defaultArgs()

	var dealer Dealer
	dealer.masterController = newMasterDB(master,
		cfg.MasterTable, cfg.MasterData.getKey(),
		cfg.MasterData.getTenant(),
		cfg.Mapper)

	dealer.tenantController = newTenantDB(global, cfg.Connection, cfg.Driver, cfg.Template)

	return &dealer
}

func (d *Dealer) Create(id string) (*sql.DB, error) {
	name, err := d.masterController.Create(id)
	if err != nil {
		return nil, err
	}
	return d.tenantController.Create(name)
}

func (d *Dealer) Get(id string) (*sql.DB, error) {
	name, err := d.masterController.Get(id)
	if err != nil {
		return nil, err
	}
	return d.tenantController.Get(name)
}

func (d *Dealer) Delete(id string) error {
	name, err := d.masterController.Delete(id)
	if err != nil {
		return err
	}
	return d.tenantController.Delete(name)
}
