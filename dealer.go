package swolf

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Dealer struct {
	masterController *masterDB
	tenantController *tenantDB
}

func Setup(cfg Config) *Dealer {
	global, err := sql.Open(cfg.Driver, cfg.Connection)
	if err != nil {
		panic("Cannot open sql environment")
	}

	var master *sql.DB
	switch cfg.Driver {
	case "postgres":
		master, err = sql.Open(cfg.Driver, cfg.Connection+" dbname="+cfg.MasterDatabase)
	case "mysql":
		master, err = sql.Open(cfg.Driver, cfg.Connection+"/"+cfg.MasterDatabase)
	}
	if err != nil {
		panic("Cannot open master database")
	}

	cfg.Mapper.defaultArgs()

	var dealer Dealer
	dealer.masterController = newMasterDB(master,
		cfg.MasterTable, cfg.MasterData.getKey(),
		cfg.MasterData.getTenant(),
		cfg.Mapper)

	dealer.tenantController = newTenantDB(global, cfg.Connection, cfg.Template)

	return &dealer
}

func (d *Dealer) Create(id string) (string, error) {
	return d.masterController.Create(id)
}

func (d *Dealer) Get(id string) (string, error) {
	return d.masterController.Get(id)
}
