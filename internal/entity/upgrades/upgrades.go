package upgrades

import "time"

type Item struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Price         int        `json:"price"`
	ProfitPerHour int        `json:"profitPerHour"`
	Condition     *Condition `json:"condition,omitempty"`
	Section       string     `json:"section"`
	Level         int        `json:"level"`
	CurrentProfit int        `json:"currentProfitPerHour"`
	ProfitDelta   int        `json:"profitPerHourDelta"`
	IsAvailable   bool       `json:"isAvailable"`
	IsExpired     bool       `json:"isExpired"`
	Cooldown      *int       `json:"cooldownSeconds"`
	TotalCooldown *int       `json:"totalCooldownSeconds"`
	ReleaseAt     *string    `json:"releaseAt"`
	ExpiresAt     *string    `json:"expiresAt"`
	MaxLevel      *int       `json:"maxLevel"`
}

func (i *Item) Tick() {
	if i.ExpiresAt != nil && !i.IsExpired && *i.ExpiresAt != "" {
		expire_time, _ := time.Parse(time.RFC3339, *i.ExpiresAt)
		if expire_time.Sub(time.Now()) < 0 {
			i.IsExpired = true
		}
	}
	if i.Cooldown != nil && *i.Cooldown > 0 {
		(*i.Cooldown)--
	}
}

type Section struct {
	Name        string `json:"section"`
	IsAvailable bool   `json:"isAvailable"`
}

type DailyCombo struct {
	UpgradeIDs       []string `json:"upgradeIds"`
	BonusCoins       int      `json:"bonusCoins"`
	IsClaimed        bool     `json:"isClaimed"`
	RemainingSeconds int      `json:"remainSeconds"`
}

type Upgrades map[string]*Item

type Response struct {
	UpgradesArray []*Item    `json:"upgradesForBuy"`
	Upgrades      Upgrades   `json:"-"`
	Sections      []Section  `json:"sections"`
	DailyCombo    DailyCombo `json:"dailyCombo"`
}

func (i *Item) PriceByLevel(level int) int {
	return 0
}

func (i *Item) ProfitDeltaByLevel(level int) int {
	return 0
}

func (i *Item) CooldownByLevel(level int) int {
	return 0
}

// def profit_delta_by_level(self, level: int) -> int:
//     return round(self.profit_per_hour_delta * 1.07 ** level)
//
// def price_by_level(self, level: int) -> int:
//     return round(self.price * 1.05 ** ((level + 3) * level / 2))
//
// def cooldown_by_level(self, level: int) -> int:
//     return self.cooldown_seconds * 2 ** level if self.cooldown_seconds else 0
