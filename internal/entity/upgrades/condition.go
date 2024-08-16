package upgrades

import (
	"encoding/json"
)

type Type string

const (
	byUpgrade                Type = "ByUpgrade"
	referralCount            Type = "ReferralCount"
	moreReferralCount        Type = "MoreReferralsCount"
	subscribeTelegramChannel Type = "SubscribeTelegramChannel"
)

type Condition struct {
	Type                     Type
	ByUpgrade                *ByUpgrade
	ReferralCount            *ReferralCount
	MoreReferralCount        *MoreReferralCount
	SubscribeTelegramChannel *SubscribeTelegramChannel
}

type ByUpgrade struct {
	Level     int
	UpgradeID string
}

type ReferralCount struct {
	ReferralCount int
}

type MoreReferralCount struct {
	MoreReferralCount int
}

type SubscribeTelegramChannel struct {
	ChannelID int
	Link      string
}

func (c *Condition) UnmarshalJSON(d []byte) error {
	type tmpStruct struct {
		Type Type `json:"_type"`
		ByUpgrade
		ReferralCount
		MoreReferralCount
		SubscribeTelegramChannel
	}
	var tmp = new(tmpStruct)
	if err := json.Unmarshal(d, tmp); err != nil {
		return err
	}

	c.Type = tmp.Type
	switch tmp.Type {
	case byUpgrade:
		c.ByUpgrade = &tmp.ByUpgrade
	case referralCount:
		c.ReferralCount = &tmp.ReferralCount
	case moreReferralCount:
		c.MoreReferralCount = &tmp.MoreReferralCount
	case subscribeTelegramChannel:
		c.SubscribeTelegramChannel = &tmp.SubscribeTelegramChannel
	}

	return nil
}
