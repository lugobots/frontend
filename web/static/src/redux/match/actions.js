import {PANEL_UPDATE} from "./actionTypes";

export function updatePanel(data) {
  return {
    type: PANEL_UPDATE,
    data
  }
}

export default {
  updatePanel,
}