package mysql

import (
	"fmt"
	"server/component/config"
	"server/logic/orm/dal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	dal.SetDefault(db)
}
