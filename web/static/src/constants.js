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


export {GameDefinitions, GameSettings};