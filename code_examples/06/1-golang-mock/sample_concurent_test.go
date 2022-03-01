package golang_mock

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"gitlab.betsys.com/go-guild-mocks/1-golang-mock/mocks"
	"testing"
)

func call(ctx context.Context, qm *QuickMaths) (int, error) {
	result := make(chan int)
	go func() {
		result <- qm.Sum(1, 2)
		close(result)
	}()
	select {
	case r := <-result:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func TestQuickMathsConcurrent(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	defer ctrl.Finish()

	m := mocks.NewMockMath(ctrl)

	m.EXPECT().
		Sum(gomock.Eq(1), gomock.Eq(2)).
		Return(3)

	qm := &QuickMaths{
		Math: m,
	}

	res, err := call(ctx, qm)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if res != 3 {
		t.Errorf("Expected 3, got %d", res)
	}
}

func TestQuickMathsConcurrentError(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	defer ctrl.Finish()

	m := mocks.NewMockMath(ctrl)

	qm := &QuickMaths{
		Math: m,
	}

	res, err := call(ctx, qm)
	fmt.Printf("%v %v", res, err)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
