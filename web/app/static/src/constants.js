const GameDefinitions = {
  Field: {
    Width: 20000,
    Height: 10000,
  },
  Player: {
    Size:400,
  },
  BallTimeInGoalZone: 40, // two sec
}

const GameSettings = {
  START_MODE: {
    WAIT: 'WAIT',
    NO_WAIT: 'NO_WAIT',
    DELAY: 'DELAY',
  },
  LISTENING_MODE: {
    // respect the timer defined by listening duration
    TIMER: 'TIMER',
    // immediately after all orders
    RUSH: 'RUSH',
    // wait external remote control (dev only)
    REMOTE: 'REMOTE',
  }
}

const GameStates = {
  WAITING: 'WAITING',
  GET_READY: 'GET_READY',
  LISTENING: 'LISTENING',
  PLAYING: 'PLAYING',
  SHIFTING: 'SHIFTING',
  OVER: 'OVER',
}

const AppStatus = {
  Connecting: "connecting",
  Setting: "setting-up",
  Listening: "listening",
  UpstreamDown: "upstream-down",
  Broken: "broken",
}

const EventTypes = {
  NewPlayer: "new_player",
  Breakpoint: "breakpoint",
  StateChange: "state_change",
  DebugReleased: "debug_released",
  Goal: "goal",
  GameOver: "game_over",
  LostPlayer: "lost_player",
  ConnectionLost: "connection_lots",
  ConnectionReestablished: "connection_reestablished",
  Buffering: "buffering",
  BufferReady: "buffer_ready",
}

const StadiumStatus = {
  PLAYING: "PLAYING",
  GOAL: "GOAL",
  ALERT: "ALERT",
  DEBUGGING: "DEBUGGING",
  REARRANGING: "REARRANGING",
  OVER: "OVER",
  BUFFERING: "BUFFERING",
}

const BreakpointType = {
  ORDERS: 'ORDERS',
  TURN: 'TURN',
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
