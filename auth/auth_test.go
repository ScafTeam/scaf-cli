package auth

import (
  "testing"
  "os"
)

func TestSignIn(t *testing.T) {
  // set backend url
  os.Setenv("SCAF_BACKEND_URL", "http://localhost:8000")

  // correct email and password
  resp, err := signIn("test0@test.com", "123456")
  if err != nil {
    t.Error(err)
  }
  if resp.StatusCode != 200 {
    t.Error("status code should be 200")
  }

  // correct email and wrong password
  resp, err = signIn("test0@test.com", "1234567")
  if err != nil {
    t.Error(err)
  }
  if resp.StatusCode != 401 {
    t.Error("status code should be 401")
  }

  // wrong email and correct password
  resp, err = signIn("ttt", "123456")
  if err != nil {
    t.Error(err)
  }
  if resp.StatusCode != 401 {
    t.Error("status code should be 401")
  }
}
