import React from 'react';
import $ from 'jquery';
import 'jquery-ui/ui/widgets/draggable';

import {GameDefinitions} from '../constants'
import {ShouldRender, getSizeRatio} from "../helpers";

import {SetPlayerProperties} from './ToolBarTabDebug'

class FieldPlayer extends React.Component {

  constructor(props) {
    super(props);
    this.myDOM = React.createRef();
    this.myDirectionDOM = React.createRef();
    this.position = {y: 0, x: 0}
  }

  shouldComponentUpdate(nextProps, nextState) {
    $(this.myDOM.current).draggable("disable");
    if (nextProps.stadium_state.mode === StadiumStates.StadiumStateDebugging && nextProps.stadium_state.action === "rearranging") {
      $(this.myDOM.current).draggable("enable");
    }
    return ShouldRender(this.props, nextProps);
  }

  componentDidMount() {
    const meInJquery = $(this.myDOM.current)
    this.props.setOnNewFrameListener(player => {
      this.position.x = player.position.x ?? 0
      this.position.y = player.position.y ?? 0
      const left = 100 * this.position.x / GameDefinitions.Field.Width
      const top = 100 * (GameDefinitions.Field.Height - this.position.y) / GameDefinitions.Field.Height

      this.myDOM.current.style.left = `${left}%`;
      this.myDOM.current.style.top = `${top}%`;
      this.myDirectionDOM.current.style.transform = `rotate(${-player.velocity.direction.ang + 90}deg)`;
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
        const l = (left / GameDefinitions.Field.Width ).toFixed(2)
        const t = (top / GameDefinitions.Field.Height).toFixed(2)

        this.myDOM.current.style.left = `${100 * l}%`;
        this.myDOM.current.style.top = `${100 * t}%`;
        console.log(GameDefinitions.Field.Height, top)
        SetPlayerProperties(this.props.team_side, this.props.number, {
          x: left,
          y: GameDefinitions.Field.Height - top,
        })
      }
    });
  }

  render() {
    let classNames = ['player', `player-${this.props.team_side}-team`];
    if (this.props.stadium_state.mode === StadiumStates.StadiumStateDebugging && this.props.stadium_state.action === "rearranging") {
      classNames.push('rearranging')
    }
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