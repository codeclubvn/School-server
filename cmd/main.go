package cmd

import (
	"elearning/config"
	"elearning/infra/mysql"
	"elearning/infra/mysql/repository"
	"elearning/pkg/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"reflect"
)

type App struct {
	config   *config.Environment
	database *mysql.Database
	logger   *logrus.Entry
}

var listEnvSecret = []string{
	"Constants",
	"MysqlPassword",
	"RedisPassword",
	"AccessTokenSecretKey",
	"RefreshTokenSecretKey",
	"AwsSecretAccessKey",
}

func main() {
	logger := initLog()
	loggerStartServer := initStartServerLog()
	cfg := loadEnvironment()
	v := reflect.ValueOf(cfg).Elem()
	for i := 0; i < v.NumField(); i++ {
		varName := v.Type().Field(i).Name
		varValue := v.Field(i).Interface()
		isLog := true
		for _, envSecret := range listEnvSecret {
			if varName == envSecret {
				isLog = false
				break
			}
		}
		if isLog {
			fmt.Printf("EnvKeyAndValue %s: '%v'\n", varName, varValue)
		}
	}

	gin.SetMode(cfg.RunMode)
	loggerStartServer.Infof("System is running with %s mode", cfg.RunMode)
	// Connect to database
	db, err := mysql.ConnectPostgresql(cfg)
	if err != nil {
		loggerStartServer.Fatalf("Connect PostgresSQL Server Failed. Error: %s", err.Error())
	}
	loggerStartServer.Infof("Connect PostgresSQL Server Successfully")

	app := &App{
		config:   cfg,
		database: db,
		logger:   logger,
	}

	if app.config.MysqlMigrateMode {
		loggerStartServer.Info("Start Migrating")
		err = app.database.AutoMigrate()
		if err != nil {
			loggerStartServer.Fatal("Migrating Failed. Error: " + err.Error())
		}
		loggerStartServer.Info("Migrating Successfully")
	}

	// Service
	dataService := data.NewDataService()

	// Repository
	userRepository := repository.NewUserRepository(app.database, dataService)
}

func initLog() *logrus.Entry {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		DisableColors:   true,
		FieldMap: logrus.FieldMap{
			"level": "logLevel",
		},
	})
	log := logrus.WithFields(logrus.Fields{
		"module": "backend",
	})
	return log
}

func initStartServerLog() *logrus.Entry {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		DisableColors:   true,
		DisableSorting:  true,
		FieldMap: logrus.FieldMap{
			"level": "logLevel",
		},
	})
	log := logrus.WithFields(logrus.Fields{
		"module":  "backend",
		"logType": "startServer",
	})
	return log
}

func loadEnvironment() *config.Environment {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Fail loading environment variables: ", err)
	}
	return cfg
}
