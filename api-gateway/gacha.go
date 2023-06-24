package main

import (
	gachapb "api-gateway/pkg/grpc/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client gachapb.GachaServiceClient

func init() {
	if client != nil {
		return
	}

	host := os.Getenv("GACHA_SERVICE_HOST")
	if host == "" {
		log.Fatal("GACHA_SERVICE_HOST is not set")
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", host, "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	client = gachapb.NewGachaServiceClient(conn)
}

type DrawRequest struct {
	UserId int64 `json:"user_id"`
}

// /draw
func Draw(c echo.Context) error {
	dr := new(DrawRequest)
	if err := c.Bind(dr); err != nil {
		return err
	}

	req := &gachapb.DrawRequest{
		UserId: dr.UserId,
	}
	res, err := client.Draw(context.Background(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
