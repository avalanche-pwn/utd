package srv

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type ClientSrv struct {
	conn   net.Conn
	seqNum int16
}

func (client *ClientSrv) Initialize(port string) error {
	var err error
	client.conn, err = net.Dial("tcp", port)
	client.seqNum = 0
	return err
}

func (client *ClientSrv) Close() {
	client.conn.Close()
}

func sendRequest[T any](client *ClientSrv, message Request[T]) {
	res, error := json.Marshal(message)
	if error != nil {
		panic("dupa")
	}
	length := len(res)

	fmt.Printf("Content-Length: %d\r\n\r\n%s", length, res)
	fmt.Fprintf(client.conn, "Content-Length: %d\r\n\r\n%s", length, res)
}

func (client *ClientSrv) InitializeDAP() {
	var req Request[InitializeRequestArguments]
	req.Seq = client.seqNum
	req.Type = "request"
	req.Command = "initialize"
	req.Arguments.ClientName = "utd"
	req.Arguments.AdapterId = "delve"
	req.Arguments.SupportsVariableType = false
	req.Arguments.SupportsVariablePaging = false
	req.Arguments.SupportsRunInTerminalRequest = false
	req.Arguments.SupportsMemoryReferences = false
	req.Arguments.SupportsProgressReporting = false
	req.Arguments.SupportsInvalidatedEvent = false
	req.Arguments.SupportsMemoryEvent = false
	req.Arguments.SupportsArgsCanBeInterpretedByShell = false
	req.Arguments.SupportsStartDebuggingRequest = false
	req.Arguments.SupportsANSIStyling = false
	sendRequest(client, req)
}

func (client *ClientSrv) Test() {
	time.Sleep(8 * time.Second)
	var ret []byte
	client.conn.Read(ret)
	fmt.Println(ret)
}
