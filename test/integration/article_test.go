package integration

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	application_articles "example.com/mittaus/blog/application/articles"
	"example.com/mittaus/blog/config"
	"example.com/mittaus/blog/domain"
	mssql "example.com/mittaus/blog/infraestructure/persistence/mssql"
	"github.com/joho/godotenv"
)

func TestGetEntries(t *testing.T) {

	loadTestEnvConfig()

	serverPort, _ := strconv.Atoi(os.Getenv("server.port"))
	ginServer := config.NewServer(
		serverPort,
		config.DebugMode,
	)

	repoArticle := loadRepo()
	serviceArticle := application_articles.NewArticleManager(repoArticle)
	ginServer.Router.GET("/api/v1/articles/:id", serviceArticle.Get)

	ts := httptest.NewServer(ginServer.Router)
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/v1/articles/2", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func loadRepo() domain.IArticleRepository {
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
		panic(error)
	}

	repositoryArticle := mssql.NewArticleRepository(db)
	return repositoryArticle
}

func loadTestEnvConfig() {
	err := godotenv.Load("../../config.test.yaml")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
