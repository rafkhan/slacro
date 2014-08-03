package main;

import (
  "fmt"
  "net/url"
  "testing"
)

func TestSlackbotCheck(t *testing.T) {
  queryBot := "?asd=raf&zxc=42&user_id=USLACKBOT&qwe=7"
  queryNoBot := "?asd=raf&user_id=1234567890&qwe=7&zxc=42"

  valsBot, errBot := url.ParseQuery(queryBot);
  if errBot != nil { t.Error(errBot); }

  valsNoBot, errNoBot := url.ParseQuery(queryNoBot);
  if errNoBot != nil { t.Error(errNoBot); }

  if !IsSlackbot(valsBot) { t.Error("Should have found slackbot ID."); }
  if IsSlackbot(valsNoBot) { t.Error("User is not a slackbot."); }
}

func TestTriggerText(t *testing.T) {
  a := "~lol";
  b := "lol~";
  c := "l~ol";

  if !HasTrigger(a) { t.Errorf("%s should trigger.", a); }
  if HasTrigger(b) { t.Errorf("%s should not trigger.", b); }
  if HasTrigger(c) { t.Errorf("%s should not trigger.", c); }
}

func TestRedisConn(t *testing.T) {
  c := GetRedisConn();
  defer c.Close();

  resp, err := c.Do("PING");
  if err != nil || resp != "PONG" {
    t.Error("No redis connection.")
  }
}

func TestGetImage(t *testing.T) {
  key := "lie";
  val := "lie.png";

  conn := GetRedisConn();
  defer conn.Close();

  conn.Do("HSET", HASH, key, val);

  resp := GetImage(conn, key);
  if resp != val {
    t.Errorf("HGET %s %s should return %s but returned %s instead.",
        HASH, key, val, resp);
  };
}

// I should actually write a test for this...
func TestGenerateResp(t *testing.T) {
  ICON = fmt.Sprintf("%s%s/public/%s", HOST, PORT, "lie.png");
  fmt.Println(GenerateResp("lie.png", NAME, ICON));
}
