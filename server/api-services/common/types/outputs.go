package types

type RESPONSE_PARAMETERS struct {
	IsSuccess bool `json:"isSuccess"`
	Result    any  `json:"result"`
}

type GET_SEARCHED_PRODUCTS_OUTPUT struct {
	SearchProductView
	Document        string `db:"document" json:"-"`
	SsQuery         string `db:"ss_query" json:"-"`
	RankProductName string `db:"rank_product_name" json:"-"`
	Similarity      string `db:"similarity" json:"-"`
}
