import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";
import {renderLogger} from "../helpers";

class Panel extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    return <header id="lugobot-header"
                   className={`container `}
      // ${/*this.getStadiumStateMode() === StadiumStates.StadiumStateGoal ? 'active-modal' : ''*/}`
    >
      <section id="game-panel">
        <PanelTeamsInfo/>
        <PanelGameInfo/>
      </section>
    </header>;
  }
}

export default Panel;