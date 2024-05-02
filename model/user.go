package model

type UserModel struct {
	ID                   string `json:"id"`
	TwitterID            string `json:"twitterId"`
	TwitterHandle        string `json:"twitterHandle"`
	DiscordID            string `json:"discordId"`
	DiscordHandle        string `json:"discordHandle"`
	Name                 string `json:"name"`
	Avatar               string `json:"avatar"`
	ReferralCode         string `json:"referralCode"`
	OnboardInputReferral bool   `json:"onboardInputReferral"`
	EthWalletAddress     string `json:"ethWalletAddress"`
	SeiWalletAddress     string `json:"seiWalletAddress"`
	Email                string `json:"email"`
	Role                 string `json:"role"`
}
