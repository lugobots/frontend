import {RESET, ALERT, DEBUG, GOAL, PANEL_UPDATE, RESUME, REARRANGE, OVER} from './actionTypes'
import {StadiumStatus} from "../../constants";


const defaultPanel = {
  time_remaining: "--:--",
  shot_time: "--:--",
  home_score: 0,
  away_score: 0,
  team_goal: "",
}
const defaultSnapshot = {
  turn: 0,
  home_team: {
    players: [],
    score: 0,
  },
  away_team: {
    players: [],
    score: 0,
  },
  ball: {
    position: {
      x: 0,
      y: 0,
    }
  }
}


const initialState = {
  status: StadiumStatus.PLAYING,
  panel: defaultPanel,
  event_data: null,
  snapshot: defaultSnapshot,
  lastSnapshot: null,
}

export default function match(state = initialState, action) {
  switch (action.type) {
    case RESET:
      return initialState
    case RESUME:
      return Object.assign({}, state, {
        status: initialState.status,
        event_data: null,
      })
    case PANEL_UPDATE:
      return Object.assign({}, state, {
        panel: action.data,
      })
    case GOAL:
      return Object.assign({}, state, {
        status: StadiumStatus.GOAL,
        event_data: {team_goal: action.team_side},
      })
    case ALERT:
      return Object.assign({}, state, {
        status: StadiumStatus.ALERT,
        event_data: {modal: action.modal},
      })
    case DEBUG:
      return Object.assign({}, state, {
        status: StadiumStatus.DEBUGGING,
        event_data: null,
      })
    case REARRANGE:
      return Object.assign({}, state, {
        status: StadiumStatus.REARRANGING,
        event_data: null,
      })
    case OVER:
      return Object.assign({}, state, {
        status: StadiumStatus.OVER,
        event_data: null,
      })
    default:
      return state
  }
}