# Principle of Go Context
```go=
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

- Deadline(): return context 完成工作的 deadline
- Done(): return 一個 channel, 該 channel 會在工作完成或是上下文被讀取完畢後關閉. 多次調用 Done() 會 return 同一個 channel
- Err(): return context 結束的原因, 只會在 Done() return 的 channel 被關閉才會返回非 nil:
  - 如果 context 被取消, return Canceled
  - 如果 context 超時, return DeadlineExceeded
- Value(): 從 context 中返回鍵對應的值. 同一個上下文中, 相同 Key 多次調用 Value 會返回相同結果

- Background() and TODO():兩者都是沒有deadline, 也沒有攜帶任何值的 context
  - Background(): 主要用於 main(), 初始化和測試. 在 context tree 中的最頂層 (即 root context)
  - TODO(): 不知道使用什麼 context 時可以使用
