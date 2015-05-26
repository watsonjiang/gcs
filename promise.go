// +build ignore

package gcs

/*
   Allows a thread to submit an asynchronous request and to wait for the result.
   The caller may choose to check for the result at a later time, or immediately and
   it may block or not. Both the caller and responder have to know the promise.
   When the result is abailable, 'HasResult' will always return true and 'GetResult'
   will return the result. In order to block for a different result, 'Reset' has to be 
   called first.
*/
type Promise struct {
   lock sync.Mutex
   result interface{}
   hasResult bool
   c  chan interface{}
}

/*
  Blocks until a result is available, or timeout have elapsed.
*/
var ERR_PROMISE_TIMEOUT = errors.New("Timeout")
func (p *Promise) GetResultWithTimeout(timeout time.Duration) (interface{}, error) {
   if p.HasResult() {
      return p.result, nil
   }
   if timeout <= 0 {
      <-p.c
   }else{
      select {
      case <-p.c:
      case time.After(timeout):
          return nil, ERR_PROMISE_TIMEOUT
      }
   }
   return p.result, nil
}

/*
  Blocks until a result is available
*/
func (p *Promise) GetResult() (interface{}, error) {
   res, err := GetResultWithTimeout(0)
   return res, err
}

/*
  Checks whether result is available. Does not block.
*/
func (p *Promise) HasResult() bool {
   p.lock.Lock()
   defer p.lock.Unlock()
   return p.hasResult
}

/*
  Sets the result and notifies any goroutines waiting for it.
*/
func (p *Promise) SetResult(obj interface{}) {
   p.lock.Lock()
   defer p.lock.Unlock()
   p.result = obj
   p.hasResult = true
   p.c <- obj
}

/*
  Causes all waiting goroutines to return
*/
func (p *Promise) Reset() {
   p.lock.Lock()
   defer p.lock.Unlock()
   p.result = nil
   p.hasResult = false
   close(p.c)
   p.c = make(chan interface{})
}
