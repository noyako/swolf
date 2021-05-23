package swolf

import (
	"database/sql"
	"sync"
	"time"
)

// Pool is used to cache database connections
type Pool interface {
	Get(string) (*sql.DB, bool)
	Put(string, *sql.DB)
	SetMaxConn(int)
}

type databasePool struct {
	mutex   *sync.Mutex
	conns   map[string]*sql.DB
	timing  map[string]uint64
	maxConn int
}

func newDatabasePool(cons int) *databasePool {
	return &databasePool{
		mutex:   &sync.Mutex{},
		conns:   make(map[string]*sql.DB),
		timing:  make(map[string]uint64),
		maxConn: cons,
	}
}

func (p *databasePool) Get(id string) (*sql.DB, bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	db, ok := p.conns[id]
	return db, ok
}

func (p *databasePool) Put(id string, db *sql.DB) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.Get(id); ok || (len(p.conns) < p.maxConn) {
		p.conns[id] = db
		p.timing[id] = uint64(time.Now().UnixNano())
		return
	}

	minID := findMinID(p.timing)

	p.conns[minID] = db
	p.timing[minID] = uint64(time.Now().UnixNano())
	return
}

func (p *databasePool) SetMaxConn(num int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	newCons := make(map[string]*sql.DB)
	newTimings := make(map[string]uint64)

	i := 0
	for k, v := range p.conns {
		if i >= num {
			return
		}
		i++
		newCons[k] = v
		newTimings[k] = p.timing[k]
	}

	p.conns = newCons
	p.timing = newTimings
}

func findMinID(m map[string]uint64) string {
	var minValue uint64
	minKey := ""
	for k, v := range m {
		if v < minValue || minValue == 0 {
			minKey = k
			minValue = v
		}
	}

	return minKey
}
