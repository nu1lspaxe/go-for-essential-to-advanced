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