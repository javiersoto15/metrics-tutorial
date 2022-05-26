package kibana

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.elastic.co/apm/v2"
)

func ProcessExample() {
	transactionName := "PROCESS /test"
	transactionType := "message"
	tx := apm.DefaultTracer().StartTransaction(transactionName, transactionType)
	ctx := apm.ContextWithTransaction(context.TODO(), tx)
	defer tx.End()
	err := child1(ctx)
	if err != nil {
		apm.CaptureError(ctx,err)
	}
	err = child1(ctx)
	if err != nil {
		apm.CaptureError(ctx,err)
	}
}

func ReadExample() {
	transactionName := "READ /test"
	transactionType := "message"
	tx := apm.DefaultTracer().StartTransaction(transactionName, transactionType)
	ctx := apm.ContextWithTransaction(context.TODO(), tx)
	defer tx.End()
	err := child2(ctx)
	if err != nil {
		apm.CaptureError(ctx,err)
	}
	err = child1(ctx)
	if err != nil {
		apm.CaptureError(ctx,err)
	}
	err = child3(ctx)
	if err != nil {
		apm.CaptureError(ctx,err)
	}
	err = child2(ctx)
	if err != nil {
		apm.CaptureError(ctx,err)
	}
}

func GetHandlerExample(ctx context.Context) {

}

func child1(ctx context.Context) error {
	operationType := "test.process"
	span, spanCtx := apm.StartSpan(ctx, "SUM child1", operationType)
	span.Action = "addition"
	defer span.End()
	nbr1 := generateRandomIntWithSpan(ctx, 0, 100)
	nbr2 := generateRandomIntWithSpan(ctx, 0, 100)
	sum := nbr1 + nbr2
	if sum > 190 {
		return fmt.Errorf("intended error child1")
	}
	err := child3(spanCtx)
	if err != nil {
		err = fmt.Errorf("error child1 from child3: %w", err)
	}
	return err
}

func child2(ctx context.Context) error {
	operationType := "test.process"
	span, _ := apm.StartSpan(ctx, "DIFF child2", operationType)
	span.Action = "substraction"
	defer span.End()
	nbr1 := generateRandomIntWithSpan(ctx, 50, 100)
	nbr2 := generateRandomIntWithSpan(ctx, 0, 50)
	diff := nbr1 - nbr2
	if diff < 2 {
		return fmt.Errorf("intended error child2")
	}
	return nil
}

func child3(ctx context.Context) error {
	operationType := "test.process"
	span, _ := apm.StartSpan(ctx, "MOD child3", operationType)
	span.Action = "module"
	defer span.End()
	nbr1 := generateRandomIntWithSpan(ctx, 1, 100)
	nbr2 := generateRandomIntWithSpan(ctx, 1, 10)
	_ = nbr1 % nbr2
	return nil
}

func generateRandomIntWithSpan(ctx context.Context, min, max int) int {
	operationType := "number.random"
	span, _ := apm.StartSpan(ctx, "RAND generateRandomIntWithSpan", operationType)
	span.Action = "generate"
	defer span.End()
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
