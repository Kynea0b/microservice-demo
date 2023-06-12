package main

import (
	"context"
	gachapb "gacha-service/pkg/grpc/proto"
)

type gachaServiceServer struct {
	gachapb.UnimplementedGachaServiceServer
}

func NewGachaServiceServer() gachapb.GachaServiceServer {
	return &gachaServiceServer{}
}

func (g *gachaServiceServer) Draw(ctx context.Context, req *gachapb.DrawRequest) (*gachapb.DrawResponse, error) {
	return &gachapb.DrawResponse{
		ItemId:   1,
		ItemName: "item1",
		Rarity:   gachapb.Rarity_RARITY_N,
	}, nil
}
