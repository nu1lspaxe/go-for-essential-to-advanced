# Memory Management

- [A Guide to the Go Garbage Collector](https://tip.golang.org/doc/gc-guide)

## pprof
```go
import _ "net/http/pprof"
```

```bash
go tool pprof http://localhost:6060/debug/pprof/heap

# Install graphviz from https://graphviz.org/download/
# command: top, list [function name], web...
```

## GC logs
```bash
GODEBUG=gctrace=1 ./<filename>.go
```

## Optimization
1. 減少短命物件
   - 減少臨時變數，將重複分配物件的重用
   - 使用 `sync.Pool` 管理頻繁分配和釋放的物件
2. 優化大內存結構
   - 縮小結構體並按需分配，避免一次分配過多內存
   - 使用切片代替數組，減少內存拷貝
3. 控制 Goroutine 數量
   - 避免過多 goroutine 導致頻繁的內存分配和回收
4. 優化 GC 觸發頻率 (垃圾回收百分比)
   ```bash
   # 預設 GOGC=100
   GOGC=50 ./<filename>.go
   ```
5. 手動釋放未使用的內存
   - 調用 `runtime.GC()` 或 `debug.FreeOSMemory()`，但僅在確定程序負載下降時使用