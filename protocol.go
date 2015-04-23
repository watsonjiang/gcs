package gcs

/*
   Protocol interface defines common behaviors for protocol layers.
   Each layer has to implement this interface
*/
type Protocol interface {
   /*
     return the name of protocol
    */
   Name() string
   /*
     An event was received from the layer below.
   */
   Up(e Event}) interface{}
   /*
     An event is  to be sent down the stack
   */
   Down(e Event) interface{}
}

/*
  protocol base
*/
type ProtocolBase struct {
   up_prot Protocol
   down_prot Protocol
}

