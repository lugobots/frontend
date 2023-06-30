import React from 'react';
import $ from 'jquery';
import 'jquery-ui/ui/widgets/draggable';

import {GameDefinitions} from '../constants'
import {getSizeRatio, renderLogger} from "../helpers";

// import {SetPlayerProperties} from './ToolBarTabDebug'

class FieldPlayer extends React.Component {

  constructor(props) {
    super(props);
    this.myDOM = React.createRef();
    this.myDirectionDOM = React.createRef();
    this.position = {y: 0, x: 0}
  }

  shouldComponentUpdate(nextProps, nextState) {
    return false
  }

  componentDidMount() {
    const meInJquery = $(this.myDOM.current)
    this.props.setOnNewFrameListener((player, isBallHolder) => {

      if(!player) {
        this.myDOM.current.style.opacity = 0.2
        this.position.x = this.props.defaultPosition.x
        this.position.y = this.props.defaultPosition.y
      } else {
        this.myDOM.current.style.opacity = 1
        this.position.x = player.position.x
        this.position.y = player.position.y
      }

      const left = 100 * this.position.x / GameDefinitions.Field.Width
      const top = 100 * (GameDefinitions.Field.Height - this.position.y) / GameDefinitions.Field.Height

      const hasClass = this.myDOM.current.classList.contains("ball-holder")
      if(!hasClass && isBallHolder) {
        this.myDOM.current.classList.add("ball-holder")
      } else if(this.myDOM.current.classList.contains("ball-holder")) {
        this.myDOM.current.classList.remove("ball-holder")
      }

      this.myDOM.current.style.left = `${left}%`;
      this.myDOM.current.style.top = `${top}%`;
      let ang = 90
      if(player?.velocity?.direction) {
        ang = ((Math.atan2(player?.velocity?.direction.y, player?.velocity?.direction.x)*180)/Math.PI) + 360 - 90
      }
      this.myDirectionDOM.current.style.transform = `rotate(${-ang%360}deg)`;
    })
    meInJquery.draggable({
      containment: "parent",
      stop: () => {
        let left = parseFloat(this.myDOM.current.style.left) / getSizeRatio();
        if(left < 0) {
          left = 0
        }
        let top = parseFloat(this.myDOM.current.style.top) / getSizeRatio();
        if(top > GameDefinitions.Field.Height) {
          top = GameDefinitions.Field.Height
        }
        // const l = (left / GameDefinitions.Field.Width ).toFixed(2)
      }
    });
  }

  render() {
    renderLogger(this.constructor.name)
    let classNames = ['player', `player-${this.props.team_side}-team`];
    // if (this.props.stadium_state.mode === StadiumStates.StadiumStateDebugging && this.props.stadium_state.action === "rearranging") {
    //   classNames.push('rearranging')
    // }

    return <span ref={this.myDOM}
                 id={`player-${this.props.team_side}-team-${this.props.number}`}
                 className={classNames.join(' ')}
                 // rel={this.props.number}
                 data-team={this.props.team_side}
                 data-number={this.props.number}
                 style={{left: 0, top: 0}}
    >
        <span className="player-direction" ref={this.myDirectionDOM} style={{transform: null}}/>
        <span className="player-number">{this.props.number}</span>
      </span>
  }

}

export default FieldPlayer;
