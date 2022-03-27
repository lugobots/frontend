import {ALERT, DEBUG, GOAL, OVER, PANEL_UPDATE, REARRANGE, RESET, RESUME, BUFFERING} from "./actionTypes";

function reset() {
  return {
    type: RESET,
  }
}

function buffering(percentile) {
  return {
    type: BUFFERING,
    percentile
  }
}

function updatePanel(data) {
  return {
    type: PANEL_UPDATE,
    data
  }
}

function pauseForDebug() {
  return {
    type: DEBUG,
  }
}

function pauseForRearrange() {
  return {
    type: REARRANGE,
  }
}

function displayGoal(team_side) {
  return {
    type: GOAL,
    team_side,
  }
}

function displayAlert(title, text) {
  return {
    type: ALERT,
    modal: {
      title,
      text
    }
  }
}

function resume() {
  return {
    type: RESUME,
  }
}

function over() {
  return {
    type: OVER,
  }
}

export default {
  updatePanel,
  displayGoal,
  displayAlert,
  resume,
  pauseForDebug,
  pauseForRearrange,
  reset,
  over,
  buffering,
}