package main

import (
	itempb "api-gateway/pkg/grpc/proto"
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
	itemClient itempb.ItemServiceClient
	itemConn   *grpc.ClientConn
)

func init() {
	if itemClient != nil {
		return
	}

	host := os.Getenv("ITEM_SERVICE_HOST")
	if host == "" {
		log.Fatal("ITEM_SERVICE_HOST is not set")
	}

	var err error
	itemConn, err = grpc.Dial(
		fmt.Sprintf("%s:%s", host, "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	itemClient = itempb.NewItemServiceClient(itemConn)
}

func GetInventories(c echo.Context) error {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return err
	}

	req := &itempb.GetInventoryRequest{
		UserId: userId,
	}
	res, err := itemClient.GetInventory(context.Background(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func CloseItemConnection() {
	if itemConn != nil {
		itemConn.Close()
	}
}
