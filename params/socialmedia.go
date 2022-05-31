package params

type CreateSocialMedia struct{
	Name string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

type UpdateSocialMedia struct{
	Name string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}