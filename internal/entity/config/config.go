package config

type Task struct {
	ID        string  `json:"id"`
	Link      *string `json:"link"`
	Reward    int     `json:"rewardCoins"`
	Cycle     Cycle   `json:"periodicity"`
	ChannelId *int    `json:"channelId"`
}

type ClickerConfig struct {
	MaxPassive int    `json:"maxPassiveDtSeconds"`
	Tasks      []Task `json:"tasks"`
}

type DailyCipher struct {
	Cipher        string `json:"cipher"`
	BonusCoins    int    `json:"bonusCoins"`
	IsClaimed     bool   `json:"isClaimed"`
	RemainSeconds int    `json:"remainSeconds"`
}

type DailyKeys struct {
	StartDate     string  `json:"startDate"`
	LevelConfig   string  `json:"levelConfig"`
	BonusKeys     int     `json:"bonusKeys"`
	IsClaimed     bool    `json:"isClaimed"`
	SecondsToNext int     `json:"totalSecondsToNextAttempt"`
	RemainToGuess float32 `json:"remainSecondsToGuess"`
	RemainSeconds float32 `json:"remainSeconds"`
	RemainToNext  float32 `json:"remainSecondsToNextAttempt"`
}

type Response struct {
	ClickerConfig ClickerConfig `json:"clickerConfig"`
	DailyCipher   DailyCipher   `json:"dailyCipher"`
	DailyKeys     DailyKeys     `json:"DailyKeysMiniGame"`
}

func (c *DailyCipher) Tick() {
	c.RemainSeconds--
}

func (k *DailyKeys) Tick() {
	k.RemainToGuess--
	k.RemainToNext--
	k.RemainSeconds--
	k.SecondsToNext--
}

func (r *Response) Tick() {
	r.DailyCipher.Tick()
	r.DailyKeys.Tick()
}
