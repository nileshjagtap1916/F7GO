package main

type contract struct {
	Lookupkeyhash        string  `json:"lookupkeyhash"`
	Agency_name          string  `json:"agency_name"`
	Plan_no              string  `json:"plan_no"`
	Plan_name            string  `json:"plan_name"`
	Io_no                string  `json:"io_no"`
	Media_code           string  `json:"media_code"`
	Media_name           string  `json:"media_name"`
	Menu_code            string  `json:"menu_code"`
	Menu_name            string  `json:"menu_name"`
	Start_date           string  `json:"start_date"`
	End_date             string  `json:"end_date"`
	Order_type           string  `json:"order_type"`
	Budget_amount        float64 `json:"budget_amount"`
	Currency_code        string  `json:"currency_code"`
	Agency_fee_percent   float64 `json:"agency_fee_percent"`
	Platform_name        string  `json:"platform_name"`
	Account_id           string  `json:"account_id"`
	Account_type         string  `json:"account_type"`
	Total_fee_percent    float64 `json:"total_fee_percent"`
	Optimize_fee_percent float64 `json:"optimize_fee_percent"`
}
