

class Channel {
  constructor() {
    this.listener = () => {console.error('no game state change listener')}
  }
  subscribe(cb) {
    this.listener = cb
  }

  newGameEvent(gameEvent) {
    this.listener(gameEvent)
  }
}

const channel = new Channel()

export default channel