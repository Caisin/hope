package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"hope/apps/admin/internal/cache"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/migrate"
	"hope/pkg/conf"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
	rdb *redis.Client
}

func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "admin-service/data/ent"))
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	//性能检测
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		helper.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		helper.Fatalf("failed opening connection to db: %v", err)
		return nil
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		helper.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(entClient *ent.Client, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "admin-service/data"))

	d := &Data{
		db:  entClient,
		log: helper,
		rdb: rdb,
	}
	//初始化缓存
	err := cache.InitPermission(context.Background(), rdb, entClient)
	if err != nil {
		helper.Error(err)
		return nil, nil, err
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			helper.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			helper.Error(err)
		}
	}, nil
}
