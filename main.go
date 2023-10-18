package main


import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urlMap := make(map[string]string)

func main(){
	http.HandleFunc("/", HandleFrm)
	http.HandleFunc("/shorturl", HandleShort)
	http.HandleFunc("/short/", HandleRdirect)
	http.ListenAndServe("":3030", nil)
}