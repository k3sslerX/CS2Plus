package statsFaceit

type FaceitUser struct {
	SteamID64 string `json:"steamid64"`
	CustomURL string `json:"custom_url"`
	FaceitID  string `json:"faceit_id"`
}
