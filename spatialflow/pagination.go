package spatialflow

import "context"

// Pager provides iteration over paginated results.
// Uses offset-based pagination (matches SpatialFlow API).
type Pager[T any] struct {
	fetchPage func(ctx context.Context, offset, limit int) (items []T, total int, err error)
	offset    int
	limit     int
	total     int
	exhausted bool
}

// NewPager creates a new pager with the given fetch function.
// The fetch function should return (items, totalCount, error).
// For endpoints without total, return -1 and pager will use items length.
func NewPager[T any](limit int, fetch func(ctx context.Context, offset, limit int) ([]T, int, error)) *Pager[T] {
	if limit <= 0 {
		limit = 100
	}
	return &Pager[T]{
		fetchPage: fetch,
		limit:     limit,
		total:     -1, // Unknown until first fetch
	}
}

// Next fetches the next page of results.
// Returns nil, nil when exhausted.
func (p *Pager[T]) Next(ctx context.Context) ([]T, error) {
	if p.exhausted {
		return nil, nil
	}

	items, total, err := p.fetchPage(ctx, p.offset, p.limit)
	if err != nil {
		return nil, err
	}

	p.total = total
	p.offset += len(items)

	// Check if exhausted
	if len(items) == 0 || len(items) < p.limit {
		p.exhausted = true
	}
	// Also check total if known
	if p.total >= 0 && p.offset >= p.total {
		p.exhausted = true
	}

	return items, nil
}

// HasMore returns true if more pages may be available.
func (p *Pager[T]) HasMore() bool {
	return !p.exhausted
}

// Total returns the total count if known, or -1 if unknown.
func (p *Pager[T]) Total() int {
	return p.total
}

// CollectAll fetches all pages and returns all items.
func (p *Pager[T]) CollectAll(ctx context.Context) ([]T, error) {
	var all []T
	for p.HasMore() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		items, err := p.Next(ctx)
		if err != nil {
			return nil, err
		}
		all = append(all, items...)
	}
	return all, nil
}

// Iter returns a channel that yields items one at a time.
// The channel is closed when iteration completes or context is cancelled.
func (p *Pager[T]) Iter(ctx context.Context) <-chan IterResult[T] {
	ch := make(chan IterResult[T])
	go func() {
		defer close(ch)
		for p.HasMore() {
			items, err := p.Next(ctx)
			if err != nil {
				select {
				case ch <- IterResult[T]{Err: err}:
				case <-ctx.Done():
				}
				return
			}
			for _, item := range items {
				select {
				case ch <- IterResult[T]{Item: item}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return ch
}

// IterResult wraps an item or error from iteration.
type IterResult[T any] struct {
	Item T
	Err  error
}
