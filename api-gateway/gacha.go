package main

import (
	gachapb "api-gateway/pkg/grpc/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	gachaClient gachapb.GachaServiceClient
	gachaConn   *grpc.ClientConn
)

func init() {
	if gachaClient != nil {
		return
	}

	host := os.Getenv("GACHA_SERVICE_HOST")
	if host == "" {
		log.Fatal("GACHA_SERVICE_HOST is not set")
	}

	var err error
	gachaConn, err = grpc.Dial(
		fmt.Sprintf("%s:%s", host, "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	gachaClient = gachapb.NewGachaServiceClient(gachaConn)
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
	res, err := gachaClient.Draw(context.Background(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func GetHistories(c echo.Context) error {
	p := c.Param("user_id")
	userId, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return err
	}

	req := &gachapb.GetHistoriesRequest{
		UserId: userId,
	}
	res, err := gachaClient.GetHistories(context.Background(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func CloseGachaConnection() error {
	if gachaConn != nil {
		return gachaConn.Close()
	}
	return nil
}
