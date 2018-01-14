package discord

type Attachment struct {
	ID       uint64 `json:"id"`
	Filename string `json:"filename"`
	Size     uint   `json:"size"`
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   uint   `json:"height"`
	Width    uint   `json:"width"`
}
