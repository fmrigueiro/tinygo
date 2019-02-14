// +build !byollvm

package loader

/*
#cgo CFLAGS: -I/usr/lib/llvm-8/include
#cgo LDFLAGS: -L/usr/lib/llvm-8/lib -lclang
*/
import "C"
