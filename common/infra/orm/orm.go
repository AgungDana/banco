package orm

import (
	"banco/common/config"
	"banco/common/ctxutil"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	pool *gormPool

	ErrEmptyTransaction   = errors.New("empty transaction")
	ErrWithoutTransaction = errors.New("without transaction")
)

func getPool() *gormPool {
	if pool == nil {
		once.Do(func() {
			pool = &gormPool{
				txns:  make(map[uuid.UUID]transaction),
				pools: make(map[string]*gorm.DB),
			}
		})
	}
	return pool
}

type transaction struct {
	tx *gorm.DB

	// begin transaction counter+1
	// commit transacntio counter-1, commit ketika 0
	// roolback force 0
	counter int
}

type gormPool struct {
	mutex sync.Mutex
	txns  map[uuid.UUID]transaction
	pools map[string]*gorm.DB
}

func DbCon(conf config.Config) (*gorm.DB, error) {
	log.Println("trying to establish a db connection")
	var (
		dsn, dsnWithoutPassword string
		db                      *gorm.DB
		exist                   bool
		err                     error
	)
	switch conf.Driver {
	case "mysql":
		dsn = conf.UserName + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.DbName + "?parseTime=True&loc=Local"
		// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// if err != nil {
		// 	return nil, err
		// }
		// return db, nil
	default:
		return nil, fmt.Errorf("invalid driver supported -> driver using %s", conf.Driver)
	}
	dsnWithoutPassword = strings.Replace(dsn, conf.Password, "", -1)

	if db, exist = getPool().pools[dsnWithoutPassword]; exist {
		log.Printf("reusing db connection -> use: %s", dsnWithoutPassword)
		return db, nil
	}

	log.Printf("create a new connection db")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, fmt.Errorf("cannot build db %s", dsnWithoutPassword)
	}

	getPool().pools[dsnWithoutPassword] = db
	return db, nil

}

func BeginTx(ctx context.Context, db *gorm.DB) *gorm.DB {
	id, withTransaction := ctxutil.GetTransactionId(ctx)

	if !withTransaction {
		return db
	}

	p := getPool()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	txn, exist := p.txns[id]
	if !exist {
		txn = transaction{tx: db.Begin()}
	} else {
		txn.counter += 1
	}
	p.txns[id] = txn
	return txn.tx
}

func ComitTx(ctx context.Context) error {
	id, withTransaction := ctxutil.GetTransactionId(ctx)

	if !withTransaction {
		return ErrWithoutTransaction
	}

	p := getPool()

	txn, exist := p.txns[id]
	if !exist {
		return ErrEmptyTransaction
	}

	if txn.counter > 0 {
		txn.counter -= 0
		p.txns[id] = txn
		return nil
	}

	err := txn.tx.Commit().Error
	endTx(ctx)
	return err
}

func RoolBackTxn(ctx context.Context) error {
	id, withTransaction := ctxutil.GetTransactionId(ctx)

	if !withTransaction {
		return ErrWithoutTransaction
	}

	p := getPool()

	txn, exist := p.txns[id]
	if !exist {
		return ErrEmptyTransaction
	}

	if txn.counter > 0 {
		txn.counter -= 1
		p.txns[id] = txn
		return nil
	}

	err := txn.tx.Rollback().Error
	endTx(ctx)
	return err
}

func endTx(ctx context.Context) error {
	id, withTransaction := ctxutil.GetTransactionId(ctx)

	if !withTransaction {
		return ErrWithoutTransaction
	}

	p := getPool()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	txn, exist := p.txns[id]
	if !exist {
		return ErrEmptyTransaction
	}

	if txn.counter > 0 {
		return nil
	}

	delete(p.txns, id)
	return nil
}
