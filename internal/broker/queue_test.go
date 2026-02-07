package broker_test

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		maxMessages int64
		want        *broker.Queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := broker.NewQueue(tt.maxMessages)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
