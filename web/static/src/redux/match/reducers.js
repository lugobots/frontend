import {PANEL_UPDATE, STATE_CHANGE} from './actionTypes'

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
    Score: 0,
  },
  away_team: {
    players: [],
    Score: 0,
  },
  ball: {
    Position: {
      X: 0,
      Y: 0,
    }
  }
}

const initialState = {
  panel: defaultPanel,
  snapshot: defaultSnapshot,
  lastSnapshot: null,
}

export default function match(state = initialState, action) {
  switch (action.type) {
    case PANEL_UPDATE:
      return Object.assign({}, initialState, {
        panel: action.data,
      })
    case STATE_CHANGE:
      return Object.assign({}, initialState, {
        lastSnapshot: action.data,
        snapshot: action.data,
      })
    default:
      return state
  }
}