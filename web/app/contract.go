package app

import "github.com/lugobots/lugo4go/v2/lugo"

type EventsBroker interface {
	StreamEventsTo(uuid string) chan FrontEndUpdate
	GetGameConfig() *lugo.GameSetup
}

const (
	EventNewPlayer     = "new_player"
	EventBreakpoint    = "breakpoint"
	EventStateChange   = "state_change"
	EventDebugReleased = "debug_released"
	EventGameOver      = "game_over"
	EventLostPlayer    = "lost_player"
)

// Error is used to identify internal errors
type Error string

// Error implements the native golang error interface
func (e Error) Error() string { return string(e) }

const (
	// ErrGRPCConnectionClosed identifies when the error returned is cased by the connection has been closed
	ErrGRPCConnectionClosed = Error("grpc connection closed by the server")

	// ErrGRPCConnectionLost identifies that something unexpected broke the gRPC connection
	ErrGRPCConnectionLost = Error("grpc stream error")

	// ErrUnknownGameEvent identifies an error when the service receives an event of an unknown type
	ErrUnknownGameEvent = Error("unknown event type")

	ErrGameOver = Error("the game is over")

	ErrStopRequested = Error("it was requested to stop")
)
