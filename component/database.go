package component

import (
	"fmt"

	"github.com/gedelumbung/go-catalog/config"
	"github.com/gedelumbung/go-catalog/repository"
	"github.com/gedelumbung/go-catalog/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabaseConnection(config *conf.Configuration) (repository.Repository, error) {
	/**
	* Multiple Database Purpose
	* We can define another databases here
	 */
	if config.DB.Driver == "mysql" {
		return mysql.Connect(config.DB.Mysql.URL)
	}
	return nil, fmt.Errorf("unknown store type: %s", config.DB.Driver)
}
