package migrate

import (
	"context"
	"strings"
)

type Migrate interface {
	InitDB(ctx context.Context) error
	MigrateUp(ctx context.Context) error
	MigrateDown(ctx context.Context) error
}

type Config struct {
	configFilePath     string
	initFilePath       string
	migrateRootDirPath string
}

type Option func(config *Config)

func WithConfigFilePath(configFilePath string) Option {
	return func(config *Config) {
		config.configFilePath = configFilePath
	}
}

func WithInitFilePath(initFilePath string) Option {
	return func(config *Config) {
		config.initFilePath = initFilePath
	}
}

func WithMigrateRootDirPath(migrateRootDirPath string) Option {
	return func(config *Config) {
		config.migrateRootDirPath = migrateRootDirPath
	}
}

func Process(migrates ...Migrate) error {
	ctx := context.Background()
	for _, m := range migrates {
		if err := m.InitDB(ctx); err != nil {
			return err
		}
		if err := m.MigrateUp(ctx); err != nil {
			return err
		}
		if err := m.MigrateDown(ctx); err != nil {
			return err
		}
	}
	return nil
}

func parseInitSqlFile(sqlStr string, specialStr string) map[string]string {
	var (
		sb                string //stringbuffer
		sqlSb             string //sql stringbuffer
		isComment         bool   //是否是注释
		dbNames           = make(map[string]string)
		currentDbName     string
		nextIsDBName      bool
		commentBeginIndex = -2
	)
	//解析0_init.sql文件中的database名称
	for i, r := range []rune(sqlStr) {
		if isComment {
			if string(r) == "\n" {
				isComment = false
				sqlSb += "\n"
			}
			continue
		}
		if string(r) == "-" {
			if commentBeginIndex+1 == i {
				isComment = true
			} else {
				commentBeginIndex = i
				continue
			}
		} else {
			if commentBeginIndex+1 == i {
				sb += "-"
				sqlSb += "-"
			}
		}
		if string(r) != "-" {
			sqlSb += string(r)
		}
		if string(r) != "\n" && string(r) != " " && string(r) != "-" && string(r) != ";" {
			sb += string(r)
		} else {
			if strings.ToLower(sb) == specialStr {
				nextIsDBName = true
			} else {
				if sb != "" {
					if nextIsDBName {
						nextIsDBName = false
						currentDbName = sb
					}
				}
			}
			sb = ""
			if string(r) == ";" {
				dbNames[currentDbName] = sqlSb
				sqlSb = ""
			}
		}
	}
	return dbNames
}
