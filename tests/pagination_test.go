package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/spatialflow-io/spatialflow-go/spatialflow"
)

func TestPager_Next(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}
	fetchCalls := 0

	pager := spatialflow.NewPager(2, func(ctx context.Context, offset, limit int) ([]string, int, error) {
		fetchCalls++
		end := offset + limit
		if end > len(items) {
			end = len(items)
		}
		if offset >= len(items) {
			return nil, len(items), nil
		}
		return items[offset:end], len(items), nil
	})

	ctx := context.Background()

	// First page
	page1, err := pager.Next(ctx)
	if err != nil {
		t.Fatalf("Next() error = %v", err)
	}
	if len(page1) != 2 || page1[0] != "a" || page1[1] != "b" {
		t.Errorf("page1 = %v, want [a, b]", page1)
	}

	// Second page
	page2, err := pager.Next(ctx)
	if err != nil {
		t.Fatalf("Next() error = %v", err)
	}
	if len(page2) != 2 || page2[0] != "c" || page2[1] != "d" {
		t.Errorf("page2 = %v, want [c, d]", page2)
	}

	// Third page (partial)
	page3, err := pager.Next(ctx)
	if err != nil {
		t.Fatalf("Next() error = %v", err)
	}
	if len(page3) != 1 || page3[0] != "e" {
		t.Errorf("page3 = %v, want [e]", page3)
	}

	// Exhausted
	if pager.HasMore() {
		t.Error("HasMore() = true, want false")
	}

	page4, err := pager.Next(ctx)
	if err != nil {
		t.Fatalf("Next() error = %v", err)
	}
	if page4 != nil {
		t.Errorf("page4 = %v, want nil", page4)
	}
}

func TestPager_CollectAll(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7}

	pager := spatialflow.NewPager(3, func(ctx context.Context, offset, limit int) ([]int, int, error) {
		end := offset + limit
		if end > len(items) {
			end = len(items)
		}
		if offset >= len(items) {
			return nil, len(items), nil
		}
		return items[offset:end], len(items), nil
	})

	ctx := context.Background()
	all, err := pager.CollectAll(ctx)
	if err != nil {
		t.Fatalf("CollectAll() error = %v", err)
	}

	if len(all) != 7 {
		t.Errorf("len(all) = %d, want 7", len(all))
	}

	for i, v := range all {
		if v != i+1 {
			t.Errorf("all[%d] = %d, want %d", i, v, i+1)
		}
	}
}

func TestPager_Iter(t *testing.T) {
	items := []string{"x", "y", "z"}

	pager := spatialflow.NewPager(2, func(ctx context.Context, offset, limit int) ([]string, int, error) {
		end := offset + limit
		if end > len(items) {
			end = len(items)
		}
		if offset >= len(items) {
			return nil, len(items), nil
		}
		return items[offset:end], len(items), nil
	})

	ctx := context.Background()
	var collected []string

	for result := range pager.Iter(ctx) {
		if result.Err != nil {
			t.Fatalf("Iter() error = %v", result.Err)
		}
		collected = append(collected, result.Item)
	}

	if len(collected) != 3 {
		t.Errorf("len(collected) = %d, want 3", len(collected))
	}
}

func TestPager_ErrorHandling(t *testing.T) {
	expectedErr := errors.New("fetch error")
	callCount := 0

	pager := spatialflow.NewPager(10, func(ctx context.Context, offset, limit int) ([]string, int, error) {
		callCount++
		if callCount == 2 {
			return nil, 0, expectedErr
		}
		// Return full page so pager doesn't think we're exhausted
		items := make([]string, 10)
		for i := range items {
			items[i] = "item"
		}
		return items, 100, nil // Pretend there are more
	})

	ctx := context.Background()

	// First call succeeds
	_, err := pager.Next(ctx)
	if err != nil {
		t.Fatalf("first Next() should succeed: %v", err)
	}

	// Second call fails
	_, err = pager.Next(ctx)
	if err != expectedErr {
		t.Errorf("second Next() error = %v, want %v", err, expectedErr)
	}
}

func TestPager_ContextCancellation(t *testing.T) {
	pager := spatialflow.NewPager(10, func(ctx context.Context, offset, limit int) ([]string, int, error) {
		return []string{"item"}, 1000, nil // Many items
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := pager.CollectAll(ctx)
	if err != context.Canceled {
		t.Errorf("CollectAll() error = %v, want context.Canceled", err)
	}
}

func TestPager_Total(t *testing.T) {
	pager := spatialflow.NewPager(10, func(ctx context.Context, offset, limit int) ([]string, int, error) {
		return []string{"a", "b"}, 42, nil
	})

	// Before first fetch, total is unknown
	if pager.Total() != -1 {
		t.Errorf("Total() before fetch = %d, want -1", pager.Total())
	}

	ctx := context.Background()
	_, _ = pager.Next(ctx)

	// After fetch, total is known
	if pager.Total() != 42 {
		t.Errorf("Total() after fetch = %d, want 42", pager.Total())
	}
}
