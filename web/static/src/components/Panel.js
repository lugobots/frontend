import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";

class Panel extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    return <section id="game-panel">
        <PanelTeamsInfo game={this.props.game} />
        <PanelGameInfo game={this.props.game} />
    </section>;
  }
}

export default Panel;