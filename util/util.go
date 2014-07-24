package util

import (
  "log"
  "strconv"
  "time"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

func ReadJson(r *http.Request, v interface{}) (err error) {
    defer r.Body.Close()

    var (
        body []byte
    )

    body, err = ioutil.ReadAll(r.Body)

    if err != nil {
        log.Printf("ReadJson couldn't read request body %v", err)
        return err
    }

    if err = json.Unmarshal(body, v); err != nil {
        log.Printf("ReadJson couldn't parse request body %v", err)
        return err
    }

    return err
}

func WriteJson(w http.ResponseWriter, v interface{}) {
    // avoid json vulnerabilities, always wrap v in an object literal

    if data, err := json.Marshal(v); err != nil {
        log.Printf("Error marshalling json: %v", err)
    } else {
        w.Header().Set("Content-Length", strconv.Itoa(len(data)))
        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    }
}

func TodayAsInteger() (int) {
  todayAsInt, err := strconv.Atoi(time.Now().Local().Format("20060102"))
  if(err != nil) { panic(err) }
  return todayAsInt
}