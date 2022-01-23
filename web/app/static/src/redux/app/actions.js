import {BACK_CONNECT, BACK_DISCONNECT, BROKEN, SETUP, UPSTREAM_CONNECT, UPSTREAM_DISCONNECT} from "./actionTypes";

function backConnect() {
  return {
    type: BACK_CONNECT,
  }
}

function backendDisconnected() {
  return {
    type: BACK_DISCONNECT,
  }
}

function upstreamConnected() {
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

function broken() {
  return {
    type: BROKEN,
  }
}

export default {
  backConnect,
  backendDisconnected,
  upstreamConnected,
  upstreamDisconnect,
  setup,
  broken,
}
