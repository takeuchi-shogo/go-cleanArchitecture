package domain

type Users struct {
	ID                int     `json:"id"`
	UserName          string  `json:"userName"`
	Sex               string  `json:"sex"`
	Age               int     `json:"age"`
	AddressPrefecture *string `json:"addressPrefecture"`
	AddressCity       *string `json:"addressCity"`
	Tel               *string `json:"tel"`
	Email             *string `json:"email"`
	Password          string  `json:"password"`
}

type UserForMe struct {
	ID                int     `json:"id"`
	UserName          string  `json:"userName"`
	Sex               string  `json:"sex"`
	Age               int     `json:"age"`
	AddressPrefecture *string `json:"addressPrefecture"`
	AddressCity       *string `json:"addressCity"`
	Tel               *string `json:"tel"`
	Email             *string `json:"email"`
	Password          string  `json:"password"`
}

type UserForPatch struct {
	ID                int     `json:"id"`
	UserName          string  `json:"userName"`
	Sex               string  `json:"sex"`
	Age               int     `json:"age"`
	AddressPrefecture *string `json:"addressPrefecture"`
	AddressCity       *string `json:"addressCity"`
	Tel               *string `json:"tel"`
	Email             *string `json:"email"`
	Password          string  `json:"password"`
}
