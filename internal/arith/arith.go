package arith

import "errors"

// Args struct
type Args struct {
	A, B int
}

// Quotient struct
type Quotient struct {
	Quo, Rem int
}

// Arith type
type Arith int

// Multiply method
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide method
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
