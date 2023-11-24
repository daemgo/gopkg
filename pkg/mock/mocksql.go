package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetNewMySqlMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName:                "mysql",
			Conn:                      db,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})

	if err != nil {
		return gormDB, mock, err
	}

	return gormDB, mock, err
}

func GetNewPostgresMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(postgres.Dialector{
		Config: &postgres.Config{
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		},
	}, &gorm.Config{})
	if err != nil {
		return gormDB, mock, err
	}

	return gormDB, mock, nil
}
