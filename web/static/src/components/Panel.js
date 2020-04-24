import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";

class Panel extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    return <section id="game-panel">
        <PanelTeamsInfo game={this.props.game} />
    </section>;
  }
}

export default Panel;