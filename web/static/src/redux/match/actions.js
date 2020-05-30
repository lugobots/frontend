import {ALERT, GOAL, PANEL_UPDATE} from "./actionTypes";
import {ModalModes} from "../../constants";

function updatePanel(data) {
  return {
    type: PANEL_UPDATE,
    data
  }
}

function displayGoal(team_side) {
  return {
    type: GOAL,
    data: {
      activate: true,
      mode: ModalModes.GOAL,
      team_side,
    }
  }
}

function displayModal(title, text) {
  return {
    type: ALERT,
    data: {
      activate: true,
      mode: ModalModes.ALERT,
      title,
      text
    }
  }
}

export default {
  updatePanel,
  displayGoal,
  displayModal,
}