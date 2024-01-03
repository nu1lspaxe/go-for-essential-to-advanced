# VS Code Tips
- `shift + F9`: debug inline

<br>

# Useful third part packages

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
- 