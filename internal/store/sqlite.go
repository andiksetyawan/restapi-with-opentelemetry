package store

import (
	"github.com/rs/zerolog/log"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"restapi-with-opentelemetry/internal/entity"
)

func NewSQLLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect sqllite database")
		return nil
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatal().Err(err).Msg("failed to set otelgorm middleware")
		return nil
	}

	//auto migrate user, article entity
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Article{})
	return db
}
