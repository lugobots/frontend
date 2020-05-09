const GameDefinitions = {
  Field: {
    Width: 20000,
    Height: 10000,
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

export {GameDefinitions, GameSettings, GameStates};