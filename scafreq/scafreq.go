package scafreq

import (
  "net/http"
  "net/http/cookiejar"
  "os"
  "log"
  "net/url"
  "bytes"
)

// GetClient returns a client with cookiejar, and JWT if exists
func GetClient() (*http.Client, error) {
  jar, err := cookiejar.New(nil)
  if err != nil {
    return nil, err
  }
  client := &http.Client{
    Jar: jar,
  }
  jwtCookie, err := LoadCookie("jwt")
  if err != nil {
    return client, nil
  }
  backend_url, err := url.Parse(os.Getenv("SCAF_BACKEND_URL"))
  if err != nil {
    return nil, err
  }
  client.Jar.SetCookies(
    backend_url,
    []*http.Cookie{
      jwtCookie,
    },
  )

  return client, nil
}

func NewRequest(method string, path string, body []byte) (*http.Request, error) {
  req, err := http.NewRequest(
    method,
    os.Getenv("SCAF_BACKEND_URL") + path,
    bytes.NewBuffer(body),
  )
  if err != nil {
    return nil, err
  }

  return req, nil
}

// DoRequest makes a request with JWT, and save the cookies
// have no responsibility to close the resp.body
func DoRequest(req *http.Request) (*http.Response, error) {
  client, err := GetClient()
  if err != nil {
    return nil, err
  }
  log.Println("DoRequest: ", req.Method, req.URL)
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  err = SaveCookies(resp.Cookies())
  if err != nil {
    return nil, err
  }

  return resp, nil
}
