package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"restapi-with-opentelemetry/internal/entity"
)

type postgreSQLDb struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func (p *postgreSQLDb) getDsnString() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", p.host, p.port, p.user, p.dbName, p.password)
	log.Debug().Msg(dsn)
	return dsn
}

func NewPostgreSQLDb() *postgreSQLDb {
	return &postgreSQLDb{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
	}
}

//NewPostgreeSQLDbTest creating database server for unit test with testcontainers
func NewPostgreeSQLDbTest() *postgreSQLDb {
	ctx := context.TODO()

	sqlDb := postgreSQLDb{
		user:     "postgres",
		password: "postgres",
		dbName:   "test_db",
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres:13-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       sqlDb.dbName,
			"POSTGRES_USER":     sqlDb.user,
			"POSTGRES_PASSWORD": sqlDb.password,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithPollInterval(1 * time.Second),
	}

	containers, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create test container, please check 5432 port on this OS")
	}

	port, err := containers.Ports(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to get host test container")
	}
	sqlDb.host = port["5432/tcp"][0].HostIP
	sqlDb.port = port["5432/tcp"][0].HostPort

	return &sqlDb
}

func (p *postgreSQLDb) Connect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: p.getDsnString()}), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open db connection")
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatal().Err(err).Msg("failed to set otelgorm middleware")
		return nil
	}

	db.AutoMigrate(&entity.User{})
	return db
}
