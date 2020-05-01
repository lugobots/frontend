import React from 'react';


class PanelGameInfo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      snapshot: null,
      time_remaining: "00:00"
    };

  }

  render() {
    const snapshot = this.props.update.snapshot;
    return <div id="game-info">
        <span id="score-info">
          <span id="score-home-team" className="score-team">{snapshot.home_team.score}</span>
          <span id="timer">
            <span id="remaining">{this.props.time_remaining}</span>
            <span id="buffering" >
              <span className="label">Buffering</span>
              <span className="percent">50%</span>
              <span className="progress"/>
            </span>
          </span>
          <span id="score-away-team" className="score-team">{snapshot.home_team.score}</span>
        </span>
    </div>;
  }
}

export default PanelGameInfo;

