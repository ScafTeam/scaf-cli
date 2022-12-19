package auth

import (
  "os"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "errors"
)

func readJWT() (string, error) {
  cookies, err := readCookies()
  if err != nil {
    return "", err
  }

  for _, cookie := range cookies {
    if cookie.Name == "jwt" {
      return cookie.Value, nil
    }
  }

  return "", errors.New("jwt not found")
}

func readCookies() ([]*http.Cookie, error) {
  if _, err := os.Stat(os.Getenv("SCAF_CONFIG_DIR") + "/cookies.json"); err != nil {
    os.Create(os.Getenv("SCAF_CONFIG_DIR") + "/cookies.json")
    return []*http.Cookie{}, nil
  }
  file, err := os.Open(os.Getenv("SCAF_CONFIG_DIR") + "/cookies.json")
  if err != nil {
    return nil, err
  }
  defer file.Close()

  data, err := ioutil.ReadAll(file)
  if err != nil {
    return nil, err
  }

  var cookies []*http.Cookie
  err = json.Unmarshal(data, &cookies)
  if err != nil {
    os.Remove(os.Getenv("SCAF_CONFIG_DIR") + "/cookies.json")
    return []*http.Cookie{}, nil
  }

  return cookies, nil
}

func saveCookies(resp *http.Response) error {
  cookies, err := readCookies()
  if err != nil {
    return err
  }

  for _, cookie := range resp.Cookies() {
    var found bool = false
    for i, c := range cookies {
      if c.Name == cookie.Name {
        cookies[i] = cookie
        found = true
        break
      }
    }
    if !found {
      cookies = append(cookies, cookie)
    }
  }

  data, err := json.Marshal(cookies)
  if err != nil {
    return err
  }

  err = ioutil.WriteFile(os.Getenv("SCAF_CONFIG_DIR") + "/cookies.json", data, 0777)
  if err != nil {
    return err
  }

  return nil
}

// TODO: refactor deleteCookies
func deleteCookies() error {
  return os.Remove(os.Getenv("SCAF_CONFIG_DIR") + "/cookies.json")
}
