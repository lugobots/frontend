import React from 'react';

const BackEndPoint = "http://localhost:8080"

class ToolBarTabDebug extends React.Component {
  constructor(props) {
    super(props);

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

  render() {
    return <div className={`${this.props.className} debug-tab`}>
      <button id="btn-resume" className="btn btn-main" onClick={this.pauseResume}>Resume</button>
      <button id="btn-next-order" className="btn" onClick={this.nextTurn}>Next Order</button>
      <button id="btn-next-cycle" className="btn" onClick={this.nextOrder}>Next Cycle</button>
      <button id="btn-rearrange" className="btn">Rearrange</button>
      <button id="btn-save-positions" className="btn">Save Positions</button>
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
              <span id="axis-x" className="axis">X:<span className="value-x">1059.65</span></span>
              <span id="axis-y" className="axis">Y:<span className="value-y">1397.17</span></span>
            </span>
            <span id="icon"/>
          </span>
    </div>;
  }

}

export default ToolBarTabDebug;