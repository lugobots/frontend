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
    let status = false
    fetch(`${BackendConfig.BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(response => {
        status = response.status
        return response.json()
      })
      .then(result => {
          if (status !== 200) {
            throw new Error(result.error)
          }
          console.log(result)
          this.props.dispatch(appAction.setup(result))
        }
      ).catch(e => {
      this.props.dispatch(appAction.broken())
      this.props.dispatch(stadiumAction.displayAlert("Error getting game configuration",
        <span>Error: {e.message}</span>
      ))
    })


  }

  onStateChange(data) {
    const s = store.getState().stadium.status
    const blockingStatus = [
      StadiumStatus.ALERT,
      StadiumStatus.PLAYING,
      StadiumStatus.DEBUGGING,
      StadiumStatus.REARRANGING,
    ]
    if (!blockingStatus.includes(s)) {
      this.props.dispatch(stadiumAction.resume())
    }
    this.updateTimer(data)
    channel.newGameFrame(data.game_event.game_snapshot)
  }

  parse(event) {
    return JSON.parse(event.data);
  }

  componentDidMount() {
    let upstreamConnTries = 0;
    let backConnTries = 0;
    this.evtSource = new EventSource(`${BackendConfig.BackEndPoint}/game-state/${gameID}/${uuid}/`);
    this.evtSource.onerror = () => {
      backConnTries++
      this.props.dispatch(appAction.backDisconnect())
      this.props.dispatch(stadiumAction.displayAlert("Connecting to backend",
        <span>Wait the connection be established<br/><br/>Retrying {backConnTries}</span>))
    };
    this.evtSource.addEventListener('open', () => {
      upstreamConnTries = 0;
      backConnTries = 0;
      this.props.dispatch(appAction.backConnect())
    });
    this.evtSource.addEventListener("ping", () => {
      console.debug("ping")
    });
    this.evtSource.addEventListener(EventTypes.ConnectionLost, () => {
      if (store.getState().stadium.status === StadiumStatus.OVER) {
        return
      }
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
    this.evtSource.addEventListener(EventTypes.StateChange, (e) => this.onStateChange(this.parse(e)));
    this.evtSource.addEventListener(EventTypes.Goal, (e) => {
      const g = this.parse(e);
      this.updatePanel(g)
      //ignore celebrations on dev mode.
      if (!store.getState().app.setup.dev_mode) {
        this.props.dispatch(stadiumAction.displayGoal(g.game_event.goal.side.toLowerCase()))
      }
    });
    this.evtSource.addEventListener(EventTypes.GameOver, (e) => {
      const g = this.parse(e);
      this.updatePanel(g)
      this.props.dispatch(stadiumAction.over())
    });
    this.evtSource.addEventListener(EventTypes.Breakpoint, (e) => {
      this.props.dispatch(stadiumAction.pauseForDebug());
      this.onStateChange(this.parse(e))
    });
    this.evtSource.addEventListener(EventTypes.DebugReleased, () => {
      this.props.dispatch(stadiumAction.resume())
    });
    this.evtSource.addEventListener(EventTypes.Buffering, (e) => {
      this.props.dispatch(stadiumAction.buffering(this.parse(e).buffer_percentile))
    });
    this.evtSource.addEventListener(EventTypes.BufferReady, () => {
      this.props.dispatch(stadiumAction.resume())
    });
  }

  render() {
    if (this.props.status === AppStatus.Setting) {
      this.setup()
      this.props.dispatch(stadiumAction.reset())
    }
    return <Stadium/>;
  }

  updatePanel(g) {
    this.props.dispatch(stadiumAction.updatePanel({
      time_remaining: g.time_remaining,
      shot_time: g.shot_time,
      home_score: g.game_event?.game_snapshot?.home_team?.score ?? 0,
      away_score: g.game_event?.game_snapshot?.away_team?.score ?? 0,
    }))
  }

  updateTimer(g) {
    if (
      g.time_remaining !== store.getState().stadium.panel.time_remaining
      ||
      g.shot_time !== store.getState().stadium.panel.shot_time
    ) {
      this.updatePanel(g)
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