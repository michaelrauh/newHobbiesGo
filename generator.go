package main

import (
  "math/rand"
  "time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Generate(length int) string {
  b := make([]rune, length)
  rand.Seed(time.Now().UTC().UnixNano())
  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}
