# Package Caching

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/caching)](https://goreportcard.com/report/github.com/yourusername/caching)

This Go package provides a simple in-memory caching mechanism with basic operations like Set, Get, and Delete.

## Installation

```bash
go get -u github.com/fentezi/cache
```
## Creating a new cache instance
```
cache := caching.New()
```
## Setting a value in the cache
```
cache.Set("key", "value")
```
## Setting a value in the cache
```
cache.Set("key", "value")
```
## Getting a value from the cache
```
value := cache.Get("key")
fmt.Println(value)
```
## Deleting a value from the cache
```
cache.Delete("key")
```

## Example
```
package main

import (
	"fmt"
	"github.com/fentezi/caching"
)

func main() {
	// Create a new cache instance
	cache := caching.New()

	// Set a value in the cache
	cache.Set("username", "john_doe")

	// Get the value from the cache
	username := cache.Get("username")
	fmt.Println("Username:", username)

	// Delete the value from the cache
	cache.Delete("username")
}
```
