package types

// TransferActionData data for transfer action
type TransferActionData struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity Asset  `json:"quantity"`
	Memo     string `json:"memo"`
}
