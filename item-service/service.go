package main

import (
	"context"
	"database/sql"
	itempb "item-service/pkg/grpc/proto"
	"time"
)

type itemServiceServer struct {
	itempb.UnimplementedItemServiceServer
	db *sql.DB
}

type Inventory struct {
	Id        int64
	UserId    int64
	ItemId    int64
	ItemName  string
	Rarity    string
	Count     int
	CreatedAt time.Time
}

func NewItemServiceServer(db *sql.DB) itempb.ItemServiceServer {
	return &itemServiceServer{
		db: db,
	}
}

func (s *itemServiceServer) GetItem(ctx context.Context, req *itempb.GetItemRequest) (*itempb.GetItemResponse, error) {
	// アイテムを所持しているか確認
	inventry, err := get(req.UserId, req.ItemId, s.db)
	if err = handleError(err); err != nil {
		return nil, err
	}

	// DB更新
	if inventry.Id == 0 {
		// 未所持
		if _, err = insert(req.UserId, req.ItemId, req.ItemName, req.Rarity.String(), s.db); err != nil {
			return nil, err
		}
	} else {
		// 所持済み
		if err = update(req.UserId, req.ItemId, s.db); err != nil {
			return nil, err
		}
	}

	return &itempb.GetItemResponse{
		ItemId:   req.ItemId,
		ItemName: req.ItemName,
		Rarity:   req.Rarity,
		Count:    int32(inventry.Count) + 1,
	}, nil
}

func get(userId, itemId int64, db *sql.DB) (Inventory, error) {
	row := db.QueryRow("SELECT id FROM inventry WHERE user_id = ? AND item_id = ?", userId, itemId)

	var inventry Inventory
	if err := row.Scan(
		&inventry.Id,
		&inventry.UserId,
		&inventry.ItemId,
		&inventry.ItemName,
		&inventry.Rarity,
		&inventry.Count,
		&inventry.CreatedAt,
	); err != nil {
		return Inventory{}, err
	}
	return inventry, nil
}

func handleError(err error) error {
	if err == nil || err == sql.ErrNoRows {
		return nil
	}

	return err
}

func insert(userId, itemId int64, itemName, rarity string, db *sql.DB) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO inventry (user_id, item_id, item_name, rarity, count, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		userId,
		itemId,
		itemName,
		rarity,
		1,
		time.Now().Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func update(userId, itemId int64, db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE inventry SET count = count + 1 WHERE user_id = ? AND item_id = ?",
		userId,
		itemId,
	)
	if err != nil {
		return err
	}
	return nil
}
