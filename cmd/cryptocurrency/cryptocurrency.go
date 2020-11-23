package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/jmoiron/sqlx" 
	"github.com/seminarioGo/internal/config"
	"github.com/seminarioGo/internal/database"
	"github.com/seminarioGo/internal/service/cryptocurrency"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := readConfig()
	
	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	//Creacion de tabla
 	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := cryptocurrency.New(db, cfg)
	httpService := cryptocurrency.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config{
	configFile := flag.String("config","./config.yaml","this is the service config")
	flag.Parse() 


	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
	
		fmt.Println(err.Error())
		os.Exit(1) 
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS cryptocurrency (
		id integer primary key autoincrement,
		type varchar,
		quantity integer
		);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	
	return nil
} 