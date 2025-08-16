import React from 'react';
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import appAction from "./redux/app/actions";
import stadiumAction from "./redux/stadium/actions";
import {AppStatus, EventTypes, GameStates, StadiumStatus} from "./constants";
import store from "./store";
import Stadium from "./components/Stadium";
import channel from "./channel";
import audio from "./audio_manager";
import {OVERTIME} from "./redux/stadium/actionTypes";
// import {Howl} from "howler";
// import audioKick from "./sounds/kick.mp3";
// import audioNewPlayer from "./sounds/new-player.wav";
// import audioRefereeStart from "./sounds/referee-whistle.mp3";
// import audioBackground from "./sounds/gb.wav";
// import audioPublic from "./sounds/public.mp3";
// import goal from "./sounds/goal1.wav";
// import audioConnectionLost from "./sounds/connection-lost.mp3";
// import audioReconnected from "./sounds/reconnected.mp3";

class App extends React.Component {
  constructor(props) {
    super(props);

    this.audioManager = audio

    this.ballOnHold = false;
    this.previousGameState = GameStates.WAITING;
  }


  setup() {
    let status = false
    fetch(`setup/${gameID}/${uuid}`)
      .then(response => {
        status = response.status
        return response.json()
      })
      .then(result => {
          if (status !== 200) {
            throw new Error(result.error)
          }
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

    // if (data.game_event.game_snapshot?.state !== GameStates.PLAYING && data.game_event.game_snapshot?.state !== "LISTENING") {
    //   console.log(data.game_event.game_snapshot?.state)
    // }

    // detecting game start or restart
    if (data.game_event.game_snapshot?.state === GameStates.LISTENING && this.previousGameState === GameStates.GET_READY) {
      this.audioManager.onGameRestart()
      if(data.game_event.game_snapshot?.period === "OVERTIME") {
        console.log(`changed the period`, data.game_event.game_snapshot?.period)
        this.props.dispatch(stadiumAction.overtime())
        this.audioManager.onOvertime()
      }
      // } else if (data.game_event.game_snapshot?.state !== GameStates.WAITING && this.previousGameState === GameStates.WAITING) {
      //   this.audioManager.onGameStarts()
    }
    if (data.game_event.game_snapshot?.state === GameStates.LISTENING && !this.audioManager.isAmbienceOn()) {
      this.audioManager.onGameResume()
    }


    if (data.game_event.game_snapshot?.ball?.holder?.number === 1) {
      // this.audio.audioPublic.play("claps_good_try")
    }

    // detecting ball kick
    if (!data.game_event.game_snapshot?.ball?.holder) {
      if (this.ballOnHold) {
        this.audioManager.onKick()

        this.ballOnHold = false
      }
    } else {
      this.ballOnHold = true
    }


    this.previousGameState = data.game_event.game_snapshot?.state
    this.updateTimer(data)
    channel.newGameFrame(data.game_event.game_snapshot)
  }

  parse(event) {
    return JSON.parse(event.data);
  }

  componentDidMount() {
    let upstreamConnTries = 0;
    let backConnTries = 0;
    let gameIsOver = false;
    this.evtSource = new EventSource(`game-state/${gameID}/${uuid}/`);
    this.evtSource.onerror = () => {
      if (backConnTries === 0) {
        this.audioManager.onBackendConnectionLost()
      }
      backConnTries++
      this.props.dispatch(appAction.backendDisconnected())
      this.props.dispatch(stadiumAction.displayAlert("Connecting to backend",
        <span>Wait the connection be established<br/><br/>Retrying {backConnTries}</span>))
    };
    this.evtSource.addEventListener('open', () => {
      if (backConnTries > 0) {
        this.audioManager.onBackendReconnected()
      }
      upstreamConnTries = 0;
      backConnTries = 0;
      gameIsOver = false
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
      if (upstreamConnTries === 0) {
        this.audioManager.onUpstreamConnectionLost()
      }

      upstreamConnTries++
      this.props.dispatch(appAction.upstreamDisconnect())
      this.props.dispatch(stadiumAction.displayAlert("Upstream connection lost",
        <span>The frontend application is not connected to the GameServer.
           <br/>Wait the connection be reestablished <br/><br/>Retrying {upstreamConnTries}</span>))
    });
    this.evtSource.addEventListener(EventTypes.ConnectionReestablished, () => {
      upstreamConnTries = 0;
      this.audioManager.onUpstreamReconnected()
      console.log("%cupstream connection reestablished", "color: green")
      this.props.dispatch(appAction.upstreamConnected())
    });

    this.evtSource.addEventListener(EventTypes.StateChange, (e) => this.onStateChange(this.parse(e)));
    this.evtSource.addEventListener(EventTypes.Goal, (e) => {
      this.audioManager.onGoal()
      const g = this.parse(e);
      this.updatePanel(g)
      //ignore celebrations on dev mode.
      if (!store.getState().app.setup.dev_mode) {
        this.props.dispatch(stadiumAction.displayGoal(g.game_event.goal.side.toLowerCase()))
      }
    });
    this.evtSource.addEventListener(EventTypes.GameOver, (e) => {
      gameIsOver = true
      const g = this.parse(e);
      this.updatePanel(g)
      this.audioManager.onGameOver()
      this.props.dispatch(stadiumAction.over(g.game_event?.game_over?.reason))
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

    this.evtSource.addEventListener(EventTypes.NewPlayer, (e) => {
      this.onStateChange(this.parse(e))
      this.audioManager.onNewPlayer()
    });
    this.evtSource.addEventListener(EventTypes.LostPlayer, (e) => {
      if(!gameIsOver) {
        this.onStateChange(this.parse(e))
        this.audioManager.onLostPlayer()
      }

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
