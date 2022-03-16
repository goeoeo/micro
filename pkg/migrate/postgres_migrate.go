package migrate

import (
	"context"
	"io/fs"
	"io/ioutil"
	"path"

	m "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"git.internal.yunify.com/benchmark/benchmark/pkg/infra/db/pgsql"
	"git.internal.yunify.com/benchmark/benchmark/pkg/util/config"
)

var _ Migrate = (*postgresMigrate)(nil)

type postgresMigrate struct {
	configFilePath     string
	initFilePath       string
	migrateRootDirPath string
}

func NewPostgresMigrateCli(options ...Option) Migrate {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return &postgresMigrate{
		configFilePath:     config.configFilePath,
		initFilePath:       config.initFilePath,
		migrateRootDirPath: config.migrateRootDirPath,
	}
}

func (p *postgresMigrate) InitDB(ctx context.Context) error {
	db, err := p.openDB("postgres")
	if err != nil {
		logrus.Fatalf("init db err: %v", err)
		return err
	}
	defer p.closeDB(db)
	initBytes, err := ioutil.ReadFile(p.initFilePath)
	if err != nil {
		logrus.Fatalf("migrate files read err: %v", err)
		return err
	}
	dbNames := parseInitSqlFile(string(initBytes), "database")
	for dbName, dbSql := range dbNames {
		var count int64
		db.Table("pg_catalog.pg_database").Where("datname", dbName).Count(&count)
		if count == 0 {
			result := db.Exec(dbSql)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

func (p *postgresMigrate) MigrateUp(ctx context.Context) error {
	files, err := ioutil.ReadDir(p.migrateRootDirPath)
	if err != nil {
		return err
	}
	up := func(file fs.FileInfo) error {
		dbName := file.Name()
		db, err := p.openDB(dbName)
		if err != nil {
			logrus.Fatalf("init db err: %v", err)
			return err
		}
		defer p.closeDB(db)
		sqlDb, err := db.DB()
		if err != nil {
			logrus.Fatalf("init db err: %v", err)
			return err
		}
		driver, err := postgres.WithInstance(sqlDb, &postgres.Config{DatabaseName: dbName})
		if err != nil {
			logrus.Fatalf("migrate sqldb init err: %v", err)
			return err
		}
		mg, err := m.NewWithDatabaseInstance("file://"+path.Join(p.migrateRootDirPath, file.Name()), "postgres", driver)
		if err != nil {
			logrus.Fatalf("migrate database init err: %v", err)
			return err
		}
		defer mg.Close()
		if err := mg.Up(); err != nil {
			if err != m.ErrNoChange {
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

func (p *postgresMigrate) MigrateDown(ctx context.Context) error {
	return nil
}

func (p *postgresMigrate) openDB(dbName string) (db *gorm.DB, err error) {
	dbConfig := pgsql.Config{}
	dbConfig.Database = dbName
	err = config.NewConfig(config.WithFilePath(p.configFilePath)).Register("pgsql", &dbConfig).Parse()
	if err != nil {
		return
	}
	return gorm.Open(pg.Open(dbConfig.Dsn()))
}

func (p *postgresMigrate) closeDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
