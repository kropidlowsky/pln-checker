package attacker

import (
	"fmt"

	"github.com/kropidlowsky/pln-checker/internal/config/options"
	"github.com/kropidlowsky/pln-checker/internal/request"
)

type Attacker struct {
	host      string
	rate      uint
	frequency uint
	results   []request.RequestResult

	// rw sync.RWMutex
}

func NewAttacker(opts options.LoadOpts) *Attacker {
	return &Attacker{
		host:      opts.Host.String(),
		rate:      opts.Rate,
		frequency: opts.Frequency,
	}
}

// Attack performs
func (a *Attacker) Attack() {
	for i := 0; i < int(a.rate); i++ {
		a.attack()
	}
}

// attack performs one attack.
func (a *Attacker) attack() {
	req := request.NewRequest(a.host)
	result, err := req.Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", result)
	a.results = append(a.results, result)
}
