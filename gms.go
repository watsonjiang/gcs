// +build ignore

package gcs

import(
  "errors"
  "io"
)

/*
   type of GmsHeader
*/
const (
   GMS_HEADER_JOIN_REQ = iota
   GMS_HEADER_JOIN_RSP
   GMS_HEADER_LEAVE_REQ
   GMS_HEADER_LEAVE_RSP
)

type GmsHeader struct {
   Type byte
   Mbr  Address
}

func NewGmsHeader() *GmsHeader {
   return &GmsHeader{}
}

func (h *GmsHeader) WriteTo(w io.Writer) error {

}

func (h *GmsHeader) ReadFrom(r io.Reader) error {

}

//JoinRsp


//Request
type Request struct {

}

//------------role implementation-----------
var ERR_METHOD_NOT_IMPLEMENTED = errors.New("Method not implemented")
//base of different role in GMS.(coordinator, participant, client)
type GmsRole interface {
   Join(mbr Address) error
   Leave(mbr Address) error
   HandleJoinResponse(rsp JoinRsp) error
   HandleLeaveResponse() error
   Suspect(mbr Address) error
   Unsuspect(mbr Address) error

   HandleMembershipChange(requests []Request) error
   HandleViewChange(new_view View, digest Digest) error
}

/*
  Coordinator role, Accesspts JOIN and LEAVE request and emits view changes
  acordingly.
*/
type GmsCoordinator struct {
}

func (r *GmsCoordinator) Join(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) Leave(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) HandleJoinResponse(rsp JoinRsp) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) HandleLeaveResponse() error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) Suspect(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) Unsuspect(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) HandleMembershipChange(requests []Request) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsCoordinator) HandleViewChange(new_view View, digest Digest) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

//participant role
type GmsParticipant struct {
}

func (r *GmsParticipant) Join(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) Leave(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) HandleJoinResponse(rsp JoinRsp) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) HandleLeaveResponse() error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) Suspect(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) Unsuspect(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) HandleMembershipChange(requests []Request) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsParticipant) HandleViewChange(new_view View, digest Digest) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

//----------protocol handler in stack---------------
type ViewHandler struct {}

/*
   Group membership protocol. Handles joins/leaves/crashes 
   and emits new views accordingly.
 */
type GMS struct {
   view_handler ViewHandler
}

func (p *GMS) Up(e Event) interface{} {
   switch e.Type {
      case EVT_MSG:
      case EVT_SUSPECT:
         ret:=p.up_prot.Up(e)
         addr:=e.Arg.(*Address)
      case EVT_UNSUSPECT:
   }
   return up_prot.Up(e)
}

func (p *GMS) Down(e Event) interface{} {

}
