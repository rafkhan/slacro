package main;

import (
  "os"
  "fmt"
  "net/url"
  "net/http"
  "io/ioutil"
);

func main() {
  token := os.Getenv("SLACK_TOKEN");
  fmt.Println(token);

  /*
  makeSlackReq(url.Values{"token": {token},
                          "channel": {"#bot-test"},
                          "text": {"ayy"}});
  */

  http.HandleFunc("/", handler);
  http.ListenAndServe(":7777", nil);
}

func handler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body);

  if err != nil {
    return;
  }

  fmt.Println(string(body));
  fmt.Println(r.Header);
  w.Write([]byte("{\"text\":\"ayyy\"}"));
}

func makeSlackReq(v url.Values) (*http.Response, error) {
  url := "https://slack.com/api/chat.postMessage?" + v.Encode();
  resp, err := http.Get(url);
  return resp, err;
}
