import $ from "jquery";
import {GameDefinitions} from "./constants";

function ShouldRender(currProps, nextProps) {
  return nextProps.v !== currProps.v
}

let ratio = 1;
function getSizeRatio() {
  return ratio
}
function updateRatio() {
  ratio = $("#field").innerWidth() / GameDefinitions.Field.Width
}

function renderLogger(className) {
   console.log(`${className} rendered`)
}


export {ShouldRender, renderLogger, getSizeRatio, updateRatio};