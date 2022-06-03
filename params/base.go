package params

type Response struct {
	Status              int         `json:"status"`
	Message             string      `json:"message,omitempty"`
	Error               string      `json:"error,omitempty"`
	AdditionalInfo      interface{} `json:"additional_info,omitempty"`
	Data                interface{} `json:"data,omitempty"`
	UserResponse        interface{} `json:"user_data,omitempty"`
	SocialMediaResponse interface{} `json:"social_media_data,omitempty"`
}

type UserResponse struct {
	ID        int         `json:"id,omitempty"`
	Email     string      `json:"email,omitempty"`
	Username  string      `json:"username,omitempty"`
	Age       int         `json:"age,omitempty"`
	Token     string      `json:"token,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
}

type SocialMediaResponse struct {
	ID             uint          `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	SocialMediaURL string        `json:"social_media_url,omitempty"`
	UserID         uint          `json:"user_id,omitempty"`
	CreatedAt      interface{}   `json:"created_at,omitempty"`
	UpdatedAt      interface{}   `json:"updated_at,omitempty"`
	User           UserResponse `json:"user,omitempty"`
}
