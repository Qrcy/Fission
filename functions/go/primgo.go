package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "math"
    "strconv"
)

type msg struct {
    ID string `json:"id"`
}

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func FindPrime(x int) int {
    var temp int = 0
        for i := 0; i < x; i++ {
                if (IsPrime(i) != false) {
                        temp = i
                }
        }
    return temp
}

func Handler(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    m := msg{}
    err = json.Unmarshal(reqBody, &m)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    i, err := strconv.Atoi(m.ID)

    c := FindPrime(i)
    
    respBody, err := json.Marshal(c)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
    }

    w.Write([]byte(respBody))
}

