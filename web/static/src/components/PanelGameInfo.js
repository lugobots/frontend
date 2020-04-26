import React from 'react';

class PanelGameInfo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      game: null,
    };
  }

  render() {
    return <div id="game-info">
        <span id="score-info">
          <span id="score-home-team" className="score-team">{this.props.game.homeTeam.score}</span>
          <span id="timer">
            <span id="remaining">{this.props.game.timeRemaining}</span>
            <span id="buffering" >
              <span className="label">Buffering</span>
              <span className="percent">50%</span>
              <span className="progress"/>
            </span>
          </span>
          <span id="score-away-team" className="score-team">{this.props.game.awayTeam.score}</span>
        </span>
    </div>;
  }
}

export default PanelGameInfo;

