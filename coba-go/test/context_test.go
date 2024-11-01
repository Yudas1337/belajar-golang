package test

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")

	ctxF := context.WithValue(ctxC, "f", "F")
	ctxG := context.WithValue(ctxF, "g", "G")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)

	fmt.Println(ctxF.Value("f"))
	fmt.Println(ctxF.Value("c"))
	fmt.Println(ctxF.Value("b"))

	fmt.Println(ctxG.Value("b"))
}
