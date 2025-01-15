package attacker

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kropidlowsky/pln-checker/internal/config/options"
	"github.com/kropidlowsky/pln-checker/internal/request"
	"go.uber.org/zap"
)

type Attacker struct {
	host      string
	rate      uint
	frequency uint
	logger    *zap.Logger

	wg sync.WaitGroup
}

func NewAttacker(opts options.LoadOpts, logger *zap.Logger) *Attacker {
	return &Attacker{
		host:      opts.Host.String(),
		rate:      opts.Rate,
		frequency: opts.Frequency,
		logger:    logger,
	}
}

// InfiniteAttack frequently performs many attacks at the same time for infinite time - it can be stopped by abortting the program (`ctrl` + `c`).
func (a *Attacker) InfiniteAttack() {
	ticker := time.NewTicker(time.Duration(a.frequency) * time.Second)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for range ticker.C {
			a.Attack()
		}
	}()

	<-sigChan
}

// Attack performs many attacks at the same time.
func (a *Attacker) Attack() {
	for i := 0; i < int(a.rate); i++ {
		a.wg.Add(1)
		go func() {
			defer a.wg.Done()
			a.singleAttack()
		}()

	}
}

// singleAttack performs one attack.
func (a *Attacker) singleAttack() {
	req := request.NewRequest(a.host)

	result, err := req.Get()
	if err != nil {
		panic(err)
	}

	a.logger.Info("visited", zap.Any("result", result))
}
