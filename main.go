package main

import (
	"log"
	"os"
	"strconv"

	application_articles "example.com/mittaus/blog/application/articles"
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

	repoTag, repoArticle := loadRepo()
	serviceTag := application_tags.NewTagManager(repoTag)
	serviceArticle := application_articles.NewArticleManager(repoArticle)
	handler := server_routes.NewRouterHandler(serviceTag, serviceArticle)

	serverPort, _ := strconv.Atoi(os.Getenv("server.port"))

	log.Println("serverPort", serverPort)

	ginServer := config.NewServer(
		serverPort,
		config.DebugMode,
	)

	handler.SetRoutes(ginServer.Router)

	ginServer.Start()

	log.Println("server has started")
}

func loadRepo() (domain.ITagRepository, domain.IArticleRepository) {
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

	if databasePort, isFound = os.LookupEnv("database.port"); !isFound {
		databasePort = ""
	}

	db, error := mssql.NewRepositories(databaseDialect, databaseUser, databasePassword, databasePort, databaseServer, databaseDatabase)
	if error != nil {
		log.Println("Error al inicializar base de datos")
		panic(error)
	}
	repositoryTag := mssql.NewTagRepository(db)
	repositoryArticle := mssql.NewArticleRepository(db)
	return repositoryTag, repositoryArticle
}

func loadEnvConfig() {
	err := godotenv.Load("config.yaml")
	if err != nil {
		log.Println("Error loading .env file")
	}
}
