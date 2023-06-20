package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	gachapb "gacha-service/pkg/grpc/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3306"
	db   = "Gacha"
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

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
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
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	log.Println("Start Gacha Service")

	dsn, err := getDatasourceName()
	if err != nil {
		log.Fatal(err)
	}

	db, err := connectDb(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// grpc server の起動
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	reflection.Register(server)

	// ガチャの登録
	items := []Item{
		{Id: 1, Name: "item1", Rarity: gachapb.Rarity_RARITY_N, Weight: 40},
		{Id: 2, Name: "item2", Rarity: gachapb.Rarity_RARITY_R, Weight: 30},
		{Id: 3, Name: "item3", Rarity: gachapb.Rarity_RARITY_SR, Weight: 20},
		{Id: 4, Name: "item4", Rarity: gachapb.Rarity_RARITY_SSR, Weight: 10},
	}
	gachapb.RegisterGachaServiceServer(server, NewGachaServiceServer(db, items))

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
