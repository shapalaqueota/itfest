package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectDB() {
	var err error
	databaseUrl := "postgresql://neondb_owner:hJavESpN2s6I@ep-white-base-a5jlyjwr.us-east-2.aws.neon.tech/neondb?sslmode=require"
	log.Println("Database URL:", databaseUrl)
	DB, err = pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v \n", err))
	}

	log.Println("Successfully connected to database")
	createTables()
}

func createTables() {
	createUserTable := `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    
    CREATE TABLE IF NOT EXISTS "user" (
      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      username VARCHAR(255) NOT NULL,
      email VARCHAR(255) UNIQUE NOT NULL,
      password VARCHAR(255) NOT NULL,
      phone_number VARCHAR(255) UNIQUE
    );`

	_, err := DB.Exec(context.Background(), createUserTable)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create user table: %v \n", err))
	}

	createItemTable := `
    CREATE TABLE IF NOT EXISTS item (
      id SERIAL PRIMARY KEY,
      title VARCHAR(255) NOT NULL,
      description TEXT,
      price VARCHAR(2555) NOT NULL,
      category VARCHAR(255) NOT NULL,
      image_url VARCHAR(255)
    );`

	_, err = DB.Exec(context.Background(), createItemTable)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create item table: %v \n", err))
	}
}
