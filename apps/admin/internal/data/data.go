package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"hope/apps/admin/internal/conf"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/migrate"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewEntClient, NewData, NewGreeterRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "admin-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		helper.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		helper.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "admin-service/data"))

	d := &Data{
		db:  entClient,
		log: helper,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			helper.Error(err)
		}
	}, nil
}
