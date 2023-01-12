package scafreq

import (
  "log"
  "net/http"
  "net/http/cookiejar"
  "os"
  "log"
  "net/url"
  "bytes"
  "scaf/cli/scafio"
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
  if path[len(path)-1] != '/' {
    path += "/"
  }
  log.Printf(os.Getenv("SCAF_BACKEND_URL") + path)
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
  log.Println("DoRequest: ", req.Method, req.URL, req.Body)
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  // show body for debug
  if bodyMap, err := scafio.ReadBody(resp); err == nil {
    log.Println("DoRequest resp body: ")
    for k, v := range bodyMap {
      log.Println("  ", k, ":", v)
    }
  }
  err = SaveCookies(resp.Cookies())
  if err != nil {
    return nil, err
  }

  return resp, nil
}
