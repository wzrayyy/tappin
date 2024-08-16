package clicker

type tapRequest struct {
	Count         int   `json:"count"`
	AvailableTaps int   `json:"availableTaps"`
	Timestamp     int64 `json:"timestamp"`
}

type buyUpgradeRequest struct {
	UpgradeID string `json:"upgradeId"`
	Timestamp int64  `json:"timestamp"`
}

type buyBoostRequest struct {
	BoostID   string `json:"boostId"`
	Timestamp int64  `json:"timestamp"`
}

type checkTaskRequest struct {
	TaskID string `json:"taskId"`
}

type claimDailyCipherRequest struct {
	Cipher string `json:"cipher"`
}
