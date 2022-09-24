// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"eshop/cmd/catalog/internal/data/ent/beer"
	"eshop/cmd/catalog/internal/data/ent/migrate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Beer is the client for interacting with the Beer builders.
	Beer *BeerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Beer = NewBeerClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Beer:   NewBeerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Beer:   NewBeerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Beer.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Beer.Use(hooks...)
}

// BeerClient is a client for the Beer schema.
type BeerClient struct {
	config
}

// NewBeerClient returns a client for the Beer from the given config.
func NewBeerClient(c config) *BeerClient {
	return &BeerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `beer.Hooks(f(g(h())))`.
func (c *BeerClient) Use(hooks ...Hook) {
	c.hooks.Beer = append(c.hooks.Beer, hooks...)
}

// Create returns a create builder for Beer.
func (c *BeerClient) Create() *BeerCreate {
	mutation := newBeerMutation(c.config, OpCreate)
	return &BeerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Beer entities.
func (c *BeerClient) CreateBulk(builders ...*BeerCreate) *BeerCreateBulk {
	return &BeerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Beer.
func (c *BeerClient) Update() *BeerUpdate {
	mutation := newBeerMutation(c.config, OpUpdate)
	return &BeerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BeerClient) UpdateOne(b *Beer) *BeerUpdateOne {
	mutation := newBeerMutation(c.config, OpUpdateOne, withBeer(b))
	return &BeerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BeerClient) UpdateOneID(id int64) *BeerUpdateOne {
	mutation := newBeerMutation(c.config, OpUpdateOne, withBeerID(id))
	return &BeerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Beer.
func (c *BeerClient) Delete() *BeerDelete {
	mutation := newBeerMutation(c.config, OpDelete)
	return &BeerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BeerClient) DeleteOne(b *Beer) *BeerDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BeerClient) DeleteOneID(id int64) *BeerDeleteOne {
	builder := c.Delete().Where(beer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BeerDeleteOne{builder}
}

// Query returns a query builder for Beer.
func (c *BeerClient) Query() *BeerQuery {
	return &BeerQuery{config: c.config}
}

// Get returns a Beer entity by its id.
func (c *BeerClient) Get(ctx context.Context, id int64) (*Beer, error) {
	return c.Query().Where(beer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BeerClient) GetX(ctx context.Context, id int64) *Beer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BeerClient) Hooks() []Hook {
	return c.hooks.Beer
}
