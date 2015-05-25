package gcs

import "io"

/*
  A message digest containing -for each member- the highest seqno delivered(hd)
  and the highest seqno recieved(hr)
  The seqnos are stroed according to the order of the members in the associated 
  membership, ie. seqno[0] is the hd for members[0], seqno[1] is the hr for the 
  same member,seqnos[2] is the hd for member[1] and so.
*/
type Digest struct {
   members []Address
   seqnos  []uint64
}

//implements Streamable interface
func (d *Digest) WriteTo(w io.Writer) error {
   return nil
}

func (d *Digest) ReadFrom(r io.Reader) error {
   return nil
}
