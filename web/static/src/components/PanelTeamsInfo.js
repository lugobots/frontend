import React from 'react';

class PanelTeamsInfo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      game: null,
      setup: null,
    };
  }

  render() {
    return <div id="teams-info">
        <span id="home-team" className="team">
          <span id="profile-home-team" className="profile-team">
            <span id="picture-home-team" className="picture-team">
              <img src={this.props.setup.home_team.avatar} alt="Team Home logo"/>
            </span>
            <span id="home-team-name" className="name-team">
              <h3>{this.props.setup.home_team.name}</h3>
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
                <img src={this.props.setup.away_team.avatar} alt="Team Home logo"/>
              </span>
              <span id="name-away-team" className="name-team">
                <h3>{this.props.setup.away_team.name}</h3>
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

export default PanelTeamsInfo;

