import React from 'react';
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";

class PanelTeamsInfo extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    return <div id="teams-info">
        <span id="home-team" className="team">
          <span id="profile-home-team" className="profile-team">
            <span id="picture-home-team" className="picture-team">
              <img src={this.props.home_team_avatar}/>
            </span>
            <span id="home-team-name" className="name-team">
              <h3>{this.props.home_team_name}</h3>
            </span>
          </span>
          <span id="flag-home-team" className="flag-team">
            <span id="flag-color-1" className="flag-color"/>
            <span id="flag-color-2" className="flag-color"/>
          </span>

        </span>
      <span id="away-team" className="team">
            <span id="profile-away-team" className="profile-team">
              <span id="picture-away-team" className="picture-team">
                <img src={this.props.away_team_avatar}/>
              </span>
              <span id="name-away-team" className="name-team">
                <h3>{this.props.away_team_name}</h3>
              </span>
            </span>
            <span id="flag-away-team" className="flag-team">
              <span id="flag-color-1" className="flag-color"/>
              <span id="flag-color-2" className="flag-color"/>
            </span>

        </span>
    </div>;
  }
}

PanelTeamsInfo.propTypes = {
  home_team_avatar: PropTypes.string,
  home_team_name: PropTypes.string,
  away_team_avatar: PropTypes.string,
  away_team_name: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    home_team_avatar: state.app.setup.home_team.avatar,
    home_team_name: state.app.setup.home_team.name,
    away_team_avatar: state.app.setup.away_team.avatar,
    away_team_name: state.app.setup.away_team.name,
  }
}

export default connect(mapStateToProps)(PanelTeamsInfo)

