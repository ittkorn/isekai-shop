package main

import (
    "github.com/ittikorn/isekai-shop/config"      
    "github.com/ittikorn/isekai-shop/databases" 
    "github.com/ittikorn/isekai-shop/entities"    
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)


	PlayerMigration(db)
	AdminMigration(db)
	ItemMigration(db)
	PlayerCoinMigration(db)
	InventoryMigration(db)
	PurchaseHistoryMigration(db)
}

func PlayerMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})
}

func AdminMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
}

func ItemMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
}

func PlayerCoinMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
}

func InventoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
}

func PurchaseHistoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
}
