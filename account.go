package main

type Account struct {
	AccountName       string `json:"account_name"`
	HeadBlockNum      int    `json:"head_block_num"`
	HeadBlockTime     string `json:"head_block_time"`
	Privileged        bool   `json:"privileged"`
	LastCodeUpdate    string `json:"last_code_update"`
	Created           string `json:"created"`
	CoreLiquidBalance string `json:"core_liquid_balance"`
	RAMQuota          int64    `json:"ram_quota"`
	NetWeight         int    `json:"net_weight"`
	CPUWeight         string `json:"cpu_weight"`
	NetLimit          struct {
		Used      int `json:"used"`
		Available int `json:"available"`
		Max       int `json:"max"`
	} `json:"net_limit"`
	CPULimit struct {
		Used      int `json:"used"`
		Available int `json:"available"`
		Max       int `json:"max"`
	} `json:"cpu_limit"`
	RAMUsage    int64 `json:"ram_usage"`
	Permissions []struct {
		PermName     string `json:"perm_name"`
		Parent       string `json:"parent"`
		RequiredAuth struct {
			Threshold int           `json:"threshold"`
			Keys      []interface{} `json:"keys"`
			Accounts  []struct {
				Permission struct {
					Actor      string `json:"actor"`
					Permission string `json:"permission"`
				} `json:"permission"`
				Weight int `json:"weight"`
			} `json:"accounts"`
			Waits []interface{} `json:"waits"`
		} `json:"required_auth"`
	} `json:"permissions"`
	TotalResources struct {
		Owner     string `json:"owner"`
		NetWeight string `json:"net_weight"`
		CPUWeight string `json:"cpu_weight"`
		RAMBytes  int    `json:"ram_bytes"`
	} `json:"total_resources"`
	SelfDelegatedBandwidth struct {
		From      string `json:"from"`
		To        string `json:"to"`
		NetWeight string `json:"net_weight"`
		CPUWeight string `json:"cpu_weight"`
	} `json:"self_delegated_bandwidth"`
	RefundRequest struct {
		Owner       string `json:"owner"`
		RequestTime string `json:"request_time"`
		NetAmount   string `json:"net_amount"`
		CPUAmount   string `json:"cpu_amount"`
	} `json:"refund_request"`
	VoterInfo struct {
		Owner             string        `json:"owner"`
		Proxy             string        `json:"proxy"`
		Producers         []interface{} `json:"producers"`
		Staked            int           `json:"staked"`
		LastVoteWeight    string        `json:"last_vote_weight"`
		ProxiedVoteWeight string        `json:"proxied_vote_weight"`
		IsProxy           int           `json:"is_proxy"`
		Flags1            int           `json:"flags1"`
		Reserved2         int           `json:"reserved2"`
		Reserved3         string        `json:"reserved3"`
	} `json:"voter_info"`
	RexInfo interface{} `json:"rex_info"`
}
