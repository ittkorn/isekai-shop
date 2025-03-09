package comtroller

import (
	_itemShopService "github.com/ittikorn/isekai-shop/pkg/itemShop/service"
)

type ItemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(
	itemShopService _itemShopService.ItemShopService,
) ItemShopController {
	return &ItemShopControllerImpl{itemShopService}
}