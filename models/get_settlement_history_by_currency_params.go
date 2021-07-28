package models

type GetSettlementHistoryByCurrencyParams struct {
	Currency             string `json:"currency"`
	Type                 string `json:"type,omitempty"`
	Count                int    `json:"count,omitempty"`
	Continuation         string `json:"continuation"`
	SearchStartTimestamp uint64 `json:"search_start_timestamp"`
}
