

package main

import (
	"fmt"
	"time"
	"strconv"
	bw "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	//Connect
	cl := bw.ConnectOrExit("")
	cl.SetEntityFromEnvironOrExit()

  uri := "scratch.ns/temperature/"
	svc := cl.RegisterService(uri, "s.tempCtrl")

	//This sets a metadata key on the service
	svc.SetMetadata("tempCtrlApp", "voting for temperature")

	//You can have multiple interfaces per service. The second parameter
	//is the interface type, the first is the name of that instance of the
	//interface. We only have one interface, so underscore is a placeholder
	iface := svc.RegisterInterface("ctrl", "i.echo")

	// assume temperature is always in 'F
	var temperature float64
	temperature = 72

	//People can set what the message should be
	iface.SubscribeSlot("vote", func(m *bw.SimpleMessage) {
		if newmsg := m.GetOnePODF(bw.PODFString); newmsg != nil {
      fmt.Println("got a new vote")
			temp, err := strconv.ParseFloat(newmsg.(bw.TextPayloadObject).Value(), 64)
			if err == nil {
				fmt.Println(" ", temp)
				temperature = 0.7*temperature + 0.3*temp
			}
			// else {
			// 	fmt.Println("ERROR while parsing vote")
			// }
		}
	})

	//Also, every five seconds, we publish the message
	for {
    fmt.Println("current temperature is ", temperature)
		po := bw.CreateTextPayloadObject(bw.PONumString, strconv.FormatFloat(temperature, 'f', 3, 32))
		err := iface.PublishSignal("current", po)
		fmt.Println("Published, error was", err)
		time.Sleep(2 * time.Second)
	}
}
