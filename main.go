package main

import (
	"log"
	"os"
	"strconv"

	application_tags "example.com/mittaus/blog/application/tags"
	"example.com/mittaus/blog/config"
	domain "example.com/mittaus/blog/domain"
	server_routes "example.com/mittaus/blog/infraestructure/gin.server"
	mssql "example.com/mittaus/blog/infraestructure/persistence/mssql"
	"github.com/joho/godotenv"
)

func main() {
	//Cargamos los parámetros de config.yaml
	loadEnvConfig()

	repo := loadRepo()
	service := application_tags.NewTagManager(repo)
	handler := server_routes.NewRouterHandler(service)

	serverPort, _ := strconv.Atoi(os.Getenv("server.port"))
	ginServer := config.NewServer(
		serverPort,
		config.DebugMode,
	)

	handler.SetRoutes(ginServer.Router)

	ginServer.Start()
}

func loadRepo() domain.ITagRepository {
	var databaseDialect,
		databaseUser,
		databasePassword,
		databasePort,
		databaseServer,
		databaseDatabase string

	var isFound bool = false

	if databaseDialect, isFound = os.LookupEnv("database.dialect"); !isFound {
		databaseDialect = "sqlserver"
	}

	if databaseUser, isFound = os.LookupEnv("database.user"); !isFound {
		databaseUser = "*"
	}

	if databasePassword, isFound = os.LookupEnv("database.password"); !isFound {
		databasePassword = ""
	}

	if databaseServer, isFound = os.LookupEnv("database.server"); !isFound {
		databaseServer = ""
	}

	if databaseDatabase, isFound = os.LookupEnv("database.database"); !isFound {
		databaseDatabase = ""
	}

	if databasePort, isFound = os.LookupEnv("database.database"); !isFound {
		databasePort = ""
	}

	db, error := mssql.NewRepositories(databaseDialect, databaseUser, databasePassword, databasePort, databaseServer, databaseDatabase)
	if error != nil {
		panic(error)
	}
	repository := mssql.NewTagRepository(db)
	return repository
}

func loadEnvConfig() {
	err := godotenv.Load("config.yaml")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
