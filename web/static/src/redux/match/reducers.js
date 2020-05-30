import {GOAL, PANEL_UPDATE, STATE_CHANGE} from './actionTypes'

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

const defaultModal = {
  activate: false,
  mode: null,
}

const initialState = {
  panel: defaultPanel,
  modal: defaultModal,
  snapshot: defaultSnapshot,
  lastSnapshot: null,
}

export default function match(state = initialState, action) {
  switch (action.type) {
    case PANEL_UPDATE:
      return Object.assign({}, state, {
        panel: action.data,
        modal: defaultModal,
      })
    // case STATE_CHANGE:
    //   return Object.assign({}, state, {
    //     lastSnapshot: action.data,
    //     snapshot: action.data,
    //     modal: defaultModal,
    //   })
    case GOAL:
      return Object.assign({}, state, {
        modal: action.data
      })
    default:
      return state
  }
}