import React from 'react';
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";
import {StadiumStatus} from "../constants";


class PanelGameInfo extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    let homeTeamClass = ""
    let awayTeamClass = ""
    let remainingStyle = {}
    let bufferProgressStyle = {}
    let bufferStyle = {display: 'none'}
    switch (this.props.stadium_status) {
      case StadiumStatus.GOAL:
        if(this.props.team_goal === "home") {
          homeTeamClass = "goal"
        }
        if(this.props.team_goal === "away") {
          awayTeamClass = "goal"
        }
        break;
        case StadiumStatus.BUFFERING:
          remainingStyle = {display: 'none'}
          bufferStyle = {}
          bufferProgressStyle = {maxHeight: `${this.props.buffer_percentile}%` }

    }
    return <div id="game-info">
        <span id="score-info">
          <span id="score-home-team" className={`score-team ${homeTeamClass}`}>{this.props.home_score}</span>
          <span id="timer">
            <span id="remaining" style={remainingStyle} className="active">{this.props.time_remaining}</span>
            <span id="shot-clock" style={remainingStyle}>
              <span id="shot-clock-home" />
              <span id="shot-clock-timer">{this.props.shot_time}</span>
              <span id="shot-clock-away" />
            </span>
            <span id="buffering" style={bufferStyle} >
              <span className="label">Buffering</span>
              <span className="percent">{this.props.buffer_percentile}%</span>
              <span className="progress" style={bufferProgressStyle}/>
            </span>
          </span>
          <span id="score-away-team" className={`score-team ${awayTeamClass}`}>{this.props.away_score}</span>
        </span>
    </div>;
  }
}

PanelGameInfo.propTypes = {
  stadium_status: PropTypes.string,
  team_goal: PropTypes.string,
  home_score: PropTypes.number,
  away_score: PropTypes.number,
  time_remaining: PropTypes.string,
  shot_time: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    ...state.stadium.panel,
    stadium_status: state.stadium.status,
    team_goal: state.stadium.event_data?.team_goal,
    buffer_percentile: state.stadium.event_data?.percentile,
  }
}

export default connect(mapStateToProps)(PanelGameInfo)