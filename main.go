package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"
	"test-cicd/repository/db"
	repository "test-cicd/repository/mysql"
	"test-cicd/repository/mysql/migrate"
	"test-cicd/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if err := userService.RegisterUser(username, email, password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	fmt.Println("server is starting..")
	log.Fatal(http.ListenAndServe(":8081", r))
}
