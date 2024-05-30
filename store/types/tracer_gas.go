package types

import (
	"crypto/rand"
	"encoding/hex"
	fmt "fmt"
)

var _ GasMeter = &TracingBasicGasMeter{}

type TracingBasicGasMeter struct {
	tracing bool
	id      string
	basicGasMeter
}

// NewTracingGasMeter returns a reference to a new basicGasMeter.
func NewTracingGasMeter(limit Gas) GasMeter {
	id := make([]byte, 16)
	_, err := rand.Read(id)
	if err != nil {
		panic(err)
	}

	return &TracingBasicGasMeter{
		tracing: false,
		id:      hex.EncodeToString(id),
		basicGasMeter: basicGasMeter{
			limit:    limit,
			consumed: 0,
		},
	}
}

func (g *TracingBasicGasMeter) EnableTracing() {
	g.tracing = true
}

func (g *TracingBasicGasMeter) DisableTracing() {
	g.tracing = false
}

func (g *TracingBasicGasMeter) ConsumeGas(amount Gas, descriptor string) {
	if !g.tracing {
		g.basicGasMeter.ConsumeGas(amount, descriptor)
		return
	}
	before := g.basicGasMeter.consumed
	g.basicGasMeter.ConsumeGas(amount, descriptor)
	after := g.basicGasMeter.consumed
	fmt.Printf("ConsumeGas: %s, %d -> %d, ctx: %s\n", descriptor, before, after, g.id)
}
