package service

import
_itemShopRepository "github.com/ittikorn/isekai-shop/pkg/itemShop/repository"


type itemShopRepositoryImpl  struct{
	ItemShopRepository _itemShopRepository.ItemShopRepository

}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	) ItemShopService {
	return &itemShopRepositoryImpl{ itemShopRepository}
}