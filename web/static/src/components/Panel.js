import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";
import {GameDefinitions} from "../constants";
import {ShouldRender} from "../helpers";

class Panel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      time_remaining: "--:--",
      team_goal: "",
      home_score: 0,
      away_score: 0,
    }
  }

  // shouldComponentUpdate(nextProps, nextState) {
  //   return ShouldRender(this.props, nextProps);
  // }

  componentDidMount() {

    this.props.setOnNewEventListener(gameEvent => {

      if(gameEvent.time_remaining !== this.state.time_remaining) {
        this.setState({time_remaining: gameEvent.time_remaining})
      }
      if(gameEvent.team_goal !== this.state.team_goal) {
        this.setState({team_goal: gameEvent.team_goal})
      }

      if(gameEvent.snapshot?.home_team.Score !== this.state.home_score) {
        this.setState({home_score: gameEvent.snapshot?.home_team.Score})
      }
      if(gameEvent.snapshot?.away_team.Score !== this.state.away_score) {
        this.setState({away_score: gameEvent.snapshot?.away_team.Score})
      }
    })
  }

  render() {
    return <section id="game-panel">
      <PanelTeamsInfo setup={this.props.setup}/>
      <PanelGameInfo
        team_goal={this.state.team_goal}
        home_score={this.state.home_score || 0}
        away_score={this.state.away_score || 0}
        time_remaining={this.state.time_remaining}
      />
    </section>;
  }
}

export default Panel;