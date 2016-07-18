// Simple test application to change the echo message on the bw2 example hellosvc

package main

import (
  "fmt"
  //"github.com/immesys/spawnpoint/spawnable"
  bw "gopkg.in/immesys/bw2bind.v5"
)

func main() {
  // connect
  cl := bw.ConnectOrExit("")
  cl.SetEntityFromEnvironOrExit()
  uri := "scratch.ns/temperature/s.tempCtrl/ctrl/i.echo/slot/vote"
	fmt.Println("Enter your desired temperature")
	var temperature string
	fmt.Scanf("%s",&temperature)
  po := []bw.PayloadObject{bw.CreateStringPayloadObject(temperature),}
  fmt.Println("made a po, ready to publish")
  err :=  cl.Publish(&bw.PublishParams{
                                  URI : uri,
                            			AutoChain : true,
                            			PayloadObjects : po,
                            	    })
  fmt.Println("published, err was", err)
}
