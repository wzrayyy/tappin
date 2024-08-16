package user

type Boost struct {
	ID              string `json:"id"`
	Level           int    `json:"level"`
	LastUpgradeTime int    `json:"lastUpgradeAt"`
}

type Boosts struct {
	BoostMaxTaps           Boost `json:"boostMaxTaps"`
	BoostEarnPerTap        Boost `json:"boostEarnPerTap"`
	BoostFullAvailableTaps Boost `json:"boostFullAvailableTaps"`
}

type Upgrade struct {
	ID            string `json:"id"`
	Level         int    `json:"level"`
	LastUpgradeAt int    `json:"lastUpgradeAt"`
}

type StreakDays struct {
	ID          string `json:"id"`
	CompletedAt string `json:"completedAt"`
	Days        int    `json:"days"`
}

type Tasks struct {
	StreakDays StreakDays `json:"streak_days"`
}

type ClickerUser struct {
	ID               string  `json:"id"`
	TotalCoins       float32 `json:"totalCoins"`
	Balance          float32 `json:"balanceCoins"`
	Level            int     `json:"level"`
	AvailableTaps    int     `json:"availableTaps"`
	Boosts           Boosts  `json:"boosts"`
	Tasks            Tasks   `json:"tasks"`
	ReferralsCount   int     `json:"referralsCount"`
	MaxTaps          int     `json:"maxTaps"`
	EarnPerTap       int     `json:"earnPerTap"`
	PassivePerSecond float32 `json:"earnPassivePerSec"`
	PassivePerHour   float32 `json:"earnPassivePerHour"`
	RecoverPerSecond int     `json:"tapsRecoverPerSec"`
	CreatedAt        string  `json:"createdAt"`
	BalanceTickets   int     `json:"balanceTickets"`
	TotalKeys        int     `json:"totalKeys"`
	BalanceKeys      int     `json:"balanceKeys"`
}

func (u *ClickerUser) Tick() {
	if u.AvailableTaps+u.RecoverPerSecond < u.MaxTaps {
		u.AvailableTaps += u.RecoverPerSecond
	} else {
		u.AvailableTaps = u.MaxTaps
	}

	u.Balance += u.PassivePerSecond
	u.TotalCoins += u.PassivePerSecond
}

type Response struct {
	ClickerUser `json:"clickerUser"`
}
