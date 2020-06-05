import React from 'react';
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";


class PanelGameInfo extends React.Component {
  constructor(props) {
    super(props);
  }


  componentDidMount() {
    /* GAMBIS< APAGAR*/

    document.getElementById('remaining').style.display = "flex"
    document.getElementById('buffering').style.display = "none"
  }

  render() {
    renderLogger(this.constructor.name)
    let homeTeamClass = ""
    if(this.props.team_goal === "home") {
      homeTeamClass = "goal"
    }
    let awayTeamClass = ""
    if(this.props.team_goal === "away") {
      awayTeamClass = "goal"
    }

    return <div id="game-info">
        <span id="score-info">
          <span id="score-home-team" className={`score-team ${homeTeamClass}`}>{this.props.home_score}</span>
          <span id="timer">
            <span id="remaining" className="active">{this.props.time_remaining}</span>
            <span id="shot-clock" >
              <span id="shot-clock-home" />
              <span id="shot-clock-timer">{this.props.shot_time}</span>
              <span id="shot-clock-away" />
            </span>
            <span id="buffering" >
              <span className="label">Buffering</span>
              <span className="percent">50%</span>
              <span className="progress"/>
            </span>
          </span>
          <span id="score-away-team" className={`score-team ${awayTeamClass}`}>{this.props.away_score}</span>
        </span>
    </div>;
  }
}

PanelGameInfo.propTypes = {
  team_goal: PropTypes.string,
  home_score: PropTypes.number,
  away_score: PropTypes.number,
  time_remaining: PropTypes.string,
  shot_time: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    ...state.stadium.panel,
    team_goal: state.stadium.event_data?.team_goal,
  }
}

export default connect(mapStateToProps)(PanelGameInfo)