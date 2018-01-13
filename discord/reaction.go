package discord

type Reaction struct {
	Count uint   `json:"count"`
	Me    bool   `json:"me"`
	Emoji *Emoji `json:"Emoji"`
}
