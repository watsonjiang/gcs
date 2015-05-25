package gcs

import "testing"

func TestAddrComp1(t *t.Testing) {
   a := &IpAddress{Ip:"127.0.0.1", Port:9900}
   b := &IpAddress{Ip:"127.0.0.2", Port:9900}
   if a.CompareTo(b) {
      t.Error("address compare not correct!")
   }
}
