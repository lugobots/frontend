import React from 'react';

import {Field} from '../constants'

class FieldPlayer extends React.Component {
  render() {
    const left = 100 * (this.props.position.X ?? 0) / Field.Width
    const bottom = 100 * (this.props.position.Y ?? 0) / Field.Height

    const playerStyle = {
      left: `${left}%`,
      bottom: `calc(${bottom}%)`
    }

    return <span
      id={"player-"+this.props.team_side+"-team-1"}
      className={"player player-"+this.props.team_side+"-team"}
      style={playerStyle}

    >
        <span className="player-direction"/>
        <span className="player-number">{this.props.number}</span>
      </span>
  }
}

export default FieldPlayer;