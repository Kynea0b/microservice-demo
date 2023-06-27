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