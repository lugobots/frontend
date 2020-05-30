import React from 'react';
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import appAction from "./redux/app/actions";
import matchAction from "./redux/match/actions";
import {AppStatus, BackendConfig, EventTypes} from "./constants";
import store from "./store";
import Stadium from "./components/Stadium";

class App extends React.Component {

  setup() {
    fetch(`${BackendConfig.BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(res => res.json())
      .then(
        (result) => {
          // console.log("%cSetup", "color: blue")
          // console.log(result)
          this.props.dispatch(appAction.setup(result))
        },
        (error) => {
          // @todo handle the error
          alert("oops... something went wrong")
        }
      )


  }

  onStateChange(event) {
    const g = JSON.parse(event.data);
    this.updateScoreBoard(g)
  }



  componentDidMount() {

    this.evtSource = new EventSource(`${BackendConfig.BackEndPoint}/game-state/${gameID}/${uuid}/`);
    // addEventListener version
    this.evtSource.addEventListener('open', () => {
      this.props.dispatch(appAction.backConnect())
    });
    this.evtSource.onerror = () => {
      this.props.dispatch(appAction.backDisconnect())
    };

    this.evtSource.addEventListener("ping", () => {
      console.debug("ping")
    });

    this.evtSource.addEventListener(EventTypes.ConnectionLost, () => {
      console.error("%cupstream connection lost", "color: #AA0000")
      this.props.dispatch(appAction.upstreamDisconnect())
    });
    this.evtSource.addEventListener(EventTypes.ConnectionReestablished, () => {
      console.log("%cupstream connection reestablished", "color: green")
      this.props.dispatch(appAction.upstreamConnect())
    });
    this.evtSource.addEventListener(EventTypes.StateChange, (e) => this.onStateChange(e));
    this.evtSource.addEventListener(EventTypes.Goal, (e) => {
      const g = JSON.parse(e.data);
      this.props.dispatch(matchAction.displayGoal(g.game_event.goal.side.toLowerCase()))
    });
  }

  render() {
    if (this.props.status === AppStatus.Setting) {
      this.setup()
    }
    return <Stadium/>;
  }

  updateScoreBoard(g) {
    if (
      g.time_remaining !== store.getState().match.panel.time_remaining
      ||
      g.shot_time !== store.getState().match.panel.shot_time
    ) {
      this.props.dispatch(matchAction.updatePanel({
        time_remaining: g.time_remaining,
        shot_time: g.shot_time,
        home_score: g.game_event?.game_snapshot?.home_team?.Score ?? 0,
        away_score: g.game_event?.game_snapshot?.away_team?.Score ?? 0,
        team_goal: "",
      }))
    }
  }
}

App.propTypes = {
  dispatch: PropTypes.func.isRequired,
  status: PropTypes.string.isRequired,
}

const mapStateToProps = state => {

  return {
    status: state.app.status
  }
}

export default connect(mapStateToProps)(App)