package gcs

/*
  ViewIds are used for ordering views (each view has a ViewId
  and a list of members
*/
type ViewId struct {
   creator Address
   id      int64
}

/*
 Establishes an order between 2 ViewIds. The comparison is done 
 on the IDs, if they are equal, use the creator.
 return 0 for equality, value less than 0 if smaller, greater than 0 if greater.
*/
func (id *ViewId) CompareTo(o *ViewId) int{
   if id.id > o.id {
      return 1
   }else if id.id < o.id {
      reutrn -1
   }else if id.id == o.id {
      return id.createor.CompareTo(o.creator)
   }
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
   for _,m:=range v.members {
      if mbr.CompareEqual(m) {
         return true
      }
   }
   return false
}

func (v *View) CompareTo(o *ViewId) int {
   return 0
}
