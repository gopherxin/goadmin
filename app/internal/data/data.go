package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gopherxin/goadmin/app/internal/biz"
	"github.com/gopherxin/goadmin/app/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewSystemUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "goadmin/base/data"))
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := &Data{
		db:  db,
		log: l,
	}
	return d, cleanup, nil
}

// NewDB gorm Connecting to a Database
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "order-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&biz.SystemUser{}); err != nil {
		log.Fatal(err)
	}
	return db
}
