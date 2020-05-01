import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";

class Panel extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    return <section id="game-panel">
      <PanelTeamsInfo setup={this.props.setup}/>
      <PanelGameInfo update={this.props.update} time_remaining={this.props.update.time_remaining} />
    </section>;
  }
}

export default Panel;