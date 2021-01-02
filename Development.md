# Development

## Testing

For now, test coverage should never fall below 50%

```bash
go test ./... -coverprofile="coverage.txt" -covermode=atomic -count=1
go tool cover -html="coverage.txt"
```

## Sanity Tools

### Errors

New errors should be custom. Test files can be ignored.

*Finds all usages of funcs from the `errors` package*
```bash
grep -r --exclude="*_test.go" --exclude="*md" "errors." ./source/hookboy
```

*Finds all usages of error*
```bash
grep -r --exclude="*_test.go" --exclude="*md" "error" ./source/hookboy
```