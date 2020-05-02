import React from 'react';
import FieldPlayer from "./FieldPlayer";
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
        position={p.Position}
      />)
    }
    console.log(this.props.snapshot)
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

    return <div id="field">
      <span id="ball"/>
      {items}
    </div>;
  }
}

export default Field;