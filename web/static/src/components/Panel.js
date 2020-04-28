import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";

class Panel extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    return <section id="game-panel">
      <PanelTeamsInfo game={this.props.game} setup={this.props.setup}/>
      <PanelGameInfo game={this.props.game} setup={this.props.setup}/>
    </section>;
  }
}

export default Panel;