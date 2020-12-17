# Development

## Testing

For now, test coverage should never fall below 50%

```bash
go test -coverprofile="coverage.out" 
go tool cover -html="coverage.out"
```