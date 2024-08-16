package clicker

import (
	"fmt"
	"time"
)

func (c *Clicker) genericWorker(fn func() error, period *int, done EmptyChannel) error {
	return c.genericWorkerWithChannel(func(EmptyChannel) error { return fn() }, period, done)
}

func (c *Clicker) genericWorkerWithChannel(fn func(EmptyChannel) error, period *int, done EmptyChannel) error {
	for {
		select {
		case <-done:
			return nil
		case <-c.channels.Global:
			fmt.Println("got global stop")
			if done != nil {
				close(done)
			}
			return nil
		case <-time.After(time.Second * time.Duration(*period)):
			err := fn(done)
			if err != nil {
				close(c.channels.Global)
				return err
			}
		}
	}
}

func (c *Clicker) tapWorker() error {
	return c.genericWorkerWithChannel(func(done EmptyChannel) error {
		c.locks.User.RLock()
		taps_consumed := int(c.Config.TapInterval) * c.user.EarnPerTap
		taps_left := c.user.AvailableTaps - taps_consumed
		c.locks.User.RUnlock()

		fmt.Println("tap")

		if taps_left < 0 {
			c.locks.Boosts.RLock()
			b := c.boosts.SelectById("BoostFullAvailableTaps").CooldownSeconds
			c.locks.Boosts.RUnlock()

			if b != nil && *b <= 0 {
				fmt.Println("buy taps")
				c.BuyBoost("BoostFullAvailableTaps")
			} else {
				c.locks.User.RLock()
				time_sleep := (c.user.MaxTaps - c.user.AvailableTaps) / c.user.RecoverPerSecond
				c.locks.User.RUnlock()
				select {
				case <-c.channels.Global:
					return nil
				case <-done:
					return nil
				case <-time.After(time.Duration(time_sleep) * time.Second):
					break
				}
			}
		}
		return c.Tap(c.Config.TapInterval * c.Config.TapsPerSecond)
	}, &c.Config.TapInterval, c.channels.Tap)
}

func (c *Clicker) updateWorker() error {
	return c.genericWorker(func() error {
		return c.Update()
	}, &c.Config.UpdateFrequency, c.channels.Update)
}

func (c *Clicker) tickWorker() error {
	interval := 1
	return c.genericWorker(func() error {
		c.Tick()
		return nil
	}, &interval, nil)
}

func (c *Clicker) Start() error {
	c.errorGroup.Go(c.tapWorker)
	c.errorGroup.Go(c.updateWorker)
	c.errorGroup.Go(c.tickWorker)

	return c.errorGroup.Wait()
}

func (c *Clicker) Stop() error {
	close(c.channels.Global)
	return c.errorGroup.Wait()
}
