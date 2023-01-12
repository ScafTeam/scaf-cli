package scafreq

import (
  "os"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "errors"
)

func LoadCookie(name string) (*http.Cookie, error) {
  cookies, err := LoadCookies()
  if err != nil {
    return nil, err
  }
  for _, cookie := range cookies {
    if cookie.Name == name {
      return cookie, nil
    }
  }

  return nil, errors.New("cookie not found")
}

func LoadCookies() ([]*http.Cookie, error) {
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

func SaveCookies(new_cookies []*http.Cookie) error {
  cookies, err := LoadCookies()
  if err != nil {
    return err
  }
  for _, cookie := range new_cookies {
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
  err = WriteCookies(cookies)
  if err != nil {
    return err
  }

  return nil
}

func WriteCookies(cookies []*http.Cookie) error {
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
func DeleteCookies(names []string) error {
  cookies, err := LoadCookies()
  if err != nil {
    return err
  }
  new_cookies := []*http.Cookie{}
  for _, cookie := range cookies {
    var found bool = false
    for _, name := range names {
      if cookie.Name == name {
        found = true
        break
      }
    }
    if !found {
      new_cookies = append(new_cookies, cookie)
    }
  }
  err = WriteCookies(new_cookies)
  if err != nil {
    return err
  }

  return nil
}

func DeleteAllCookies() error {
  return WriteCookies([]*http.Cookie{})
}

func LoadCookieValue(name string) (string, error) {
  cookie, err := LoadCookie(name)
  if err != nil {
    return "", err
  }

  return cookie.Value, nil
}
