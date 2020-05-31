import React from 'react';
import Panel from "./Panel";
import Field from "./Field";
import Events from "./Events";
import ToolBar from "./ToolBar";
import {Howl, Howler} from 'howler';
import tickAudio from '../sounds/kick.wav';

import {GameSettings, GameStates, StadiumStates, EventTypes, BackendConfig, ModalModes} from '../constants';
import {renderLogger, updateRatio} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";


class Stadium extends React.Component {
  constructor(props) {
    super(props);
    // this.reset(true)
    // let evtSource = null
    // let lastFrame = null
    // let onNewEventListener = []
    // this.onNewEvent = (snapshot) => {
    //   lastFrame = snapshot
    //   onNewEventListener.forEach(cb => {
    //     cb(snapshot)
    //   })
    // }
    // this.addOnNewEventListener = (cb) => {
    //   if (lastFrame !== null) {
    //     cb(lastFrame)
    //   }
    //   onNewEventListener.push(cb)
    // }
  }

  componentDidMount() {
    // let AAAAA = new Howl({
    //   src: [tickAudio]
    // });
    // AAAAA.play()
    updateRatio()
    // const me = this;
    // let upstreamConnTries = 0;
    // this.evtSource = new EventSource(`${BackendConfig.BackEndPoint}/game-state/${gameID}/${uuid}/`);
    // // addEventListener version
    // this.evtSource.addEventListener('open', () => {
    //   console.log("%cconnection opened", "color: green")
    //   upstreamConnTries = 0;
    //   me.gotoStateSettingUp()
    // });
    // this.evtSource.onerror = () => {
    //   console.error("stream connection lost. trying to reconnect...");
    //   me.gotoStateConnecting()
    // };
    // this.evtSource.addEventListener("ping", () => {
    //   console.debug("ping")
    // });
    // this.evtSource.addEventListener(EventTypes.ConnectionLost, function (event) {
    //   console.error("upstream connection lost")
    //   me.gotoStateConnectingUpstream(upstreamConnTries++)
    // });
    // this.evtSource.addEventListener(EventTypes.ConnectionReestablished, function (event) {
    //   console.log("%cupstream connection reestablished", "color: green")
    //   upstreamConnTries = 0;
    //   me.gotoStateSettingUp()
    // });
    //
    // evtSource.addEventListener(EventTypes.StateChange, (e) => this.processGameEvent(e));
    // evtSource.addEventListener(EventTypes.NewPlayer, (e) => this.processGameEvent(e));
    // evtSource.addEventListener(EventTypes.Breakpoint, (e) => {
    //   const g = JSON.parse(e.data);
    //   console.log("%cDEBUG ON", "color: blue", g)
    //   this.processGameEvent(e)
    // });
    // evtSource.addEventListener(EventTypes.DebugReleased, (e) => {
    //   console.log("%cDEBUG OFF", "color: gray")
    //   this.processGameEvent(e)
    // });
  }

  processGameEvent(event) {
    if(!this.state.isSetup) {
      console.log("%cIgnoring event while not set up", "color: #AA0000")
      return
    }
    const g = JSON.parse(event.data);
    let team_goal = ""
    const s = this.getStadiumStateMode();
    if (s === StadiumStates.StadiumStateGoal) {
      this.gotoStateListening()
    }
    const frameState = g.game_event?.game_snapshot?.state || GameStates.WAITING;
    console.log("event", frameState, g.debug_mode)
    if (frameState === GameStates.GET_READY) {
      const homeScoreOld = this.state.event.snapshot.home_team.Score ?? 0
      const awayScoreOld = this.state.event.snapshot.away_team.Score ?? 0

      const homeScoreNew = g.game_event?.game_snapshot?.home_team?.Score ?? 0
      const awayScoreNew = g.game_event?.game_snapshot?.away_team?.Score ?? 0

      if (homeScoreOld !== homeScoreNew) {
        team_goal = "home";
      } else if (awayScoreOld !== awayScoreNew) {
        team_goal = "away";
      }
    }

    let state = {
      event: {
        debug_mode: g.debug_mode,
        shot_time: g.shot_time,
        time_remaining: g.time_remaining,
        snapshot: g.game_event?.game_snapshot,
      }
    }
    if (team_goal !== "") {
      this.gotoStateGoal(team_goal)
    }
    this.onNewEvent(state.event)
  }

  setMainColor(name, colors) {
    const lis = [colors.red ?? 0, colors.green ?? 0, colors.blue ?? 0]
    document.documentElement.style.setProperty(name, lis.toString());
  }

