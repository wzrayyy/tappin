package clicker

import (
	"sync"

	"github.com/wzrayyy/tappin/internal/entity/boosts"
	"github.com/wzrayyy/tappin/internal/entity/config"
	"github.com/wzrayyy/tappin/internal/entity/upgrades"
	"github.com/wzrayyy/tappin/internal/entity/user"
	"golang.org/x/sync/errgroup"
)

func fetchAndUpdate[T any](c *Clicker, endpoint string, lock *sync.RWMutex, setter func(*T)) error {
	var resp T

	err := c.requestAndDecode(endpoint, nil, &resp)
	if err != nil {
		return err
	}

	lock.Lock()
	setter(&resp)
	lock.Unlock()

	return nil
}

func (c *Clicker) Update() error {
	errs := errgroup.Group{}

	errs.Go(func() error {
		return fetchAndUpdate(c, "sync", &c.locks.User, func(r *user.Response) {
			c.user = r
		})
	})

	errs.Go(func() error {
		return fetchAndUpdate(c, "config", &c.locks.Config, func(r *config.Response) {
			c.clickerConfig = r
		})
	})

	errs.Go(func() error {
		return fetchAndUpdate(c, "boosts-for-buy", &c.locks.Boosts, func(r *boosts.Response) {
			c.boosts = r
		})
	})

	errs.Go(func() error {
		return fetchAndUpdate(c, "upgrades-for-buy", &c.locks.Upgrades, func(r *upgrades.Response) {
			c.upgrades = r
		})
	})

	return errs.Wait()
}
