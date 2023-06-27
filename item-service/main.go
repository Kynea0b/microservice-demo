package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	itempb "item-service/pkg/grpc/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "3306"
	db   = "Item"
)

func getDatasourceName() (string, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		return "", fmt.Errorf("DB_USER is not set")
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return "", fmt.Errorf("DB_PASSWORD is not set")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", fmt.Errorf("DB_HOST is not set")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		pass,
		host,
		port,
		db,
	), nil
}

func connectDb(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("failed to connect database")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("failed to ping")
		return nil, err
	}
	return db, nil
}

func createTable(db *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS inventry (
		id BIGINT PRIMARY KEY AUTO_INCREMENT,
		user_id BIGINT NOT NULL,
		item_id BIGINT NOT NULL,
		item_name VARCHAR(255) NOT NULL,
		rarity VARCHAR(20) NOT NULL,
		count INT NOT NULL,
		created_at DATETIME NOT NULL
	)
	`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.Println("Start Item Service")

	dsn, err := getDatasourceName()
	if err != nil {
		log.Fatal(err)
	}

	db, err := connectDb(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	// grpc server の起動
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	reflection.Register(server)

	itempb.RegisterItemServiceServer(server, NewItemServiceServer(db))

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	server.GracefulStop()
}
