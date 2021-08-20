package tools

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const defaultLifeTime = time.Minute * 5
const defaultMaxIdleConns = 5
const defaultMaxOpenConns = 5

type singleton struct {
	Connection *gorm.DB
}

type MultipleDatabaseOnPoolError struct{}

type PoolWithoutInstanceError struct{}

var (
	poolGormDb    map[string]*gorm.DB = make(map[string]*gorm.DB)
	LogModeEnable                     = func() bool {
		return true
	}()
)

//LoadGormMySQL with the following parameters:
/**
 * user string, pass string, host string, port int, dbName string
**/
func LoadGormMySQL(user string, pass string, host string, port int, dbName string) error {
	return LoadGorm("mysql", user, pass, host, port, dbName, false)
}

//LoadGormPostGres with the following parameters:
/**
 * user string, pass string, host string, port int, dbName string
**/
func LoadGormPostGres(user string, pass string, host string, port int, dbName string, sslMode bool) error {
	return LoadGorm("postgres", user, pass, host, port, dbName, sslMode)
}

// LoadGorm with the following parameters:
/**
 * driverName string (such as mysql) user string, pass string, host string, port int, dbName string
**/
func LoadGorm(driverName string, user string, pass string, host string, port int, dbName string, sslMode bool) error {
	var err error

	if poolGormDb[dbName] == nil {
		poolGormDb[dbName], err = getGormConnection(driverName, user, pass, host, port, dbName, sslMode)
	}

	return err
}

//GetGormDb return gorm.DB instance
func GetGormDb(dbNameParam ...string) (*gorm.DB, error) {
	dbName, err := defineDatabaseName(dbNameParam, len(poolGormDb), func() string {
		return firstKeyFromGormPool(poolGormDb)
	})

	if err != nil {
		return nil, err
	}
	if poolGormDb[dbName] == nil {
		return nil, errors.New("LoadGorm/SetGormDb wasn't called for database: " + dbName)
	}

	return poolGormDb[dbName], nil
}

func SetGormDb(gormDb *gorm.DB, dbName string) {
	poolGormDb[dbName] = gormDb
}

func getGormConnection(driverName string, user string, pass string, host string, port int, dbName string, sslMode bool) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var dsn string
	dsn, err = generateDsn(driverName, user, pass, host, port, dbName, sslMode)
	db, err = gorm.Open(driverName, dsn)

	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetConnMaxLifetime(defaultLifeTime)
	db.DB().SetMaxIdleConns(defaultMaxIdleConns)
	db.DB().SetMaxOpenConns(defaultMaxOpenConns)
	db.LogMode(LogModeEnable)
	return db, err
}

func firstKeyFromGormPool(object map[string]*gorm.DB) string {
	for k := range object {
		return k
	}

	return ""
}

// PurgeGormPool - Cleans pool closing connections
func PurgeGormPool() {
	for k, v := range poolGormDb {
		v.Close()
		delete(poolGormDb, k)
	}
}

func (e *MultipleDatabaseOnPoolError) Error() string {
	return fmt.Sprintf("It isn't allowed define a default database. You should pass the database name instead.")
}

func (e *PoolWithoutInstanceError) Error() string {
	return fmt.Sprintf("Can't define a default database. You should Set or Load a instance first.")
}

func defineDatabaseName(dbNameParam []string, poolSize int, firstKeyOfPool func() string) (string, error) {
	var dbName string

	if len(dbNameParam) == 0 {
		if poolSize == 0 {
			return "", &PoolWithoutInstanceError{}
		}

		if poolSize > 1 {
			return "", &MultipleDatabaseOnPoolError{}
		}

		dbName = firstKeyOfPool()
	} else {
		dbName = dbNameParam[0]
	}

	return dbName, nil
}

func generateDsn(driverName string, user string, pass string, host string, port int, dbName string, sslMode bool) (string, error) {
	if driverName == "mysql" {
		return user + ":" + pass + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8&parseTime=True&loc=America%2FSao_Paulo", nil
	}

	if driverName == "postgres" {

		var sslOption string

		if sslMode {
			sslOption = "enable"
		} else {
			sslOption = "disable"
		}

		return "host=" + host +
			" port=" + strconv.Itoa(port) +
			" user=" + user +
			" dbname=" + dbName +
			" password=" + pass +
			" sslmode=" + sslOption, nil
	}

	return "", errors.New("Can't generate DSN for " + driverName)
}
