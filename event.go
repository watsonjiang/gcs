package gcs

const (
   EVT_MSG = iota
   EVT_CONNECT
   EVT_DISCONNECT
   EVT_VIEW_CHANGE
   EVT_SUSPECT
   EVT_UNSUSPECT
)

type Event struct {
   Type int
   Arg  interface{}
}