  setup() {
    fetch(`${BackendConfig.BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(res => res.json())
      .then(
        (result) => {
          console.log("%cSetup", "color: blue")
          console.log(result)
          this.setState({
            isSetup: true,
            setup: result.game_setup,
            upstream_state: result.connection_state === "up",
          });

          this.setMainColor('--team-home-color-primary', this.state.setup.home_team.colors.primary);
          this.setMainColor('--team-home-color-secondary', this.state.setup.home_team.colors.secondary);
          this.setMainColor('--team-away-color-primary', this.state.setup.away_team.colors.primary);
          this.setMainColor('--team-away-color-secondary', this.state.setup.away_team.colors.secondary);
          this.gotoStateListening()
        },
        (error) => {
          this.setState({
            isSetup: false,
            upstream_up: false,
            error
          });
        }
      )
  }

  reset(initialise) {
    document.getElementById('lugobot-view').classList.add("loading");
    document.title = "Lugo"
    const setup = {
      dev_mode: false,
      listening_mode: GameSettings.LISTENING_MODE.TIMER,
      start_mode: GameSettings.START_MODE.WAIT,
      listening_duration: 50,
      game_duration: 6000,
      home_team: {
        name: "Home",
        avatar: "external/profile-team-home.jpg",
        colors: {
          primary: {
            red: 240
          },
          secondary: {
            red: 255,
            green: 255,
            blue: 255
          }
        }
      },
      away_team: {
        name: "Away",
        side: 1,
        avatar: "external/profile-team-away.jpg",
        colors: {
          primary: {
            green: 200
          },
          secondary: {
            green: 240,
            blue: 240
          }
        }
      }
    }

    const snapshot = {
      turn: 0,
      home_team: {
        players: [],
        Score: 0,
      },
      away_team: {
        players: [],
        Score: 0,
      },
      ball: {
        Position: {
          X: 0,
          Y: 0,
        }
      }
    }

    const initialState = {
      v: (new Date()).getTime(),
      isConnected: false,
      isSetup: false,
      upstream_up: false,
      error: null,
      setup: setup,
      modal: null,
      stadium: {
        mode: StadiumStates.StadiumStateConnecting
      },
      event: {
        type: "",
        time_remaining: "00:00",
        shot_time: "00",
        snapshot: snapshot,
      }
    }
    console.log(`SETUP: `, initialState.v)
    if (initialise) {
      this.state = initialState
      return
    }
    console.log("RESETED!")
    this.setState(initialState);
  }

  openModal(title, text, buttons) {
    document.getElementById('lugobot-page').classList.add("active-modal");
    this.setState({
      modal: {title, text, buttons}
    })
  }

  closeModal() {
    document.getElementById('lugobot-page').classList.remove("active-modal");
    this.setState({
      modal: null
    })
  }

  setStadiumState(state) {
    this.resetStadiumState()
    this.closeModal()
    this.setState({
      stadium: state,
      v: (new Date()).getTime(),
    })
    console.log("New state: ", this.state.stadium)
  }

  resetStadiumState() {
    this.setState({
      stadium: {mode: null},
    })
  }

  getStadiumStateMode() {
    return this.state.stadium.mode
  }

  gotoStateConnecting() {
    this.reset()
    this.setState({
      isConnected: false,
      isSetup: false,
    });
    this.setStadiumState({mode: StadiumStates.StadiumStateConnecting})
    this.openModal("Connecting to backend", <span>Wait the connection be established</span>)
  }

  gotoStateSettingUp() {
    document.getElementById('lugobot-view').classList.remove("loading");
    this.setState({
      isConnected: true,
    });
    this.setStadiumState({mode: StadiumStates.StadiumStateSetting})
    this.openModal("Loading game", <span>Loading game state</span>)
    this.setup()
  }

  gotoStateListening() {
    this.setStadiumState({mode: StadiumStates.StadiumStateListening})
  }

  gotoStateConnectingUpstream(tries) {
    this.reset()
    this.setStadiumState({mode: StadiumStates.StadiumStateConn})
    this.openModal("Upstream connection lost",
      <span>The frontend application is not connected to the GameServer.
          <br/>Wait the connection be reestablished <br/><br/>Retrying {tries}</span>)

  }

  gotoStateGoal(team_side) {
    this.setStadiumState({mode: StadiumStates.StadiumStateGoal, side: team_side})
  }

  gotoStateDebugging(action) {
    this.setStadiumState({mode: StadiumStates.StadiumStateDebugging, action})
  }

  render() {
    renderLogger(this.constructor.name)
    this.setMainColor('--team-home-color-primary', this.props.setup.home_team.colors.primary);
    this.setMainColor('--team-home-color-secondary', this.props.setup.home_team.colors.secondary);
    this.setMainColor('--team-away-color-primary', this.props.setup.away_team.colors.primary);
    this.setMainColor('--team-away-color-secondary', this.props.setup.away_team.colors.secondary);

    let  stadium_class = ""
    if(this.props.modal_mode === ModalModes.ALERT) {
      stadium_class = "active-modal"
    }
    return <div id="stadium" className={stadium_class}>
        <Panel />

      {/*<main id="lugobot-stadium" className="container">*/}
      {/*  <Field*/}
      {/*    v={this.state.v}*/}
      {/*    stadium_state={this.state.stadium}*/}
      {/*    setup={this.state.setup}*/}
      {/*    setOnNewEventListener={(cb) => {*/}
      {/*      this.addOnNewEventListener(cb)*/}
      {/*    }}*/}
      {/*  />*/}
      {/*</main>*/}
      {/*<ToolBar*/}
      {/*  v={this.state.v}*/}
      {/*  setup={this.state.setup}*/}
      {/*  stadium_state={this.state.stadium}*/}
      {/*  setOnNewEventListener={(cb) => {*/}
      {/*    this.addOnNewEventListener(cb)*/}
      {/*  }}*/}
      {/*  gotoStateDebugging={this.gotoStateDebugging.bind(this)}*/}
      {/*/>*/}
      <Events />
    </div>;
  }
}

Stadium.propTypes = {
  setup: PropTypes.object,
  modal_mode: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    setup: state.app.setup,
    modal_mode: state.stadium.modal.mode,
  }
}

export default connect(mapStateToProps)(Stadium)

