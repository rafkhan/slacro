package main;

import (
  "net/url"
  "testing"
)

func TestSlackbotCheck(t *testing.T) {
  queryBot := "?asd=raf&zxc=42&user_id=USLACKBOT&qwe=7"
  queryNoBot := "?asd=raf&user_id=1234567890&qwe=7&zxc=42"

  valsBot, errBot := url.ParseQuery(queryBot);
  if errBot != nil {
    t.Error(errBot);
  }

  valsNoBot, errNoBot := url.ParseQuery(queryNoBot);
  if errNoBot != nil {
    t.Error(errNoBot);
  }

  if !IsSlackbot(valsBot) {
    t.Error("Should have found slackbot ID.");
  }

  if IsSlackbot(valsNoBot) {
    t.Error("User is not a slackbot.");
  }
}
