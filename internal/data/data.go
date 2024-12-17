package data

import (
	"fmt"
	"kratosTestApp/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserDb,
	NewUserRepo,
)

// Data .
type Data struct {
	UserDb *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, userDb *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	DataIns := &Data{
		UserDb: userDb,
	}
	return DataIns, cleanup, nil
}
func NewUserDb(c *conf.Data, logger log.Logger) *gorm.DB {
	fmt.Println("NewUserDb...")
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", c.Database.Source)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
