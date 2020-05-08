import React from 'react';

import {GameDefinitions} from '../constants'

class FieldPlayer extends React.Component {
  render() {
    const left = 100 * (this.props.position.X ?? 0) / GameDefinitions.Field.Width
    const bottom = 100 * (this.props.position.Y ?? 0) / GameDefinitions.Field.Height

    const playerStyle = {
      left: `${left}%`,
      bottom: `calc(${bottom}%)`
    }

    const direction = {
      transform: `rotate(${-this.props.ang + 90}deg)`
    }

    return <span
      id={"player-"+this.props.team_side+"-team-1"}
      className={"player player-"+this.props.team_side+"-team"}
      style={playerStyle}

    >
        <span className="player-direction" style={direction} />
        <span className="player-number">{this.props.number}</span>
      </span>
  }
}

export default FieldPlayer;