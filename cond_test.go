package syncutils

import (
	"testing"
	"time"
)

func TestCond_Wait(t *testing.T) {
	cond := NewCond()
	ch := cond.Wait()

	// Channel should be blocked initially
	select {
	case <-ch:
		t.Error("Channel should be blocked")
	case <-time.After(100 * time.Millisecond):
		// Expected behavior
	}
}

func TestCond_Signal(t *testing.T) {
	cond := NewCond()
	ch := cond.Wait()

	// Signal should unblock the waiting goroutine
	go cond.Signal()

	select {
	case <-ch:
		// Expected behavior
	case <-time.After(100 * time.Millisecond):
		t.Error("Channel should be unblocked after Signal")
	}
}

func TestCond_Broadcast(t *testing.T) {
	cond := NewCond()
	channels := make([]<-chan struct{}, 3)

	// Create multiple waiting goroutines
	for i := range channels {
		channels[i] = cond.Wait()
	}

	// Broadcast should unblock all waiting goroutines
	go cond.Broadcast()

	// Check all channels are unblocked
	for i, ch := range channels {
		select {
		case <-ch:
			// Expected behavior
		case <-time.After(100 * time.Millisecond):
			t.Errorf("Channel %d should be unblocked after Broadcast", i)
		}
	}
}

func TestCond_SignalOrder(t *testing.T) {
	cond := NewCond()
	ch1 := cond.Wait()
	ch2 := cond.Wait()

	// Signal should unblock the last waiting goroutine
	go cond.Signal()

	select {
	case <-ch2:
		// Expected behavior
	case <-time.After(100 * time.Millisecond):
		t.Error("Second channel should be unblocked after Signal")
	}

	// First channel should still be blocked
	select {
	case <-ch1:
		t.Error("First channel should still be blocked")
	case <-time.After(100 * time.Millisecond):
		// Expected behavior
	}
}

func TestCond_EmptySignal(t *testing.T) {
	cond := NewCond()
	// Signal on empty cond should not panic
	cond.Signal()
}

func TestCond_EmptyBroadcast(t *testing.T) {
	cond := NewCond()
	// Broadcast on empty cond should not panic
	cond.Broadcast()
}
