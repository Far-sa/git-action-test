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

	dsn := "user:password@tcp(localhost:3306)/testdb?parseTime=true"

	dbWrapper, err := db.NewMySQLDB(dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Successfully connected to the database")

	// Access the underlying sql.DB using the DB method
	migrator := migrate.NewMySQLMigrator(dbWrapper.DB()) // Create migrator with the database connection

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "./database/migrations")

	err = migrator.Up(dir)

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
	log.Fatal(http.ListenAndServe(":8080", r))
}
