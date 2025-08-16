import {ALERT, DEBUG, GOAL, OVER, PANEL_UPDATE, REARRANGE, RESET, RESUME, BUFFERING, OVERTIME} from "./actionTypes";

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
function overtime() {
  return {
    type: OVERTIME,
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

function over(reason) {
  return {
    type: OVER,
    reason,
  }
}

export default {
  updatePanel,
  displayGoal,
  overtime,
  displayAlert,
  resume,
  pauseForDebug,
  pauseForRearrange,
  reset,
  over,
  buffering,
}