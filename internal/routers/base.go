package routers

import (
	"fmt"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/usecase"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/akhidnukhlis/implement-gRpc-server-author-service/config/app"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var grpcServer = grpc.NewServer()

type Service struct {
	Author *usecase.AuthorHandler
}

type Serve struct {
	DB      *gorm.DB
	Router  *mux.Router
	Service Service
}

// Initialize is used to initialize db driver connection
func (se *Serve) Initialize(dbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	//Set migration table
	registry := app.SetMigrationTable()

	//initiate database driver
	if dbDriver == "mysql" {
		dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		se.DB, err = gorm.Open(dbDriver, dbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", dbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", dbDriver)
		}
	}
	if dbDriver == "postgres" {
		dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
		se.DB, err = gorm.Open(dbDriver, dbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", dbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", dbDriver)
		}
	}

	//Migration process
	se.DB.Debug().AutoMigrate(registry...) //database migration

	se.Router = mux.NewRouter()

	se.initializeRoutes()
}

// Run is used to execute connection and run our services

func (se *Serve) Run() {
	port := os.Getenv("APP_PORT")
	fmt.Println("Listening to port ", port)
	log.Fatal(http.ListenAndServe(":"+port, se.Router))
}

func (se *Serve) RunGrpc() {
	port := os.Getenv("GRPC_PORT")
	fmt.Println("Listening to port ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed connection : %v\n", err)
	}

	log.Printf("server listening at %v\n", lis.Addr())
	if grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server : %v\n", err)
	}

	log.Fatal(http.ListenAndServe(":"+port, se.Router))
}
