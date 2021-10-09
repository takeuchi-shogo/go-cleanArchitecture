package domain

type Tweets struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Contents  string `json:"contents"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updateAt"`
}

type TweetsForGet struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
	Contents string `json:"contents"`
}

type TweetsForPatch struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	Contents string `json:"contents"`
}
