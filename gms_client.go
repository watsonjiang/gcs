// +build ignore

package gcs

import(
  //"github.com/golang/glog"
)

/*
  Client role. Whenever a new member wants to join a group, it starts in the 
  CLIENT role. No multicasts to the group will be received and processed until 
  the member has been joined and truned into a SERVER(coordinator or participant)
*/
type GmsClient struct {
   joinRspChan chan *JoinRsp
   joinPromise *Promise
   leaving bool
}

/*
  Joins the process to a group. Determines the coordinator and sends a 
  unicast handleJoin() message to it. The coordinator returns a JoinRsp 
  and the broadcasts the new view, which contains a message digest and 
  current membership(including the joiner). The joiner is then supposed 
  to install the new view and the digest and starts accepting mcast
  message. Previous mcast messages were discarded.
*/
func (r *GmsClient) Join(mbr Address) error {
   joinAttempts := 0
   r.leaving = false
   r.joinPromise.Reset()
   for !r.leaving {
      if r.installViewIfValidJoinRsp(r.joinPromise, false) {
         return nil
      }
      //TODO:fetch member addresses from somewhere
      var mbrs []Address
      coords = getCoords(mbrs)

      //no coord, all are clients. if I'm the first of sorted client,
      //I'll become coordinator, The other will wait and retry the discovery
      //and join process
      if len(coords)==0 {
         //if r.firstOfAllClients(mbr, 
      }

   }
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) getCoords(mbrs []Address) []Address {
  return nil 
}

/*
  Return true if joiner is the first member of the group
*/
func (r *GmsClient) firstOfAllClients(joiner Address, mbrs []Address) bool {
  return false 
}

type ByAdr []Address
func (a ByAdr) Len() int {
   return len(a)
}

func (a ByAdr) Swap(i, j int) {
   a[i], a[j] = a[j], a[i]
}

func (a ByAdr) Less(i, j int) bool {
   return a[i].CompareTo(a[j])
}

func (r *GmsClient) Leave(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) HandleJoinResponse(rsp JoinRsp) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) HandleLeaveResponse() error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) Suspect(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) Unsuspect(mbr Address) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) HandleMembershipChange(requests []Request) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

func (r *GmsClient) HandleViewChange(new_view View, digest Digest) error {
   return ERR_METHOD_NOT_IMPLEMENTED
}

/*
  Installs a new view and sends an ack if (1) the join rsp is not null, (2) it is valid and (3)
  the view installation is successful. if true is returned, the JOIN process can be terminated,
  else it needs to JOIN-REQ again in a next iteration.
*/
func (r *GmsClient) installViewIfValidJoinRsp(joinPromise *Promise, blockForRsp bool) (succ bool) {
   var rsp JoinRsp
   var err error
   if joinPromise.HasResult() {
      rsp, err = joinPromise.GetResultWithTimeout(time.Millisecond)
   }else if(block_for_rsp) {
      rsp, err = joinPromise.GetResult()
   }
   defer func() {
      if succ {
         sendViewAck(rsp.GetView().GetCreator())
      }
   }()
   return (err==nil) && r.isJoinRspValid(rsp) && installView(rsp.GetView(), rsp.GetDigest())
}


