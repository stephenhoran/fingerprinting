# Fingerprinting

A comparison of Go's implementation of FNV-1a 64-bit vs Prometheus's "Fast" implementation (no collision checking)

`
pkg: fingerprinting
BenchmarkNewGoHash
BenchmarkNewGoHash-8     	 4800076	       248 ns/op
BenchmarkNewPromHash
BenchmarkNewPromHash-8   	 4784163	       244 ns/op
`
