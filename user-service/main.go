package main

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"user-service/handler"
	"user-service/repository/db"
	repository "user-service/repository/mysql"
	"user-service/repository/mysql/migrate"
	"user-service/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
)

func main() {

	// dsn := "user:password@tcp(mysql:3306)/testdb?parseTime=true"

	dbWrapper, err := db.NewMySQLDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer dbWrapper.DB().Close()

	log.Println("Successfully connected to the database")

	// create  path
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not determine current file path")
	}

	migrationsDir := path.Join(path.Dir(filename), "database/migrations")
	log.Printf("Migrations directory: %s", migrationsDir)

	// Access the underlying sql.DB using the DB method
	migrator := migrate.NewMySQLMigrator(dbWrapper.DB()) // Create migrator with the database connection
	err = migrator.Up(migrationsDir)

	if err != nil {
		fmt.Println("Error running migrations up:", err)
		return
	}

	userRepo := repository.NewMySQLUserRepository(dbWrapper)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()

	// Define routes
	e.POST("/register", userHandler.Register)

	// Start server
	fmt.Println("server is starting..")
	e.Start(":3000")

}
