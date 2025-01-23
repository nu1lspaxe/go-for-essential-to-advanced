// Package syncmap provides a concurrent map implementation.
// This was the prototype for sync.Map which was added to
// the standard library's sync package in Go 1.9.
package syncmap

import "sync"

// Map is a concurrent map with amortized-constant-time
// loads, stores, and deletes.
// It is safe for multiple goroutines to call a Map's methods
// concurrently.
//
// The zero Map is valid and empty.
//
// A Map must not be copied after first use.
type Map = sync.Map
