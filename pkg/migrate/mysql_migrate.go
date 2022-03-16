package migrate

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlConf "git.internal.yunify.com/benchmark/benchmark/pkg/infra/db/mysql"

	"git.internal.yunify.com/benchmark/benchmark/pkg/util/config"
)

var _ Migrate = (*mysqlMigrate)(nil)

type mysqlMigrate struct {
	configFilePath     string
	initFilePath       string
	migrateRootDirPath string
}

func NewMysqlMigrateCli(options ...Option) Migrate {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return &mysqlMigrate{
		configFilePath:     config.configFilePath,
		initFilePath:       config.initFilePath,
		migrateRootDirPath: config.migrateRootDirPath,
	}
}

func (m *mysqlMigrate) InitDB(ctx context.Context) error {
	db, err := m.openDB("")
	if err != nil {
		logrus.Fatalf("init db err: %v", err)
		return err
	}
	defer m.closeDB(db)
	initBytes, err := ioutil.ReadFile(m.initFilePath)
	if err != nil {
		logrus.Fatalf("migrate files read err: %v", err)
		return err
	}
	dbNames := parseInitSqlFile(string(initBytes), "exists")
	for _, dbSql := range dbNames {
		fmt.Printf(dbSql)
		result := db.Exec(dbSql)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (m *mysqlMigrate) MigrateUp(ctx context.Context) error {
	files, err := ioutil.ReadDir(m.migrateRootDirPath)
	if err != nil {
		return err
	}
	up := func(file fs.FileInfo) error {
		dbName := file.Name()
		db, err := m.openDB(dbName)
		if err != nil {
			logrus.Fatalf("init db err: %v", err)
			return err
		}
		defer m.closeDB(db)
		sqlDb, err := db.DB()
		if err != nil {
			logrus.Fatalf("init db err: %v", err)
			return err
		}
		driver, err := mysql.WithInstance(sqlDb, &mysql.Config{DatabaseName: dbName})
		if err != nil {
			logrus.Fatalf("migrate sqldb init err: %v", err)
			return err
		}
		mg, err := migrate.NewWithDatabaseInstance("file://"+path.Join(m.migrateRootDirPath, file.Name()), "postgres", driver)
		if err != nil {
			logrus.Fatalf("migrate database init err: %v", err)
			return err
		}
		defer mg.Close()
		if err := mg.Up(); err != nil {
			if err != migrate.ErrNoChange {
				return err
			}
		}
		return nil
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if err := up(file); err != nil {
			return err
		}
	}
	return nil
}

func (m *mysqlMigrate) MigrateDown(ctx context.Context) error {
	return nil
}

func (m *mysqlMigrate) openDB(dbName string) (db *gorm.DB, err error) {
	dbConfig := mysqlConf.Config{}
	dbConfig.Database = dbName
	err = config.NewConfig(config.WithFilePath(m.configFilePath)).Register("mysql", &dbConfig).Parse()
	if err != nil {
		return
	}
	return gorm.Open(mysqlDriver.Open(dbConfig.Dsn()))
}

func (m *mysqlMigrate) closeDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
