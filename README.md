# spread

A library to calculate key-based hash distribution


## Example Usage

```go
// Create a spread instance
spreader := spread.New(nil)

// you could also use different hash implamentations
// spreader := spread.New(sha256.New())

keyValue := "my-key-to-hash"

fraction := spreader.Key(keyValue)

fmt.Println(fraction) // 0.804535691348
```

### Testing

``make test``