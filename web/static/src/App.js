import React from 'react';
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import appAction from "./redux/app/actions";
import stadiumAction from "./redux/stadium/actions";
import {AppStatus, BackendConfig, EventTypes, StadiumStatus} from "./constants";
import store from "./store";
import Stadium from "./components/Stadium";
import channel from "./channel";

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
    const s = store.getState().stadium.status
    if(s !== StadiumStatus.ALERT && s !== StadiumStatus.PLAYING) {
      this.props.dispatch(stadiumAction.resume())
    }
    this.updateScoreBoard(g)
    channel.newGameEvent(g)
  }



  componentDidMount() {
    let upstreamConnTries = 0;
    let backConnTries = 0;
    this.evtSource = new EventSource(`${BackendConfig.BackEndPoint}/game-state/${gameID}/${uuid}/`);
    // addEventListener version
    this.evtSource.addEventListener('open', () => {
      upstreamConnTries = 0;
      backConnTries = 0;
      this.props.dispatch(appAction.backConnect())
    });
    this.evtSource.onerror = () => {
      backConnTries++
      this.props.dispatch(appAction.backDisconnect())
      this.props.dispatch(stadiumAction.displayAlert("Connecting to backend",
        <span>Wait the connection be established<br/><br/>Retrying {backConnTries}</span>))
    };

    this.evtSource.addEventListener("ping", () => {
      console.debug("ping")
    });

    this.evtSource.addEventListener(EventTypes.ConnectionLost, () => {
      console.error("%cupstream connection lost", "color: #AA0000")
      upstreamConnTries++
      this.props.dispatch(appAction.upstreamDisconnect())
      this.props.dispatch(stadiumAction.displayAlert("Upstream connection lost",
        <span>The frontend application is not connected to the GameServer.
          <br/>Wait the connection be reestablished <br/><br/>Retrying {upstreamConnTries}</span>))
    });
    this.evtSource.addEventListener(EventTypes.ConnectionReestablished, () => {
      upstreamConnTries = 0;
      console.log("%cupstream connection reestablished", "color: green")
      this.props.dispatch(appAction.upstreamConnect())
    });
    this.evtSource.addEventListener(EventTypes.StateChange, (e) => this.onStateChange(e));
    this.evtSource.addEventListener(EventTypes.Goal, (e) => {
      const g = JSON.parse(e.data);
      console.log(g)
      this.props.dispatch(stadiumAction.displayGoal(g.game_event.goal.side.toLowerCase()))
    });
  }

  render() {
    if (this.props.status === AppStatus.Setting) {
      this.setup()
      this.props.dispatch(stadiumAction.resume())
    }
    return <Stadium/>;
  }

  updateScoreBoard(g) {
    if (
      g.time_remaining !== store.getState().stadium.panel.time_remaining
      ||
      g.shot_time !== store.getState().stadium.panel.shot_time
    ) {
      this.props.dispatch(stadiumAction.updatePanel({
        time_remaining: g.time_remaining,
        shot_time: g.shot_time,
        home_score: g.game_event?.game_snapshot?.home_team?.score ?? 0,
        away_score: g.game_event?.game_snapshot?.away_team?.score ?? 0,
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
    status: state.app.status,
  }
}

export default connect(mapStateToProps)(App)