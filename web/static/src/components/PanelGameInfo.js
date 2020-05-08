import React from 'react';


class PanelGameInfo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      snapshot: null,
      time_remaining: "00:00"
    };

  }


  componentDidMount() {
    /* GAMBIS< APAGAR*/

    document.getElementById('remaining').style.display = "flex"
    document.getElementById('buffering').style.display = "none"
  }

  render() {
    return <div id="game-info">
        <span id="score-info">
          <span id="score-home-team" className="score-team">{this.props.home_score}</span>
          <span id="timer">
            <span id="remaining" className="active">{this.props.time_remaining}</span>
            <span id="buffering" >
              <span className="label">Buffering</span>
              <span className="percent">50%</span>
              <span className="progress"/>
            </span>
          </span>
          <span id="score-away-team" className="score-team">{this.props.away_score}</span>
        </span>
    </div>;
  }
}

export default PanelGameInfo;

