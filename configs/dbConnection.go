package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetEnv(key, defaultValue string,stag string) string {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if(stag=="dev"){
		viper.SetConfigFile(".env.dev")
	}else{
		viper.SetConfigFile(".env.prod")
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}

func InitDB() (*sql.DB, error, string,string) {
	stag := GetStagEnv("STAG","prod")
	dbUser := GetEnv("DB_USER","youruser",stag)
	dbPass := GetEnv("DB_PASS","yourpass",stag)
	dbHost := GetEnv("DB_HOST","yourhost",stag)
	dbPort := GetEnv("DB_PORT","3306",stag)
	dbName := GetEnv("DB_NAME","linkaja",stag)
	dbEngine := GetEnv("DB_ENGINE","mysql",stag)
	hostServer := GetEnv("MAIN_SERVER_HOST","yourhost",stag)
	portServer := GetEnv("MAIN_SERVER_PORT","8080",stag)

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(source)
	db, _ := sql.Open(dbEngine, source)
	err := db.Ping()
	if err != nil {
		log.Print(err)
	}else{
		fmt.Println("Database Connected")
	}

	return db, err, hostServer, portServer

}


func GetStagEnv(key, defaultValue string) string {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}