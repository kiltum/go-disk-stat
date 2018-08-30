package main

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func ToKb(amount uint64) float64 {
	return float64(amount / KB)
}

func ToMb(amount uint64) float64 {
	return float64(amount / MB)
}

func ToGb(amount uint64) float64 {
	return float64(amount / GB)
}
