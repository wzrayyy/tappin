package upgrades

import (
	"encoding/json"
)

func (u *Upgrades) RecurseUnavailable(item *Item) []*Item {
	if item.IsAvailable || (item.Condition == nil || item.Condition.ByUpgrade == nil) {
		return []*Item{item}
	}

	return append(u.RecurseUnavailable((*u)[item.Condition.ByUpgrade.UpgradeID]), item)
}

func (r *Response) UnmarshalJSON(data []byte) error {
	type tempType Response
	var temp *tempType = (*tempType)(r)

	err := json.Unmarshal(data, temp)
	if err != nil {
		return err
	}

	r.Upgrades = make(Upgrades)

	for _, i := range r.UpgradesArray {
		r.Upgrades[i.ID] = i
	}

	return nil
}

func (r *Response) Tick() {
	for _, e := range r.UpgradesArray {
		e.Tick()
	}
}
