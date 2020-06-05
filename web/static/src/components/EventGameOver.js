import React from 'react';
import PropTypes from "prop-types";
import {connect} from "react-redux";
import trophy from '../img/bg/modals/icon-trophy.png'

class EventGameOver extends React.Component {
  render() {
    if(!this.props.show) {
      return <span />
    }
    return <div id="modal-winner" className={`modal modal-bg active-modal zoom-In`}>
      <span className="close-modal" title="Close modal"><i className="fas fa-times"/></span>
      <span className="modal-content">
        <img id="icon-trophy" className="icon-modal" src={trophy} alt="Icon trophy" />
        <h2 className="modal-title">Final Scoreboard</h2>
      </span>

      <section className="game-panel">
        <div id="teams-info">
            <span id="home-team" className="team">
              <span id="profile-home-team" className="profile-team">
                <span id="picture-home-team" className="picture-team">
                  <img src={this.props.home_team_avatar}  />
                </span>
              </span>
            </span>
          <span id="away-team" className="team">
                <span id="profile-away-team" className="profile-team">
                  <span id="picture-away-team" className="picture-team">
                    <img src={this.props.away_team_avatar}  />
                  </span>
                </span>
          </span>
        </div>
        <div id="modal-game-info">
            <span id="score-info">
              <span id="score-home-team" className="score-team">{this.props.home_score}</span>
              <span id="score-away-team" className="score-team">{this.props.away_score}</span>
            </span>
        </div>
      </section>
    </div>
  ;
  }
}

EventGameOver.propTypes = {
  home_score: PropTypes.number,
  away_score: PropTypes.number,
  home_team_avatar: PropTypes.string,
  away_team_avatar: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    home_score: state.stadium.panel.home_score,
    away_score: state.stadium.panel.away_score,
    home_team_avatar: state.app.setup.home_team.avatar,
    away_team_avatar: state.app.setup.away_team.avatar,
  }
}


export default connect(mapStateToProps)(EventGameOver)