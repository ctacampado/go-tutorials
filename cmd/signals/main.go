package main

import "fmt"

// Flags type
type Flags uint

const (
	// FlagUp indicates if it is up
	FlagUp Flags = 1 << iota
	// FlagBroadcast indicates if it supports broadcast access capability
	FlagBroadcast
	// FlagLoopback indicates if it is a loopback interface
	FlagLoopback
	// FlagPointToPoint indicates if it belongs to a pont-to-point link
	FlagPointToPoint
	// FlagMulticast supports multicast access capability
	FlagMulticast
)

func main() {
	fmt.Printf("%08b\n", FlagUp)
	fmt.Printf("%08b\n", FlagBroadcast)
	fmt.Printf("%08b\n", FlagLoopback)
	fmt.Printf("%08b\n", FlagPointToPoint)
	fmt.Printf("%08b\n", FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

// IsUp checks if Up is set
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

// TurnDown clears the Up flag
func TurnDown(v *Flags) { *v &^= FlagUp }

// SetBroadcast sets the Broadcast flag
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }

// IsCast checks if Multicast and Broadcast flags are set
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }
