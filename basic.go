package gcs

/*
  this file contains some basic facilities for gcs.
*/

type Streamable interface {
   WriteTo(w io.Writer) error
   ReadFrom(r io.Reader) error
}
