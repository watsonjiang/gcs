package gcs

type Header interface {
   WriteTo(w io.Writer) error
   ReadFrom(r io.Reader) error
}

/*
  A Message encapsulates data sent to members of group.
*/
type Message struct {
   dest_addr Address
   src_addr  Address
   buf       []byte
   headers   []Header
   flags     uint16
}
