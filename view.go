package gcs

/*
  ViewIds are used for ordering views (each view has a ViewId
  and a list of members
*/
type ViewId struct {
   creator Address
   id      int64
}

func (id *ViewId) CompareTo(o *ViewId) int{
   return 0
}

/*
  A view is a local representation of the current membership of a group.
  Views contain the address of its creator, an ID and a list of member
  addresses. these addresses are ordered, and the first address is always
  the coordinator of the view. this way, each member of the group knows
  who the new coordinator will be if the current one crashes or leaves 
  the group. The views are sent between members using the VIEW_CHANGE
  event.
*/
type View struct {
   view_id ViewId
   members []Address
}

func (v *View) ContainsMember(mbr Address) bool {
   if mbr==nil {
      return false;
   }
   for _,m:=range members {
      if mbr.CompareTo(m)==0 {
         return true
      }
   }
   return false
}

func (v *View) CompareTo(o *ViewId) int {
   return 0
}

/*
   Returns a list of members which left from view one to two
*/
func (v *View) LeftMembers(one, two View) []Address {
   if one==nil || two==nil {
      return nil
   }
   return nil
}

/*
   Returns the difference between 2 views 'from' and 'to'.
   It assumes that view 'from' is logically prior to view 
   'to'.
*/
func (v *View) Diff(from, to View) [][]Address {

}


