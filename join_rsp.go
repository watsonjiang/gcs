// +build ignore

package gcs

type JoinRsp struct {
   view *View
   digest *Digest
   failReason string
}

func (r *JoinRsp) NewFailJoinRsp(fail_reason string) *JoinRsp {
   return &JoinRsp{
           fail_reason:fail_reason,
        }
}

func (r *JoinRsp) NewJoinRsp(v *View, d *Digest) *JoinRsp {
   return &JoinRsp{
           view:v,
           digest:d,
        }
}

func (r *JoinRsp) WriteTo(w io.Writer) error {
   return nil
}

func (r *JoinRsp) ReadFrom(w io.Reader) error {
   return nil
}

func (r *JoinRsp) String() string {
   return nil
}
