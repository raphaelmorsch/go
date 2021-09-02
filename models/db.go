package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //PostgreSQL Driver
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	fmt.Printf("postgres://%s:%s@%s/%s?port=%d\n",
		"dbUser", "yourPassword", "dbHost", "dbName", 5432)
	orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?port=%d",
		"dbUser", "yourPassword", "dbHost", "dbName", 5432))
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
	return ormObject
}
