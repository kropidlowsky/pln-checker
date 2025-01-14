package attacker

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kropidlowsky/pln-checker/internal/config/options"
	"github.com/kropidlowsky/pln-checker/internal/request"
)

type Attacker struct {
	host      string
	rate      uint
	frequency uint
	results   []request.RequestResult

	wg sync.WaitGroup
	rw sync.RWMutex
}

func NewAttacker(opts options.LoadOpts) *Attacker {
	return &Attacker{
		host:      opts.Host.String(),
		rate:      opts.Rate,
		frequency: opts.Frequency,
	}
}

// InfiniteAttack frequently performs many attacks at the same time for infinite time - it can be stopped by abortting the program (`ctrl` + `c`).
func (a *Attacker) InfiniteAttack() {
	ticker := time.NewTicker(time.Duration(a.frequency) * time.Second)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			a.Attack()
		}
	}()

	<-sigChan
}

// Attack performs many attacks at the same time.
func (a *Attacker) Attack() {
	for i := 0; i < int(a.rate); i++ {
		a.wg.Add(1)
		a.singleAttack()
	}
}

// singleAttack performs one attack.
func (a *Attacker) singleAttack() {
	defer a.wg.Done()

	req := request.NewRequest(a.host)

	result, err := req.Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", result)
	a.rw.Lock()
	defer a.rw.Unlock()
	a.results = append(a.results, result)
}
