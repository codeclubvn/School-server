package cmd

import (
	"elearning/config"
	"elearning/infra/mysql"
	"github.com/sirupsen/logrus"
)

type App struct {
	config   *config.Environment
	database *mysql.Database
	logger   *logrus.Entry
}

func main() {

}
