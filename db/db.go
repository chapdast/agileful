package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type Article struct {
	Id         int64  `json:"id"`
	Writer     string `json:"writer"`
	Body       string `json:"body"`
	CreateDate int64  `json:"create_date"`
	ViewCount  int64  `json:"view_count"`
}
type Operation int

const (
	OprEqual = Operation(iota)
	OprNotEqual
	OprSubString
)

var _ Database = database{}

type Condition struct {
	Value    interface{}
	Operator Operation
}
type option struct {
	limit     int
	offset    int
	fillerBy  map[string]Condition
	orderBy   [][]byte
	orderDesc bool
}

type Database interface {
	Raw() *pgxpool.Pool
	Create(ctx context.Context, article *Article) error
	Update(ctx context.Context, article *Article) error
	UpdateViewCount(ctx context.Context, articleID int) error
	Delete(ctx context.Context, articleID int) error
	Read(ctx context.Context, option ...Option) ([]*Article, error)
}

type database struct {
	conn *pgxpool.Pool
}

func NewDB(ctx context.Context, link string) (Database, error) {
	conn, err := pgxpool.Connect(ctx, link)
	if err != nil {
		return nil, err
	}
	return &database{
		conn: conn,
	}, nil
}
func GetDBLinkFromEnv() string {
	return os.Getenv("DB_CONNECTION_LINK")
}
func (d database) Raw() *pgxpool.Pool {
	return d.conn
}
