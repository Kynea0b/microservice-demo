package main

import (
	"context"
	"database/sql"
	gachapb "gacha-service/pkg/grpc/proto"
	"math/rand"
	"time"
)

type Item struct {
	Id     int64
	Name   string
	Rarity gachapb.Rarity
	Weight int
}

type gachaServiceServer struct {
	gachapb.UnimplementedGachaServiceServer
	db    *sql.DB
	items []Item
}

func NewGachaServiceServer(db *sql.DB, items []Item) gachapb.GachaServiceServer {
	return &gachaServiceServer{
		db:    db,
		items: items,
	}
}

func (g *gachaServiceServer) Draw(ctx context.Context, req *gachapb.DrawRequest) (*gachapb.DrawResponse, error) {
	// itemsからitemを重み付抽選する
	weights := make([]int, len(g.items))
	for i, item := range g.items {
		weights[i] = item.Weight
	}

	seed := time.Now().UnixNano()
	i := linearSearchLottery(weights, seed)
	item := g.items[i]

	// DBに保存する
	if err := save(ctx, g.db, req.UserId, item); err != nil {
		return nil, err
	}

	return &gachapb.DrawResponse{
		ItemId:   item.Id,
		ItemName: item.Name,
		Rarity:   item.Rarity,
	}, nil
}

func save(ctx context.Context, db *sql.DB, userId int64, item Item) error {
	_, err := db.ExecContext(
		ctx,
		"INSERT INTO Gacha (user_id, item_id, item_name, rarity, created_at) VALUES (?, ?, ?, ?, ?)",
		userId,
		item.Id,
		item.Name,
		item.Rarity.String(),
		time.Now().Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return err
	}
	return nil
}

/*
線形探索で重み付抽選する
@return 当選した要素のインデックス
*/
func linearSearchLottery(weights []int, seed int64) int {
	//  重みの総和を取得する
	var total int
	for _, weight := range weights {
		total += weight
	}

	// 乱数取得
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	rnd := rand.Intn(total)

	var currentWeight int
	for i, w := range weights {
		// 現在要素までの重みの総和
		currentWeight += w

		if rnd < currentWeight {
			return i
		}
	}

	// たぶんありえない
	return len(weights) - 1
}
