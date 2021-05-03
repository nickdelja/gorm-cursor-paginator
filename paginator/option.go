package paginator

var defaultConfig = Config{
	Keys:  []string{"ID"},
	Limit: 10,
	Order: DESC,
}

// Option for paginator
type Option interface {
	Apply(p *Paginator)
}

// Config for paginator
type Config struct {
	Keys   []string
	Limit  int
	Order  Order
	After  string
	Before string
}

// Apply applies config to paginator
func (c *Config) Apply(p *Paginator) {
	if len(c.Keys) != 0 {
		p.keys = c.Keys
	}
	if c.Limit > 0 {
		p.limit = c.Limit
	}
	if c.Order != "" {
		p.order = c.Order
	}
	if c.After != "" {
		p.cursor.After = &c.After
	}
	if c.Before != "" {
		p.cursor.Before = &c.Before
	}
}

// WithKeys configures keys for paginator
func WithKeys(keys ...string) Option {
	return &Config{
		Keys: keys,
	}
}

// WithLimit configures limit for paginator
func WithLimit(limit int) Option {
	return &Config{
		Limit: limit,
	}
}

// WithOrder configures order for paginator
func WithOrder(order Order) Option {
	return &Config{
		Order: order,
	}
}

// WithAfter configures after cursor for paginator
func WithAfter(c string) Option {
	return &Config{
		After: c,
	}
}

// WithBefore configures before cursor for paginator
func WithBefore(c string) Option {
	return &Config{
		Before: c,
	}
}
