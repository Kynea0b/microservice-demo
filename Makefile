DIRS = api-gateway gacha-service item-service

# 一括ビルド
.PHONY: build $(addsuffix -build,$(DIRS))

build: $(addsuffix -build,$(DIRS))

$(addsuffix -build,$(DIRS)):
	$(MAKE) -C $(patsubst %-build,%,$@) build

# 一括でproto更新
.PHONY: proto $(addsuffix -proto,$(DIRS))

proto: $(addsuffix -proto,$(DIRS))

$(addsuffix -proto,$(DIRS)):
	$(MAKE) -C $(patsubst %-proto,%,$@) proto

# ガチャを引く
draw:
	@curl -XPOST -H 'Content-Type: application/json' -d '{"user_id": $(USER_ID)}' localhost:$(PORT)/draw

# ガチャ履歴
history:
	@curl localhost:$(PORT)/histories/$(USER_ID)

# アイテムボックス取得
inventry:
	@curl localhost:$(PORT)/inventories/$(USER_ID)