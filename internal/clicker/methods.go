package clicker

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/wzrayyy/tappin/internal/entity/user"
)

func (c *Clicker) ClaimDailyKeys() error {
	m, err := json.Marshal(&claimDailyCipherRequest{
		Cipher: base64.StdEncoding.EncodeToString([]byte(strings.Repeat("0", 10) + fmt.Sprintf("|%d", c.telegramUserID))),
	})

	if err != nil {
		return err
	}

	err = c.requestAndDecode("start-keys-minigame", nil, nil)

	if err != nil {
		return err
	}

	return c.requestAndDecode("claim-daily-keys-minigame", m, nil)
}

func (c *Clicker) ClaimDailyCipher() error {
	cipher, err := c.clickerConfig.DailyCipher.Decode()
	if err != nil {
		return err
	}

	m, err := json.Marshal(claimDailyCipherRequest{
		Cipher: cipher,
	})

	return c.requestAndDecode("claim-daily-cipher", m, nil)
}

func (c *Clicker) Tap(taps int) error {
	c.locks.User.RLock()
	m, err := json.Marshal(tapRequest{
		Count:         taps,
		AvailableTaps: c.user.AvailableTaps - taps,
		Timestamp:     time.Now().UnixMilli(),
	})
	c.locks.User.RUnlock()

	if err != nil {
		return err
	}

	var u user.Response

	err = c.requestAndDecode("tap", m, &u)
	if err != nil {
		return err
	}

	c.locks.User.Lock()
	c.user = &u
	c.locks.User.Unlock()

	return nil
}

func (c *Clicker) BuyBoost(boost_id string) error {
	r, err := json.Marshal(buyBoostRequest{
		BoostID:   boost_id,
		Timestamp: time.Now().UnixMilli(),
	})
	if err != nil {
		return err
	}
	return c.requestAndDecode("buy-boost", r, nil)
}
