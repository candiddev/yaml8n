package main

import (
	"context"
	"os"
	"testing"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	r := m.Run()

	os.Exit(r)
}
