package auth

type SignInRequest struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

type ForgetPasswordRequest struct {
  Email string `json:"email"`
}
