package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"restapi/app/controller"
	"restapi/app/models"
	"restapi/app/pagination"
	"restapi/app/repository"
	"restapi/app/services"
	"sync"
	"time"
)

var Once sync.Once
var EnvFile string
var configurationLoader *ConfigurationLoader

var paginator = new(pagination.Pagination).Default()

type ConfigurationLoader struct{}

func GetConfigurationLoaderInstance() *ConfigurationLoader {
	if configurationLoader == nil {
		Once.Do(func() { configurationLoader = new(ConfigurationLoader) })
	}
	return configurationLoader
}

func init() {
	EnvFile = getEnv()
}

func getEnv() string {
	return "app.dev.yaml"
}

func (configurationLoader *ConfigurationLoader) LoadConfiguration() (models.IAppConfig, error) {
	return models.NewAppConfig(EnvFile)
}

// ConnectDatabase ConnectDB establishes a connection to the MySQL database
func (configurationLoader *ConfigurationLoader) ConnectDatabase(config models.IAppConfig) (*gorm.DB, error) {
	// Get the values from the config struct
	user := config.Username()
	password := config.Password()
	host := config.Host()
	port := config.Port()
	dbName := config.Name()

	// Construct the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	dbConfig, _ := db.DB()
	dbConfig.SetMaxOpenConns(config.MaximumOpenConnection())
	dbConfig.SetConnMaxIdleTime(time.Duration(config.MaximumIdleTime()) * time.Second)
	dbConfig.SetConnMaxLifetime(time.Duration(config.MaximumTime()) * time.Second)
	dbConfig.SetMaxIdleConns(config.MaximumIdleConnection())

	log.Println("Running Migrations")
	configurationLoader.autoMigrate(db)

	log.Println("Connected to the database")
	return db, nil
}

// RegisterBookRoutes all routes dealing with books
func (configurationLoader *ConfigurationLoader) RegisterBookRoutes(c *controller.BookController, route *gin.Engine) {
	route.GET(models.BookResourcePath, paginator, c.All)
	route.POST(models.BookResourcePath, c.Create)
	route.GET(models.BookResourcePathID, c.Read)
	route.PUT(models.BookResourcePathID, c.Update)
	route.DELETE(models.BookResourcePathID, c.Delete)
}

// GetRouteHandler bootstrap book routes
func (configurationLoader *ConfigurationLoader) GetRouteHandler(gormDb *gorm.DB, config models.IAppConfig) *gin.Engine {
	middleware := GetMiddlewareInstance()
	route := gin.Default()

	route.Use(middleware.RequestLogger())
	route.Use(middleware.ResponseLogger())
	route.Use(middleware.RequestAuthorization(config))

	bookRepository := repository.NewBookRepository(gormDb)
	bookService := services.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	configurationLoader.RegisterBookRoutes(bookController, route)

	return route
}

// autoMigrate run DB migrations
func (configurationLoader *ConfigurationLoader) autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Println(err.Error())
		return
	}
}
