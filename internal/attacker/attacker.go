package attacker

import (
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
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

	var counter int64 = -1

	go func() {
		for range ticker.C {
			a.wg.Add(1)
			go func() {
				defer a.wg.Done()
				atomic.AddInt64(&counter, 1)
				a.Attack(int(counter))
			}()
		}
	}()

	a.wg.Wait()
	<-sigChan
}

// Attack performs many attacks at the same time.
func (a *Attacker) Attack(series int) {
	var wg sync.WaitGroup
	for i := 0; i < int(a.rate); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.singleAttack(series, i)
		}()
	}

	wg.Wait()
}

// singleAttack performs one attack.
func (a *Attacker) singleAttack(series, attackNumber int) {
	req := request.NewRequest(a.host)

	result, err := req.Get()
	if err != nil {
		a.logger.Error("failed to visit", zap.Int("series", series), zap.Int("attack", attackNumber))
	}

	a.logger.Info("visited", zap.Int("series", series), zap.Int("attack", attackNumber), zap.Any("result", result))
}
