import {BACK_CONNECT, BACK_DISCONNECT, SETUP, UPSTREAM_CONNECT, UPSTREAM_DISCONNECT} from "./actionTypes";

function backConnect() {
  return {
    type: BACK_CONNECT,
  }
}

function backDisconnect() {
  return {
    type: BACK_DISCONNECT,
  }
}
function upstreamConnect() {
  return {
    type: UPSTREAM_CONNECT,
  }
}

function upstreamDisconnect() {
  return {
    type: UPSTREAM_DISCONNECT,
  }
}

function setup(gameSetup) {
  return {
    type: SETUP,
    data: gameSetup,
  }
}

export default {
  backConnect,
  backDisconnect,
  upstreamConnect,
  upstreamDisconnect,
  setup,
}