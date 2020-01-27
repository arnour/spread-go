# spread

A library to calculate hashed-based key distribution

[![Build Status](https://travis-ci.org/arnour/spread.svg?branch=master)](https://travis-ci.org/arnour/spread)

## Installation

    $ go get github.com/arnour/spread-go

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

    $ make test
