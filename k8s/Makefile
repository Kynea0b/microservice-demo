DIRS = api-gateway gacha item

# 一括インストール
.PHONY: apply $(addsuffix -apply,$(DIRS))

apply: $(addsuffix -apply,$(DIRS))

$(addsuffix -apply,$(DIRS)):
	$(MAKE) -C $(patsubst %-apply,%,$@) apply

# 一括アンインストール
.PHONY: delete $(addsuffix -delete,$(DIRS))

delete: $(addsuffix -delete,$(DIRS))

$(addsuffix -delete,$(DIRS)):
	$(MAKE) -C $(patsubst %-delete,%,$@) delete