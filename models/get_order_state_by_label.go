package models

type GetOrderStateByLabelParams struct {
	Currency string `json:"currency"`
	Label    string `json:"label"`
}
