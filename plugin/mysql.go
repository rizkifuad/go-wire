package plugin

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func NewMysqlConnection() (*gorm.DB, error) {
	dbClient := viper.GetString(`database.client`)
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open(dbClient, dsn)
	if err != nil {
		return dbConn, err
	}

	dbConn.LogMode(viper.GetBool(`debug`))

	return dbConn, nil
}
