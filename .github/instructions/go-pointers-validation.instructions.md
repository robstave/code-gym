---
description: "Use when changing Go code in go/lgwt/pointers; validate with recursive tests and errcheck."
applyTo: "go/lgwt/pointers/**,go/lgwt/pointers/**"
---

# Go validation

For Go code changes, run `go test ./...` and `errcheck ./...` from the directory where the work is being done, within the Go module. These commands cover the current directory and its subdirectories.

Report failures or missing tools explicitly; do not silently skip either check.
