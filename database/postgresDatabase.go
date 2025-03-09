package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/ittikorn/isekai-shop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	Once                      sync.Once
)

func NewPostgresDatabase(conf *config.Database) Database {
	Once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.SSLMode,
			conf.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Database connected %s", conf.DBName)
		postgresDatabaseInstance = &postgresDatabase{conn}
	})

	return postgresDatabaseInstance
}

// GetDB returns the GORM DB instance
func (db *postgresDatabase) ConnectionGetting() *gorm.DB {
	return db.DB
}