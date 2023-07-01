package api

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"restful_blog/di"
	"restful_blog/pkg/db"
	"restful_blog/pkg/routes"
)

var EnvConfig *config

func InitEnvConfigs() {
	EnvConfig = loadEnvVariables()
}

// struct to map env values
type config struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	USER            string `mapstructure:"MYSQL_USERNAME"`
	HOST            string `mapstructure:"MYSQL_HOST"`
	PORT            string `mapstructure:"MYSQL_PORT"`
	PASSWORD        string `mapstructure:"MYSQL_PASSWORD"`
	DBNAME          string `mapstructure:"MYSQL_DATABASE"`
}

func loadEnvVariables() (config *config) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("./")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

func Run() error {
	var closeDb error

	// Read banner.txt
	bs, err := os.ReadFile("cmd/api/banner.txt")
	if err != nil {
		return err
	}

	//print banner
	fmt.Println(string(bs))

	// Read env file
	InitEnvConfigs()

	Config := map[string]string{
		"USER":     EnvConfig.USER,
		"HOST":     EnvConfig.HOST,
		"PORT":     EnvConfig.PORT,
		"PASSWORD": EnvConfig.PASSWORD,
		"DBNAME":   EnvConfig.DBNAME,
	}

	// Initialize database connection
	db.InitializeDB(Config)

	//get db connection
	dbCon := db.GetDB()

	// Create new container app
	container := di.NewContainer(dbCon)

	// Create new instance echo
	e := echo.New()
	routes.Router(e, container)

	defer func() {
		if err := db.CloseDB(); err != nil {
			closeDb = err
		}
	}()

	// check closeDb has error or not
	if closeDb != nil {
		return closeDb
	}

	// Start app
	err = e.Start(EnvConfig.LocalServerPort)

	if err != nil {
		return err
	}

	return nil
}
