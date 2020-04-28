import React from 'react';
import FieldPlayer from "./FieldPlayer";
class Field extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    const items = []
    if (this.props.game.home_team.players) {
      for (const p of this.props.game.home_team.players) {
        if (p) {
          items.push(<FieldPlayer key={`home-${p.number}`} number={p.number} team_side="home" />)
        }
      }
    }
    if (this.props.game.away_team.players) {
      for (const p of this.props.game.away_team.players) {
        if (p) {
          items.push(<FieldPlayer key={`away-${p.number}`} number={p.number} team_side="away" />)
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