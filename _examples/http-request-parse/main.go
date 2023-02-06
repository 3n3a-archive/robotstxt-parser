package main

import (
  "fmt"
  "github.com/3n3a/robotstxt-parser"
  "net/http"
  "log"
  "io/ioutil"
)

func main() {
  resp, err := http.Get("https://www.suva.ch/robots.txt")
  if err != nil {
    log.Fatalln(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  inputTxt := string(body)
  st, err := RobotsTxtParser.ParseTxt(inputTxt)
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Printf("%#v\n", st)
}