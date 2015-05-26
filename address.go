package gcs

type Address interface {
   //compare two addresses.
   //return 0 for equality, -1 for smaller, 1 for bigger
   CompareTo(o Address) int
}

type IpAddress struct {
   Ip string
   Port int
}

func (a *IpAddress) CompareTo(o Address) int {
   b := o.(*IpAddress)
   if a.Ip == b.Ip {
      if a.Port == b.Port {
         return 0
      }else if a.Port < b.Port {
         return -1
      }
      return 1
   }else if a.Ip > b.Ip {
      return 1
   }
   return 0
}
