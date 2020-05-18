import React from 'react';
import $ from 'jquery';
import 'jquery-ui/ui/widgets/draggable';


const BackEndPoint = "http://localhost:8080"
const BaseUnit = 100;
const virtualCourtHeight = 100 * BaseUnit;
const virtualCourtWidth = 200 * BaseUnit;

class ToolBarTabDebug extends React.Component {
  constructor(props) {
    super(props);

    this.coordXDOM = React.createRef();
    this.coordYDOM = React.createRef();

    this.sendDebug = this.sendDebug.bind(this);
    this.pauseResume = this.pauseResume.bind(this);
    this.nextTurn = this.nextTurn.bind(this);
    this.nextOrder = this.nextOrder.bind(this);
  }

  sendDebug(path) {
    return fetch(`${BackEndPoint}/remote/${path}`, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }
    }).then(res => res.json())
  }

  pauseResume() {
    this.sendDebug("pause-resume")
      .then(
        (result) => {
          console.log(`DEBG: `, result)
        },
        (error) => {
          console.error(`Debug tool: `, error)
        }
      )
  }

  nextTurn() {
    this.sendDebug("next-turn")
      .then(
        (result) => {
          console.log(`DEBG: `, result)
        },
        (error) => {
          console.error(`Debug tool: `, error)
        }
      )
  }

  nextOrder() {
    this.sendDebug("next-order")
      .then(
        (result) => {
          console.log(`DEBG: `, result)
        },
        (error) => {
          console.error(`Debug tool: `, error)
        }
      )
  }

  componentDidMount() {
    let me = this;
    document.getElementById('field').onmousemove = function (e) {
      const ratio = $(this).innerWidth() / virtualCourtWidth;
      const parentOffset = $(this).offset();
      //or $(this).offset(); if you really just want the current element's offset
      console.log(parentOffset)
      let x = e.pageX - parentOffset.left;
      let y = e.pageY - parentOffset.top;
      x /= ratio;
      y /= ratio;
      y = virtualCourtHeight - y;
      me.coordXDOM.current.innerHTML = x.toFixed(0)
      me.coordYDOM.current.innerHTML = y.toFixed(0)
    };
  }

  render() {
    console.log(`${this.constructor.name} rendered`)
    let deactivateClass = "deactivated"
    let disabled = true
    if (this.props.debugOn) {
      deactivateClass = ""
      disabled = true
    }
    return <div className={`${this.props.className} debug-tab`}>
      <button id="btn-resume" className="btn btn-main" onClick={this.pauseResume}>Resume</button>
      <button id="btn-next-order" aria-disabled={disabled} className={`btn ${deactivateClass}`}
              onClick={this.nextTurn}>Next Order
      </button>
      <button id="btn-next-cycle" aria-disabled={disabled} className={`btn ${deactivateClass}`}
              onClick={this.nextOrder}>Next Cycle
      </button>
      <button id="btn-rearrange" aria-disabled={disabled} className={`btn ${deactivateClass}`}>Rearrange</button>
      <button id="btn-save-positions" className={`btn ${deactivateClass}`}>Save Positions</button>
      <span id="choose-preset">
            <label htmlFor="preset">Choose a pre-set Arrangement</label>
            <select name="preset">
              <option value="default">Choose a pre-set Arrangement</option>
              <option value="preset-1">Pre-set 1</option>
              <option value="preset-2">Pre-set 2</option>
              <option value="preset-3">Pre-set 3</option>
            </select>
          </span>
      <span id="coordenates">
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

export default ToolBarTabDebug;