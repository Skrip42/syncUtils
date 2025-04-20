# syncutils

A Go package providing synchronization utilities.

## Cond

`Cond` is a condition variable implementation that provides a channel-based alternative to `sync.Cond`. It allows goroutines to wait for and signal events using channels.

### Usage

Create a new condition variable using `NewCond()`:

### Methods

#### Wait
```go
func (c *Cond) Wait() <-chan struct{}
```
Returns a channel that will be closed when the condition is signaled. Multiple goroutines can wait on the same condition variable. The channel is initially blocked until either `Signal()` or `Broadcast()` is called.

#### Signal
```go
func (c *Cond) Signal()
```
Unblocks the most recently waiting goroutine (LIFO order). If no goroutines are waiting, the call has no effect.

#### Broadcast
```go
func (c *Cond) Broadcast()
```
Unblocks all waiting goroutines. If no goroutines are waiting, the call has no effect.

### Example

```go
cond := NewCond()

// In goroutine 1
ch := cond.Wait()
<-ch // Blocks until signaled

// In goroutine 2
cond.Signal() // Unblocks the waiting goroutine

// Or to unblock all waiting goroutines
cond.Broadcast()
```

The `Cond` type is thread-safe and can be used concurrently from multiple goroutines.