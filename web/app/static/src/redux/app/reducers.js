import {BACK_CONNECT, BACK_DISCONNECT, BROKEN, SETUP, UPSTREAM_CONNECT, UPSTREAM_DISCONNECT} from './actionTypes'
import {AppStatus, GameSettings} from '../../constants'

const defaultSetup = {
  dev_mode: false,
  listening_mode: GameSettings.LISTENING_MODE.TIMER,
  start_mode: GameSettings.START_MODE.WAIT,
  listening_duration: 50,
  game_duration: 6000,
  home_team: {
    name: "Home",
    avatar: "images/profile-team-home.jpg",
    colors: {
      primary: {
        red: 240
      },
      secondary: {
        red: 255,
        green: 255,
        blue: 255
      }
    }
  },
  away_team: {
    name: "Away",
    side: 1,
    avatar: "images/profile-team-away.jpg",
    colors: {
      primary: {
        green: 200
      },
      secondary: {
        green: 240,
        blue: 240
      }
    }
  }
}

const initialState = {
  status: AppStatus.Connecting,
  setup: defaultSetup,
  // isConnected: false,
  // isSetup: false,
  // upstream_up: false,
  error: null,
}

export default function app(state = initialState, action) {
  switch (action.type) {
    case BACK_CONNECT:
      return Object.assign({}, initialState, {
        status: AppStatus.Setting,
      })
    case BACK_DISCONNECT:
      return initialState
    case UPSTREAM_CONNECT:
      return Object.assign({}, initialState, {
        status: AppStatus.Setting,
      })
    case UPSTREAM_DISCONNECT:
      return Object.assign({}, initialState, {
        status: AppStatus.UpstreamDown,
      })
    case SETUP:
      return Object.assign({}, state, {
        setup: action.data.game_setup,
        status: action.data.connection_state === 'up' ? AppStatus.Listening : AppStatus.UpstreamDown,
      })
    case BROKEN:
      return Object.assign({}, initialState, {
        status: AppStatus.Broken
      })
    default:
      return state
  }
}
