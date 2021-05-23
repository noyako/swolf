package swolf

import (
	"database/sql"

	"github.com/noyako/swolf/connector"
)

const defaultConnectionsNumber = 10

// Dealer controls databases manipulation.
type Dealer struct {
	masterController MasterController
	tenantController TenantController
	pool             Pool
}

// Setup creates Dealer by config.
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

	dealer.pool = newDatabasePool(defaultConnectionsNumber)

	return &dealer
}

// Create creates new tenant with specified id.
func (d *Dealer) Create(id string) (*sql.DB, error) {
	name, err := d.masterController.Create(id)
	if err != nil {
		return nil, err
	}
	return d.tenantController.Create(name)
}

// Get returns tenant with specified id.
func (d *Dealer) Get(id string) (*sql.DB, error) {
	if db, ok := d.pool.Get(id); ok {
		return db, nil
	}

	name, err := d.masterController.Get(id)
	if err != nil {
		return nil, err
	}

	db, err := d.tenantController.Get(name)
	if err != nil {
		return nil, err
	}

	d.pool.Put(id, db)
	return db, nil
}

// Delete deletes tenant with specified id.
func (d *Dealer) Delete(id string) error {
	name, err := d.masterController.Delete(id)
	if err != nil {
		return err
	}
	return d.tenantController.Delete(name)
}
