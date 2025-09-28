package srv

import "encoding/json"

type ProtocolMessage struct {
	Seq  int16  `json:"seq"`
	Type string `json:"type"`
}

type Request[T any] struct {
	ProtocolMessage
	Command   string `json:"command"`
	Arguments T      `json:"arguments,omitempty"`
}

type Event struct {
	ProtocolMessage
	Event string `json:"event"`
	Body  string `json:"body,omitempty"`
}

type Response[T any] struct {
	ProtocolMessage
	RequestSeq int16  `json:"request_seq"`
	Success    bool   `json:"success"`
	Command    string `json:"command"`
	Message    string `json:"message,omitempty"`
	Body       T
}

type BaseRespone = Response[json.RawMessage]

type internalResponse struct {
	Err    error
	Result *BaseRespone
}

type InitializeRequestArguments struct {
	ClientID                            string `json:"clientID,omitempty"`
	ClientName                          string `json:"clientName,omitempty"`
	AdapterId                           string `json:"adapterID"`
	Locale                              string `json:"locale,omitempty"`
	LinesStartAt1                       bool   `json:"lineStartAt1,omitempty"`
	ColumnsStartAt1                     bool   `json:"columnsStartAt1,omitempty"`
	PathFormat                          string `json:"pathFormat,omitempty"`
	SupportsVariableType                bool   `json:"supportsVariableType,omitempty"`
	SupportsVariablePaging              bool   `json:"supportsVariablePaging,omitempty"`
	SupportsRunInTerminalRequest        bool   `json:"supportsRunInTerminalRequest,omitempty"`
	SupportsMemoryReferences            bool   `json:"supportsMemoryReferences,omitempty"`
	SupportsProgressReporting           bool   `json:"supportsProgressReporting,omitempty"`
	SupportsInvalidatedEvent            bool   `json:"supportsInvalidatedEvent,omitempty"`
	SupportsMemoryEvent                 bool   `json:" supportsMemoryEvent,omitempty"`
	SupportsArgsCanBeInterpretedByShell bool   `json:"supportsArgsCanBeInterpretedByShell,omitempty"`
	SupportsStartDebuggingRequest       bool   `json:"supportsStartDebuggingRequest,omitempty"`
	SupportsANSIStyling                 bool   `json:"supportsANSIStyling,omitempty"`
}
