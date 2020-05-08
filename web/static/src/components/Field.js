import React from 'react';
import FieldPlayer from "./FieldPlayer";

import {GameDefinitions} from '../constants';

class Field extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    const items = []

    const fillPlayer = (p, side) => {
      items.push(<FieldPlayer
        key={`${side}-${p.number}`}
        number={p.number}
        team_side={side}
        ang={p.velocity.direction.ang}
        position={p.Position}
      />)
    }

    if (this.props.snapshot.home_team.players) {
      for (const p of this.props.snapshot.home_team.players) {
        if (p) {
          fillPlayer(p, "home")
        }
      }
    }
    if (this.props.snapshot.away_team.players) {
      for (const p of this.props.snapshot.away_team.players) {
        if (p) {
          fillPlayer(p, "away")
        }
      }
    }

    const ball_left = 100 * (this.props.snapshot.ball.Position.X ?? 0) / GameDefinitions.Field.Width
    const ball_bottom = 100 * (this.props.snapshot.ball.Position.Y ?? 0) / GameDefinitions.Field.Height

    const ballStyle = {
      left: `${ball_left}%`,
      bottom: `calc(${ball_bottom}%)`
    }

    return <div id="field">
      <span id="ball" style={ballStyle}/>
      {items}
    </div>;
  }
}

export default Field;