package clicker

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"

	"github.com/wzrayyy/tappin/internal/entity/boosts"
	"github.com/wzrayyy/tappin/internal/entity/config"
	"github.com/wzrayyy/tappin/internal/entity/upgrades"
	"github.com/wzrayyy/tappin/internal/entity/user"
	"golang.org/x/sync/errgroup"
)

type Clicker struct {
	client         *http.Client
	authKey        string
	baseUrl        *url.URL
	clickerConfig  *config.Response
	user           *user.Response
	boosts         *boosts.Response
	upgrades       *upgrades.Response
	telegramUserID int
	errorGroup     errgroup.Group

	locks struct {
		User     sync.RWMutex
		Boosts   sync.RWMutex
		Config   sync.RWMutex
		Upgrades sync.RWMutex
	}

	channels struct {
		Update  EmptyChannel
		Tap     EmptyChannel
		Upgrade EmptyChannel
		Global  EmptyChannel
	}

	Config Config
}

type Config struct {
	UpdateFrequency int
	TapsPerSecond   int
	TapInterval     int
}

func NewClicker(auth_key string, user_id int, config Config) (*Clicker, error) {
	c := new(Clicker)

	var err error

	c.client = new(http.Client)
	c.client.Jar = new(cookiejar.Jar)
	c.authKey = auth_key
	c.telegramUserID = user_id

	c.baseUrl, err = url.Parse(apiEndpoint)
	if err != nil {
		return c, err
	}

	c.Config = config

	return c, c.Update()
}

func (c *Clicker) Tick() {
	c.locks.Config.Lock()
	c.clickerConfig.Tick()
	c.locks.Config.Unlock()

	c.locks.User.Lock()
	c.user.Tick()
	c.locks.User.Unlock()

	c.locks.Boosts.Lock()
	c.boosts.Tick()
	c.locks.Boosts.Unlock()

	c.locks.Upgrades.Lock()
	c.upgrades.Tick()
	c.locks.Upgrades.Unlock()
}
