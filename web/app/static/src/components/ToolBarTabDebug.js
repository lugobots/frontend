import React from 'react';
import $ from 'jquery';
import 'jquery-ui/ui/widgets/draggable';

import {getSizeRatio, renderLogger} from "../helpers";
import {GameDefinitions, StadiumStatus} from '../constants'
import stadiumActions from '../redux/stadium/actions'
import PropTypes from "prop-types";
import {connect} from "react-redux";
import channel from "../channel";

class ToolBarTabDebug extends React.Component {
  constructor(props) {
    super(props);

    this.coordXDOM = React.createRef();
    this.coordYDOM = React.createRef();

    this.pauseResume = this.pauseResume.bind(this);
    this.nextTurn = this.nextTurn.bind(this);
    this.nextOrder = this.nextOrder.bind(this);
    this.startRearrange = this.startRearrange.bind(this);
    this.confirmRearranging = this.confirmRearranging.bind(this);
  }


  pauseResume() {
    sendDebug("pause-resume")
      .then(
        ({status, body}) => {
          // console.log(`DEBG: `, status, body)
        },
        (error) => {
          console.error(`Debug tool: `, error)
        }
      )
  }

  nextTurn() {
    sendDebug("next-turn")
      .then(
        ({status, body}) => {
          // console.log(`DEBG: `, status, body)
        },
        (error) => {
          console.error(`Debug tool: `, error)
        }
      )
  }

  nextOrder() {
    sendDebug("next-order")
      .then(
        ({status, body}) => {
          // console.log(`DEBG: `, status, body)
        },
        (error) => {
          console.error(`Debug tool: `, error)
        }
      )
  }

  startRearrange() {
    this.props.dispatch(stadiumActions.pauseForRearrange())
    $('.player').draggable("enable");
  }

  stopRearrange() {
    this.props.dispatch(stadiumActions.pauseForDebug())
  }

  confirmRearranging() {
    let promisses = []
    $(".player").each(function () {
        let ui = $(this);
        let coordsScreen = ui.position();
        let coords = convertPixelToGameUnit(coordsScreen.left, coordsScreen.top);
        const newPosition = {X: Math.round(coords.x), Y: Math.round(coords.y)};

        promisses.push(sendDebug(`players/${this.dataset.team}/${this.dataset.number}`, newPosition, 'PATCH')
          .then(({status, body}) => {
              // console.log(`DEBG: `,status,  body)
              return body.game_snapshot
            },
            (error) => {
              console.error(`Debug tool: `, error)
            }
          ))
      }
    );

    Promise.all(promisses).then((values) => {
      channel.newGameFrame(values.pop())
      this.stopRearrange()
    });
  }

  componentDidMount() {
    let me = this;
    document.getElementById('field').onmousemove = function (e) {
      const ratio = getSizeRatio()
      const parentOffset = $(this).offset();
      //or $(this).offset(); if you really just want the current element's offset
      let x = e.pageX - parentOffset.left;
      let y = e.pageY - parentOffset.top;
      x /= ratio;
      y /= ratio;
      y = GameDefinitions.Field.Height - y;
      me.coordXDOM.current.innerHTML = x.toFixed(0)
      me.coordYDOM.current.innerHTML = y.toFixed(0)
    };
  }

  render() {
    renderLogger(this.constructor.name)
    let enabledPausePlay = true
    let enabledBreakPoint = false
    let enabledRearrange = false
    let enabledSavePos = false
    switch (this.props.stadium_status) {
      case StadiumStatus.DEBUGGING:
        enabledBreakPoint = true
        enabledRearrange = true
        break;
      case StadiumStatus.REARRANGING:
        enabledSavePos = true
        enabledPausePlay = false
        break;
    }
    return <div className={`${this.props.className} debug-tab`}>
      <button id="btn-resume" disabled={!enabledPausePlay} className="btn btn-main"
              onClick={this.pauseResume}>Resume
      </button>
      <button id="btn-next-order" disabled={!enabledBreakPoint} className="btn"
              onClick={this.nextTurn}>Next Order
      </button>
      <button id="btn-next-cycle" disabled={!enabledBreakPoint} className="btn"
              onClick={this.nextOrder}>Next Cycle
      </button>
      <button id="btn-rearrange" disabled={!enabledRearrange} className="btn"
              onClick={this.startRearrange}>Rearrange
      </button>
      <button id="btn-save-positions" disabled={!enabledSavePos} className="btn"
              onClick={this.confirmRearranging}>Save Positions
      </button>
      <span id="choose-preset">
            <label htmlFor="preset">Choose a pre-set Arrangement</label>
            <select name="preset">
              <option value="default">Choose a pre-set Arrangement</option>
              <option value="preset-1">Pre-set 1</option>
              <option value="preset-2">Pre-set 2</option>
              <option value="preset-3">Pre-set 3</option>
            </select>
          </span>
      <span id="coordinates">
            <span id="axis">
              <span id="axis-x" className="axis">X:
                <span className="value-x" ref={this.coordXDOM}>1059.65</span></span>
              <span id="axis-y" className="axis">Y:
                <span className="value-y" ref={this.coordYDOM}>1397.17</span></span>
            </span>
            <span id="icon"/>
          </span>
    </div>;
  }

}

function SetPlayerProperties(side, number, position) {
  sendDebug(`players/${side}/${number}`, {position}, 'PATCH')
    .then(
      ({status, body}) => {
        // console.log(`Rearrange`, status)
      },
      (error) => {
        console.error(`Debug tool: `, error)
      }
    )
}

function sendDebug(path, payload = {}, method = "POST") {
  let status = 0;
  return fetch(`remote/${path}`, {
    method: method,
    body: JSON.stringify(payload),
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  }).then(res => {
    status = res.status
    return res.json()
  }).then(body => {
    return {status, body}
  })
}


function convertPixelToGameUnit(left, top) {
  const ratio = getSizeRatio()
  let xPos = (parseFloat(left) / ratio) //- (GameDefinitions.Player.Size / 2);
  let yPos = (parseFloat(top) / ratio) //- (GameDefinitions.Player.Size / 2);
  yPos = GameDefinitions.Field.Height - yPos;
  return {x: xPos, y: yPos}
}

ToolBarTabDebug.propTypes = {
  stadium_status: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    stadium_status: state.stadium.status,
  }
}

const bar = connect(mapStateToProps)(ToolBarTabDebug)
export {
  bar as ToolBarTabDebug,
  SetPlayerProperties,
}

