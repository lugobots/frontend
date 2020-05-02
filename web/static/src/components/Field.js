import React from 'react';
import FieldPlayer from "./FieldPlayer";
class Field extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    const items = []
    console.log("re render field")
    if (this.props.snapshot.home_team.players) {
      console.log(this.props.snapshot.home_team)
      for (const p of this.props.snapshot.home_team.players) {
        if (p) {
          items.push(<FieldPlayer key={`home-${p.number}`} number={p.number} team_side="home" />)
        }
      }
    }
    if (this.props.snapshot.away_team.players) {
      console.log(this.props.snapshot.away_team)
      for (const p of this.props.snapshot.away_team.players) {
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