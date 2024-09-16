package pkgctx

import "context"

type ctxKeyBag struct{}

// Bag contains typed keys, values to be included as [context.WithValue].
type Bag struct {
	RequestID string
}

// WithBag inserts new [*Bag] into ctx and returns the new bag-inserted ctx.
func WithBag(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyBag{}, &Bag{})
}

// GetBag returns [*Bag] from ctx with true if exists. Otherwise, nil with true.
func GetBag(ctx context.Context) (*Bag, bool) {
	bag, ok := ctx.Value(ctxKeyBag{}).(*Bag)
	return bag, ok
}
