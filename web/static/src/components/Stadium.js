import React from 'react';
import Panel from "./Panel";
import Field from "./Field";
import Events from "./Events";

import {GameSettings, GameStates, StadiumStates} from '../constants';

const BackEndPoint = "http://localhost:8080"


class Stadium extends React.Component {
  constructor(props) {
    super(props);
    this.reset(true)
    let lastFrame = null
    let onNewEventListener = []
    this.onNewEvent = (snapshot) => {
      lastFrame = snapshot
      onNewEventListener.forEach(cb => {
        cb(snapshot)
      })
    }
    this.addOnNewEventListener = (cb) => {
      if (lastFrame !== null) {
        cb(lastFrame)
      }
      onNewEventListener.push(cb)
    }
  }

  componentDidMount() {
    console.log(`${this.constructor.name} mounted`)
    const me = this;
    let upstreamConnTries = 0;
    const evtSource = new EventSource(`${BackEndPoint}/game-state/${gameID}/${uuid}/`);
    // addEventListener version
    evtSource.addEventListener('open', () => {
      console.log("%cconnection opened", "color: green")
      upstreamConnTries = 0;
      me.gotoStateSettingUp()
    });
    evtSource.onerror = () => {
      console.error("stream connection lost. trying to reconnect...");
      me.gotoStateConnecting()
    };
    evtSource.addEventListener("connection_lots", function (event) {
      console.error("upstream connection lost")
      me.gotoStateConnectingUpstream(upstreamConnTries++)
    });
    evtSource.addEventListener("connection_Reestablished", function (event) {
      console.log("%cupstream connection reestablished", "color: green")
      upstreamConnTries = 0;
      me.gotoStateSettingUp()
    });

    evtSource.addEventListener("state_change", (e) => this.processGameEvent(e));
    evtSource.addEventListener("new_player", (e) => this.processGameEvent(e));
    evtSource.addEventListener("ping", () => {
      console.debug("ping")
    });
  }

  processGameEvent(event) {
    const g = JSON.parse(event.data);
    let team_goal = ""
    const s = this.getStadiumStateMode();
    if (s === StadiumStates.StadiumStateGoal) {
      this.gotoStateListening()
    }
    const frameState = g.game_event?.game_snapshot?.state;
    console.log("event", frameState, g.shot_time)
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
        shot_time: g.shot_time,
        time_remaining: g.time_remaining,
        snapshot: g.game_event?.game_snapshot,
      }
    }
    this.onNewEvent(state.event)
    if (team_goal !== "") {
      this.gotoStateGoal(team_goal)
    }
  }

  setMainColor(name, colors) {
    const lis = [colors.red ?? 0, colors.green ?? 0, colors.blue ?? 0]
    document.documentElement.style.setProperty(name, lis.toString());
  }

  setup() {
    fetch(`${BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(res => res.json())
      .then(
        (result) => {
          console.log(`Setup: `, result)
          this.setState({
            isLoaded: true,
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
            isLoaded: true,
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
      isLoaded: false,
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
  }

  resetStadiumState() {
    this.setState({
      stadium: {mode: null},
    })
    console.log("New state: ", this.state.stadium)
  }

  getStadiumStateMode() {
    return this.state.stadium.mode
  }

  gotoStateConnecting() {
    this.reset()
    this.setState({
      isConnected: false,
      isLoaded: false,
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

  render() {
    return <div>
      <header id="lugobot-header"
              className={`container ${
                this.getStadiumStateMode() === StadiumStates.StadiumStateGoal ? 'active-modal' : ''}`}>
        <Panel
          v={this.state.v}
          stadium_state={this.state.stadium}
          setup={this.state.setup}
          setOnNewEventListener={(cb) => {
            this.addOnNewEventListener(cb)
          }}
        />
      </header>

      <main id="lugobot-stadium" className="container">
        <Field
          v={this.state.v}
          setup={this.state.setup}
          setOnNewEventListener={(cb) => {
            this.addOnNewEventListener(cb)
          }}
        />
      </main>
      <Events
        v={this.state.v}
        stadium_state={this.state.stadium}
        modal={this.state.modal}
        setOnNewEventListener={(cb) => {
          this.addOnNewEventListener(cb)
        }}
      />
    </div>;
  }
}

export default Stadium;

