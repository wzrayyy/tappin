package boosts

type Item struct {
	ID                   string `json:"id"`
	Price                int    `json:"price"`
	EarnPerTap           int    `json:"earnPerTap"`
	MaxTaps              int    `json:"maxTaps"`
	CooldownSeconds      *int   `json:"cooldownSeconds"`
	TotalCooldownSeconds *int   `json:"totalCooldownSeconds"`
	Level                int    `json:"level"`
	MaxTapsDelta         int    `json:"maxTapsDelta"`
	EarnPerTapDelta      int    `json:"earnPerTapDelta"`
	MaxLevel             *int   `json:"maxLevel,omitempty"`
}

func (i *Item) Tick() {
	if i.TotalCooldownSeconds != nil && *i.TotalCooldownSeconds > 0 {
		(*i.TotalCooldownSeconds)--
	}
}

type Boosts []*Item

type Response struct {
	Boosts `json:"boostsForBuy"`
}

func (b *Boosts) Tick() {
	for _, el := range *b {
		el.Tick()
	}
}

func (b *Boosts) SelectById(id string) *Item {
	for _, i := range *b {
		if i.ID == id {
			return i
		}
	}
	return nil
}
