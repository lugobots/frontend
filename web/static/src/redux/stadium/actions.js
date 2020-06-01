import {ALERT, DEBUG, GOAL, PANEL_UPDATE, RESUME} from "./actionTypes";

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

export default {
  updatePanel,
  displayGoal,
  displayAlert,
  resume,
  pauseForDebug,
}