package main;

import (
  "net/url"
  "net/http"
  "io/ioutil"

  redis "github.com/garyburd/redigo/redis"
);

const HASH = "hashmap";
var conn redis.Conn;

func main() {
  conn = GetRedisConn();
  conn.Do("HSET", HASH, "lie", "lie.png");

  http.HandleFunc("/message", handler);
  http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w, r, r.URL.Path[1:]);
  });

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

func GetImage(conn redis.Conn, key string) string {
  val, err := conn.Do("HGET", HASH, key);
  if err != nil {
    return "";
  }

  // Find a better way to do this
  x := val.([]uint8);

  var buf []byte;
  buf = make([]byte, len(x));

  for i := range x {
    buf[i] = byte(x[i]);
  }

  return string(buf);
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

  img := GetImage(conn, text[1:]);

  w.Write([]byte("{\"text\":\"http://104.131.214.140/public/" + img + "\"}"));
}
