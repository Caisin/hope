// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Activity is the client for interacting with the Activity builders.
	Activity *ActivityClient
	// ActivityComponent is the client for interacting with the ActivityComponent builders.
	ActivityComponent *ActivityComponentClient
	// AdChangeLog is the client for interacting with the AdChangeLog builders.
	AdChangeLog *AdChangeLogClient
	// AdChannel is the client for interacting with the AdChannel builders.
	AdChannel *AdChannelClient
	// AgreementLog is the client for interacting with the AgreementLog builders.
	AgreementLog *AgreementLogClient
	// AmBalance is the client for interacting with the AmBalance builders.
	AmBalance *AmBalanceClient
	// AppVersion is the client for interacting with the AppVersion builders.
	AppVersion *AppVersionClient
	// AssetChangeLog is the client for interacting with the AssetChangeLog builders.
	AssetChangeLog *AssetChangeLogClient
	// AssetItem is the client for interacting with the AssetItem builders.
	AssetItem *AssetItemClient
	// BookPackage is the client for interacting with the BookPackage builders.
	BookPackage *BookPackageClient
	// ClientError is the client for interacting with the ClientError builders.
	ClientError *ClientErrorClient
	// CustomerNovelConfig is the client for interacting with the CustomerNovelConfig builders.
	CustomerNovelConfig *CustomerNovelConfigClient
	// CustomerNovels is the client for interacting with the CustomerNovels builders.
	CustomerNovels *CustomerNovelsClient
	// DataSource is the client for interacting with the DataSource builders.
	DataSource *DataSourceClient
	// ListenRecord is the client for interacting with the ListenRecord builders.
	ListenRecord *ListenRecordClient
	// Novel is the client for interacting with the Novel builders.
	Novel *NovelClient
	// NovelAutoBuy is the client for interacting with the NovelAutoBuy builders.
	NovelAutoBuy *NovelAutoBuyClient
	// NovelBookshelf is the client for interacting with the NovelBookshelf builders.
	NovelBookshelf *NovelBookshelfClient
	// NovelBuyChapterRecord is the client for interacting with the NovelBuyChapterRecord builders.
	NovelBuyChapterRecord *NovelBuyChapterRecordClient
	// NovelBuyRecord is the client for interacting with the NovelBuyRecord builders.
	NovelBuyRecord *NovelBuyRecordClient
	// NovelChapter is the client for interacting with the NovelChapter builders.
	NovelChapter *NovelChapterClient
	// NovelClassify is the client for interacting with the NovelClassify builders.
	NovelClassify *NovelClassifyClient
	// NovelComment is the client for interacting with the NovelComment builders.
	NovelComment *NovelCommentClient
	// NovelConsume is the client for interacting with the NovelConsume builders.
	NovelConsume *NovelConsumeClient
	// NovelMsg is the client for interacting with the NovelMsg builders.
	NovelMsg *NovelMsgClient
	// PayOrder is the client for interacting with the PayOrder builders.
	PayOrder *PayOrderClient
	// SocialUser is the client for interacting with the SocialUser builders.
	SocialUser *SocialUserClient
	// TaskLog is the client for interacting with the TaskLog builders.
	TaskLog *TaskLogClient
	// UserEvent is the client for interacting with the UserEvent builders.
	UserEvent *UserEventClient
	// UserMsg is the client for interacting with the UserMsg builders.
	UserMsg *UserMsgClient
	// VipUser is the client for interacting with the VipUser builders.
	VipUser *VipUserClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Activity = NewActivityClient(tx.config)
	tx.ActivityComponent = NewActivityComponentClient(tx.config)
	tx.AdChangeLog = NewAdChangeLogClient(tx.config)
	tx.AdChannel = NewAdChannelClient(tx.config)
	tx.AgreementLog = NewAgreementLogClient(tx.config)
	tx.AmBalance = NewAmBalanceClient(tx.config)
	tx.AppVersion = NewAppVersionClient(tx.config)
	tx.AssetChangeLog = NewAssetChangeLogClient(tx.config)
	tx.AssetItem = NewAssetItemClient(tx.config)
	tx.BookPackage = NewBookPackageClient(tx.config)
	tx.ClientError = NewClientErrorClient(tx.config)
	tx.CustomerNovelConfig = NewCustomerNovelConfigClient(tx.config)
	tx.CustomerNovels = NewCustomerNovelsClient(tx.config)
	tx.DataSource = NewDataSourceClient(tx.config)
	tx.ListenRecord = NewListenRecordClient(tx.config)
	tx.Novel = NewNovelClient(tx.config)
	tx.NovelAutoBuy = NewNovelAutoBuyClient(tx.config)
	tx.NovelBookshelf = NewNovelBookshelfClient(tx.config)
	tx.NovelBuyChapterRecord = NewNovelBuyChapterRecordClient(tx.config)
	tx.NovelBuyRecord = NewNovelBuyRecordClient(tx.config)
	tx.NovelChapter = NewNovelChapterClient(tx.config)
	tx.NovelClassify = NewNovelClassifyClient(tx.config)
	tx.NovelComment = NewNovelCommentClient(tx.config)
	tx.NovelConsume = NewNovelConsumeClient(tx.config)
	tx.NovelMsg = NewNovelMsgClient(tx.config)
	tx.PayOrder = NewPayOrderClient(tx.config)
	tx.SocialUser = NewSocialUserClient(tx.config)
	tx.TaskLog = NewTaskLogClient(tx.config)
	tx.UserEvent = NewUserEventClient(tx.config)
	tx.UserMsg = NewUserMsgClient(tx.config)
	tx.VipUser = NewVipUserClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Activity.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
