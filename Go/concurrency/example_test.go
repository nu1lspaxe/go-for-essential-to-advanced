package concurrency_test

import (
	"testing"

	"github.com/nu1lspaxe/go-for-essential-to-advanced/Go/concurrency"
)

func TestM1(t *testing.T) {
	processor := concurrency.DefaultProcessor{}
	m1 := concurrency.M1{Processor: processor}
	m1.Run()
}
