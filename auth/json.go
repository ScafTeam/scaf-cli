package auth

type AuthRequest struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

type ForgetPasswordRequest struct {
  Email string `json:"email"`
}
