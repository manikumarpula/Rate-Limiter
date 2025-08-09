package limiter

// Factory provides methods to create different rate limiter implementations
type Factory struct{}

// NewFactory creates a new limiter factory
func NewFactory() *Factory {
	return &Factory{}
}

// TODO: Add methods to create different algorithm implementations
// - CreateTokenBucket()
// - CreateFixedWindow()
// - CreateSlidingWindow()
// - CreateLeakyBucket()