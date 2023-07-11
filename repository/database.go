package repository

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Client *gorm.DB
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "database_name"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_DRIVER   = "mysql"
)

// InitDb initialize database connection client
func InitDb() error {
	return connectDB()
}

func connectDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT)

	err := CheckAndCreateDatabase(dsn, DB_NAME)

	if err != nil {
		fmt.Println("Error al verificar o crear la base de datos:", err)
	} else {
		fmt.Println("Verificación y creación de la base de datos completadas con éxito.")
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Set to true to avoid pluralization
		},
	})

	if err != nil {
		fmt.Println("fallo en realizar la conexion a la base de datos MySQL")
		return err
	}

	Client = db

	fmt.Println("Conexión exitosa a la base de datos MySQL")

	// Read the table creation query
	file, err := os.Open("repository/sql.sql")

	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return err
	}

	defer file.Close()

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return err
	}

	// Run the query
	err = Client.Exec(string(content)).Error

	if err != nil {
		fmt.Println("Error al crear la/s tabla/s")
		return err
	}

	fmt.Printf("Tabla/s creada/s correctamente\n")

	return nil
}

// DisconnectDB closes all database client connections
func DisconnectDB() {
	if Client != nil {
		db, _ := Client.DB()
		db.Close()
	}
}

// Check if database exist
func CheckAndCreateDatabase(dbURL, dbName string) error {
	db, err := sql.Open(DB_DRIVER, dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DB_NAME)
	if err != nil {
		return err
	}

	return nil
}