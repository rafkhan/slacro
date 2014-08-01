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

func getBody(r *http.Request) string {
  body, err := ioutil.ReadAll(r.Body);

  if err != nil {
    return "";
  }

  return string(body);
}

func isSlackbot(v url.Values) bool {
  return v["user_id"][0] == "USLACKBOT";
}

func handler(w http.ResponseWriter, r *http.Request) {
  body := getBody(r);
  vals, err := url.ParseQuery(body);

  if err != nil || isSlackbot(vals) {
    return;
  }

}

func makeSlackReq(v url.Values) (*http.Response, error) {
  url := "https://slack.com/api/chat.postMessage?" + v.Encode();
  resp, err := http.Get(url);
  return resp, err;
}
