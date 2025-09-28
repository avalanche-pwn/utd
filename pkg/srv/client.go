package srv

import (
	"encoding/json"
	"fmt"
	"net"
	"log/slog"
)

const MaxConcurentRequests = 100

type ClientSrv struct {
	conn         net.Conn
	seqNum       int16
	exit         chan bool
	sendQueue    chan Request[any]
	recvRawQueue chan string
}

func (client *ClientSrv) Initialize(port string) error {
	slog.Info("Initializing tcp connection")
	var err error
	client.conn, err = net.Dial("tcp", port)
	client.seqNum = 0
	client.exit = make(chan bool)
	client.recvRawQueue = make(chan string)
	return err
}

func (client *ClientSrv) Close() {
	client.exit <- true
	client.conn.Close()
}

func sendRequest[T any](client *ClientSrv, message Request[T]) int16 {
	message.Seq = client.seqNum
	client.seqNum += 1
	client.seqNum %= MaxConcurentRequests

	// client.responses[message.Seq] <- internalResponse{errors.New("Timeout"), nil}
	// client.responses[message.Seq] = make(chan internalResponse)
	res, error := json.Marshal(message)
	if error != nil {
		panic("dupa")
	}
	length := len(res)

	slog.Debug("Sending", "data", res)
	fmt.Fprintf(client.conn, "Content-Length: %d\r\n\r\n%s", length, res)
	return message.Seq
}

func (client *ClientSrv) InitializeDAP() {
	slog.Info("Initializing DAP protocol")
	var req Request[InitializeRequestArguments]
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

func (client *ClientSrv) reader() {
	for {
		var length int
		fmt.Fscanf(client.conn, "Content-Length: %d\r\n\r\n", &length)
		slog.Debug("Receiving data", slog.Int("excepted_length", length))

		buffer := make([]byte, length)
		client.conn.Read(buffer)
		client.recvRawQueue <- string(buffer[:])
	}
}

func (client *ClientSrv) recvJson(json string) {
	slog.Debug("Received", slog.String("data", json))
}

func (client *ClientSrv) ServeBlocking() {
	go client.reader()
	for {
		select {
		case message := <-client.sendQueue:
			sendRequest(client, message)
		case message := <-client.recvRawQueue:
			client.recvJson(message)
		case <-client.exit:
			return
		}
	}
}

func (client *ClientSrv) Serve() {
	go client.ServeBlocking()
}
