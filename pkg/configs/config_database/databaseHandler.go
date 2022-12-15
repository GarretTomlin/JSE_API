package config_database

import(
	"JSE_API/pkg/database"
)


func DbHandler(){
// Connect to the database
db, err := database.ConnectToDB("postgres",  "")
if err != nil {
	panic(err)
}

}
var Storage = &database.UrlStorage{Db: db, Cache: make(map[string]bool)}
