const GameDefinitions = {
  Field: {
    Width: 20000,
    Height: 10000,
  },
  Player: {
    Size:400,
  }
}

const GameSettings = {
  START_MODE: {
    WAIT: 0,
    NO_WAIT: 1,
    DELAY: 2,
  },
  LISTENING_MODE: {
    // respect the timer defined by listening duration
    TIMER: 0,
    // immediately after all orders
    RUSH: 1,
    // wait external remote control (dev only)
    REMOTE: 2,
  }
}

const GameStates = {
  WAITING: 0,
  GET_READY: 1,
  LISTENING: 2,
  PLAYING: 3,
  SHIFTING: 4,
  OVER: 99,
}

const AppStatus = {
  Connecting: "connecting",
  Setting: "setting-up",
  Listening: "listening",
  UpstreamDown: "upstream-down",
  Debugging: "debugging",
}

const StadiumStatus = {
  PLAYING: "PLAYING",
  GOAL: "GOAL",
  ALERT: "ALERT",
}

// const ModalModes = {
//   GOAL: "goal",
//   ALERT: "alert",
// }

// const StadiumStates = {
//   StadiumStateConnecting: "connecting",
//   StadiumStateSetting: "setting-up",
//   StadiumStateListening: "listening",
//   StadiumStateConn: "conn-upstream",
//   StadiumStateGoal: "goal",
//
//   // dev mode
//   StadiumStateDebugging: "debugging",
// }

const EventTypes = {
  NewPlayer: "new_player",
  Breakpoint: "breakpoint",
  StateChange: "state_change",
  DebugReleased: "debug_released",
  Goal: "goal",
  GameOver: "game_over",
  LostPlayer: "lost_player",
  ConnectionLost: "connection_lots",
  ConnectionReestablished: "connection_Reestablished",
}

const BreakpointType = {
  ORDERS: 0,
  TURN: 1,
}


const BackendConfig = {
  BackEndPoint: "http://localhost:8080",
}
export {
  GameDefinitions,
  GameSettings,
  GameStates,
  // StadiumStates,
  EventTypes,
  BreakpointType,
  BackendConfig,
  AppStatus,
  // ModalModes,
  StadiumStatus,
};