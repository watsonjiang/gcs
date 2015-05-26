package gcs

import (
   "testing"
)

func TestAddrComp1(t *testing.T) {
   a := &IpAddress{Ip:"127.0.0.1", Port:9900}
   b := &IpAddress{Ip:"127.0.0.2", Port:9900}
   if a.CompareEqual(b) {
      t.Error("address compare not correct!")
   }
}
