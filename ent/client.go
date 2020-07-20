// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/songfei1983/go-api-server/ent/migrate"

	"github.com/songfei1983/go-api-server/ent/adunit"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Adunit is the client for interacting with the Adunit builders.
	Adunit *AdunitClient
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
	c.Adunit = NewAdunitClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Adunit: NewAdunitClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Adunit: NewAdunitClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Adunit.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Adunit.Use(hooks...)
}

// AdunitClient is a client for the Adunit schema.
type AdunitClient struct {
	config
}

// NewAdunitClient returns a client for the Adunit from the given config.
func NewAdunitClient(c config) *AdunitClient {
	return &AdunitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `adunit.Hooks(f(g(h())))`.
func (c *AdunitClient) Use(hooks ...Hook) {
	c.hooks.Adunit = append(c.hooks.Adunit, hooks...)
}

// Create returns a create builder for Adunit.
func (c *AdunitClient) Create() *AdunitCreate {
	mutation := newAdunitMutation(c.config, OpCreate)
	return &AdunitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Adunit.
func (c *AdunitClient) Update() *AdunitUpdate {
	mutation := newAdunitMutation(c.config, OpUpdate)
	return &AdunitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AdunitClient) UpdateOne(a *Adunit) *AdunitUpdateOne {
	mutation := newAdunitMutation(c.config, OpUpdateOne, withAdunit(a))
	return &AdunitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AdunitClient) UpdateOneID(id int) *AdunitUpdateOne {
	mutation := newAdunitMutation(c.config, OpUpdateOne, withAdunitID(id))
	return &AdunitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Adunit.
func (c *AdunitClient) Delete() *AdunitDelete {
	mutation := newAdunitMutation(c.config, OpDelete)
	return &AdunitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AdunitClient) DeleteOne(a *Adunit) *AdunitDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AdunitClient) DeleteOneID(id int) *AdunitDeleteOne {
	builder := c.Delete().Where(adunit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AdunitDeleteOne{builder}
}

// Create returns a query builder for Adunit.
func (c *AdunitClient) Query() *AdunitQuery {
	return &AdunitQuery{config: c.config}
}

// Get returns a Adunit entity by its id.
func (c *AdunitClient) Get(ctx context.Context, id int) (*Adunit, error) {
	return c.Query().Where(adunit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AdunitClient) GetX(ctx context.Context, id int) *Adunit {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// Hooks returns the client hooks.
func (c *AdunitClient) Hooks() []Hook {
	return c.hooks.Adunit
}
