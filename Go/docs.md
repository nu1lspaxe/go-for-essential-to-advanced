# VS Code Tips
- `shift + F9`: debug inline

<br>

# Quick Guide on logging and Error Handling

#### Logging or others
- `log.Print`: General purpose logging of information. Doesn't terminate.
- `log.Printf`: Same as Print but allows formatting.
- `log.Fatal`: Logs and then calls os.Exit(1) to terminate. For unrecoverable errors.
- `log.Panic`: Logs and panics. For critical unrecoverable errors.
- `log.Panicf`: Same as Panic but allows formatting.
- `fmt.Errorf`: Returns an error object. Doesn't log or terminate. Used to return errors.
- `errors.New`: Same as Errorf but without formatting.


#### For testing
- `t.Error`: Logs an error but continues test execution.
- `t.Errorf`: Logs a formatted error but continues.
- `t.Fatal`: Logs and marks the test as failed. Terminates the test.
- `t.Fatalf`: Same as Fatal but with formatting.

#### Summary
- **Print | Printf** for general logging
- **Fatal** for unrecoverable errors that need to terminate
- **Panic** for critical errors that shouldn't continue
- **Errorf** to create error objects to return
- **t.Error | t.Errorf** to log errors but continue tests
- **t.Fatal | t.Fatalf** to log and fail fast during tests

<br>

# Flags
- **escape analysis**
```bash=
go build -gcflags=-m main.go
```
- Run benchmark testing
```bash=
go test -bench .
```
- Detect data race
```bash=
go run -race .
go test -race ./...
```
- `go:linkname`, only useable in condition of package `unsafe`
```go=
//go:linkname localname [importpath.name]
```

<br>

# Useful third-party packages

### strings

- **NewReplacer**
    
    used to replace multiple values, here is an example:
    ```go
    // Even index for argument be replaced, odd index for what you want to replace with.
    replacer := strings.NewReplacer(":", "", "^", "", "*", "")
	str := "Hi:, I'm *Tinaaa, Help.... m^e..:)"
	str = replacer.Replace(str)
    
	fmt.Println(str)    // Hi, I'm Tinaaa, Help.... me..)
    ```