package calculator

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/adenix/go-grpc-boilerplate/gen/go/calculatorpb/v1"
)

func TestSum(t *testing.T) {
	c := NewCalculatorServiceServer()

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		var a int32 = -300 + int32(rand.Intn(600))
		var b int32 = -100 + int32(rand.Intn(600))

		go func(a, b int32) {
			defer wg.Done()

			t.Run(fmt.Sprintf("sum-random-%d", i), func(t *testing.T) {
				res, err := c.Sum(context.Background(), &calculatorpb.SumRequest{A: a, B: b})
				if err != nil {
					t.Fatalf("failed to sum request: %q", err)
				}

				expected := a + b
				if r := res.GetResult(); r != expected {
					t.Errorf("expected %d but got %d", expected, r)
				}
			})
		}(a, b)
	}

	wg.Wait()
}
