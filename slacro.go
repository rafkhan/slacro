package main;

import (
  "os"
  "fmt"
  "net/url"
  "net/http"
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
  w.Write([]byte("hello!"))
}

func makeSlackReq(v url.Values) (*http.Response, error) {
  url := "https://slack.com/api/chat.postMessage?" + v.Encode();
  resp, err := http.Get(url);
  return resp, err;
}
