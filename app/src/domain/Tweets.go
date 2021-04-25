package domain

type Tweets struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	Contents string `json:"contents"`
}

type TweetsForPatch struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	Contents string `json:"contents"`
}
