package domain

type Diaries struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type DiariesForPatch struct {
	ID      int
	UserID  int
	Title   string
	Content string
}
