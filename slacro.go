package main;

import (
  "fmt"
  "net/url"
  "net/http"
  "io/ioutil"

  redis "github.com/garyburd/redigo/redis"
);

const HASH = "hashmap";
var conn redis.Conn;

func main() {
  conn = GetRedisConn();
  conn.Do("HSET", HASH, "asd", "HELLOHELLO");

  http.HandleFunc("/", handler);
  http.ListenAndServe(":7777", nil);
}

func GetRedisConn() redis.Conn {
  c, err := redis.Dial("tcp", "localhost:6379");
  if err != nil {
      panic(err);
  }

  return c;
}

func getBody(r *http.Request) string {
  body, err := ioutil.ReadAll(r.Body);

  if err != nil {
    return "";
  }

  return string(body);
}

func IsSlackbot(v url.Values) bool {
  return v["user_id"][0] == "USLACKBOT";
}

func HasTrigger(text string) bool {
  return text[0] == '~';
}


func handler(w http.ResponseWriter, r *http.Request) {
  body := getBody(r);
  vals, err := url.ParseQuery(body);

  if err != nil || IsSlackbot(vals) {
    return;
  }

  text := vals["text"][0];
  if !HasTrigger(text) {
    return;
  }

  val, err := conn.Do("HGET", HASH, text[1:]);
  if err != nil {
    return;
  }

  x := val.([]uint8);

  var buf []byte;
  buf = make([]byte, len(x));

  for i := range x {
    buf[i] = x[i];
  }

  fmt.Println(val);
  fmt.Println(buf);
  fmt.Println("x");

  w.Write(buf);
}
