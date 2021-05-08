package swolf

import "database/sql"

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
	switch cfg.Connection {
	case "postgres":
		master, err = sql.Open(cfg.Driver, cfg.Connection+" dbname="+cfg.MasterDatabase)
	case "mysql":
		master, err = sql.Open(cfg.Driver, cfg.Connection+"/"+cfg.MasterDatabase)
	}
	if err != nil {
		panic("Cannot open master database")
	}

	var dealer Dealer
	dealer.masterController = newMasterDB(master,
		cfg.MasterTable, cfg.MasterData.getKey(),
		cfg.MasterData.getTenant(),
		cfg.Mapper)

	dealer.tenantController = newTenantDB(global, cfg.Connection, cfg.Template)

	return &dealer
}
