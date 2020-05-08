import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";

class Panel extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    console.log(this.props.event.snapshot)
    return <section id="game-panel">
      <PanelTeamsInfo setup={this.props.setup}/>
      <PanelGameInfo
        home_score={this.props.event.snapshot.home_team.Score ?? 0}
        away_score={this.props.event.snapshot.away_team.Score ?? 0}
        time_remaining={this.props.event.time_remaining}
      />
    </section>;
  }
}

export default Panel;