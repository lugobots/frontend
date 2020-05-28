import {BACK_CONNECT, BACK_DISCONNECT, UPSTREAM_CONNECT, UPSTREAM_DISCONNECT} from './actionTypes'
import {AppStatus} from './constants'

const initialState = {
  status: AppStatus.Connecting,
  // isConnected: false,
  // isSetup: false,
  // upstream_up: false,
  error: null,
}

export default function game(state = initialState, action) {
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
    default:
      return state
  }
}