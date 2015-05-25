package gcs

type Address interface {
   CompareTo(o Address) bool
}

type IpAddress struct {
   Ip string
   Port int
}

func (a *IpAddress) CompareTo(o Address) bool {
   b := o.(*IpAddress)
   if a.Ip != b.Ip {
      return a.Ip > b.Ip
   }else if a.Port != b.Port {
      return a.Port > b.Port
   }else {
      return false
   }
}
