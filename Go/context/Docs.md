# Principle of Go Context

Context 用於簡化單個 Web 請求的多個 goroutine 之間與請求有關的數據、取消信號、截止時間等操作。

對服務器傳入的請求應該創建內容, 而對服務器的傳出調用應該接受內容。他們之間的調用鏈必須傳遞內容，或是使用 WithCancel, WithDeadline, WithTimeout, WithValue 創建。

雖然協程都是併發執行的，但一旦某個協程請求被取消時，所有用來處理該請求的 goroutine 都會迅速退出，系統釋放前者占用的資源。
<br>

### Context 介面定義
```go=
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

- ```Deadline()```
  return context 完成工作的 deadline
- ```Done()```: return 一個 channel, 該 channel 會在工作完成或是內容被讀取完畢後關閉. 多次調用 ```Done()``` 會 return 同一個 channel
- ```Err()```: return context 結束的原因, 只會在 Done() return 的 channel 被關閉才會返回非 nil:
  - 如果 context 被取消, return Canceled
  - 如果 context 超時, return DeadlineExceeded
- ```Value()```: 從 context 中返回鍵對應的值. 同一個內容中, 相同 Key 多次調用 Value 會返回相同結果

<br>

### Context 內建函數
- ```Background()``` and ```TODO()```:兩者都是沒有deadline, 也沒有攜帶任何值的 context
  - ```Background()```: 主要用於 ```main()```, 初始化和測試. 在 context tree 中的最頂層 (即 root context)
  - ```TODO()```: 不知道使用什麼 context 時可以使用

<br>

### Context With Func
當一個 context 被取消時，由其派生的 goroutine 也會被取消。
- WithCancel
  
  ```func WithCancel(parent Context) (ctx Context, cancel CancelFunc)```
- WithDeadline
  
  ```func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)```
- WithTimeout
  
  ```WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)```
- WithValue
  
  ```func WithValue(parent Context, key, val interface{}) Context```
  
  ```context.WithValue```所提供的 key 必須是可比較的，且通常是自定義型別

<br>

### Notes of Context
- 建議以參數方式顯式傳遞(值) context，並應將其作為第一個參數
  > 隱式傳遞是指針
- 如果呼叫一個方法需要傳遞 context 不知道要傳什麼時，不要傳遞 ```nil```，使用 ```context.TODO()```
- context 的 value 相關方法應傳遞請求域的必要數據，不能使用可選參數
- context 是線程安全的，可以在多個 goroutine 中傳遞
  > 線程安全 (Thread Safety) 是指在併發編程中，當有多個線程訪問同個對象時，該對象能在線程間同步數據運行，而不出現數據不一致或意外行為 (不需要額外的同步措施保護，調用結果是一致可預測的)。