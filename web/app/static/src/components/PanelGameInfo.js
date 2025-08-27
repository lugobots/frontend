import React from 'react';
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";
import {StadiumStatus} from "../constants";
import audio from "../audio_manager";

let globalPossessionEnding = false;

class PanelGameInfo extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    let homeTeamClass = ""
    let awayTeamClass = ""
    let remainingTimeStyle = {}
    let remainingShotStyle = {}
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
          remainingTimeStyle = {display: 'none'}
          bufferStyle = {}
          bufferProgressStyle = {maxHeight: `${this.props.buffer_percentile}%` }

    }

    // console.log(parseInt(this.props.time_remaining))
    // console.log(this.props.time_remaining)
    if(parseInt(this.props.time_remaining.replace(":", "")) <= 5) {
      remainingTimeStyle = {
        color: 'red',
      }
    }

    const timeToLosePossession = parseInt(this.props.shot_time)

    // we need the condition `timeToLosePossession > 0` because the shot clock may jump to zero during `get ready` state
    if(!globalPossessionEnding && timeToLosePossession > 0 && timeToLosePossession <= 3) {
      globalPossessionEnding = true
      audio.onPossessionEnding()
    }

    if(globalPossessionEnding ) {
      if(timeToLosePossession > 5) {
        globalPossessionEnding = false
      } else {
        remainingShotStyle = {
          color: 'red',
        }
      }
    }
    return <div id="game-info">
        <span id="score-info">
          <span id="score-home-team" className={`score-team ${homeTeamClass}`}>{this.props.home_score}</span>
          <span id="timer">
            <span id="remaining" style={remainingTimeStyle} className="active">{this.props.time_remaining}</span>
            <span id="shot-clock" >
              <span id="shot-clock-label" >Shot clock: </span>
              <span id="shot-clock-timer" style={remainingShotStyle}>{this.props.shot_time}</span>
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
