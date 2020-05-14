import React from 'react';
import Panel from "./Panel";
import Field from "./Field";
import Events from "./Events";

import {GameSettings, GameStates} from '../constants';

const BackEndPoint = "http://localhost:8080"

class Stadium extends React.Component {
  constructor(props) {
    super(props);
    this.reset(true)

    let lastframe = null
    let onNewEventListener = []
    this.onNewEvent = (snapshot) => {
      lastframe = snapshot
      onNewEventListener.forEach(cb => {
        cb(snapshot)
      })
    }
    this.addOnNewEventListener = (cb) =>{
      if(lastframe !== null) {
        cb(lastframe)
      }
      onNewEventListener.push(cb)
    }
  }

  componentDidMount() {
    console.log(`${this.constructor.name} mounted`)
    const me = this;
    let upstreamConnTries = 0;
    const evtSource = new EventSource(`${BackEndPoint}/game-state/${gameID}/${uuid}/`);

    const eventProcessor = function (event) {
      const g = JSON.parse(event.data);
      let team_goal = ""
      console.log(`event`)
      const newState = g.game_event?.game_snapshot?.state || GameStates.WAITING
      if (g.game_event?.game_snapshot?.state === GameStates.GET_READY) {

        const homeScoreOld = me.state.event.snapshot.home_team.Score ?? 0
        const awayScoreOld = me.state.event.snapshot.away_team.Score ?? 0

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
          team_goal: team_goal,
          time_remaining: g.time_remaining,
          snapshot: g.game_event?.game_snapshot,
        }
      }
      me.onNewEvent(state.event)
    }

    evtSource.addEventListener("state_change", eventProcessor);
    evtSource.addEventListener("new_player", eventProcessor);
    evtSource.addEventListener("ping", function (event) {
      console.debug("ping")
    });
    evtSource.addEventListener("connection_lots", function (event) {
      console.error("upstream connection lost")
      upstreamConnTries++
      me.openModal("Upstream connection lost",
        <span>The frontend application is not connected to the GameServer.
          <br/>Wait the connection be reestablished <br/><br/>Retrying {upstreamConnTries}</span>)
    });
    evtSource.addEventListener("connection_Reestablished", function (event) {
      console.log("%cupstream connection reestablished", "color: green")
      upstreamConnTries = 0;
      me.closeModal()
      me.setup()
    });

    // addEventListener version
    evtSource.addEventListener('open', () => {
      console.log("%cconnection opened", "color: green")
      me.reset()
      upstreamConnTries = 0;
      document.getElementById('lugobot-view').classList.remove("loading");
      this.setState({
        isConnected: true,
      });
      me.setup()
    });

    evtSource.onerror = () => {
      console.error("stream connection lost. trying to reconnect...");
      this.setState({
        isConnected: false,
        isLoaded: false,
      });
      me.reset()
    };


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
            v: (new Date()).getUTCDate(),
            isLoaded: true,
            setup: result.game_setup,
            upstream_state: result.connection_state === "up",
          });
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
      v: (new Date()).getUTCDate(),
      isConnected: false,
      isLoaded: false,
      upstream_up: false,
      error: null,
      setup: setup,
      modal: null,
      event: {
        team_goal: "",
        type: "",
        time_remaining: "00:00",
        snapshot: snapshot,
      }
    }

    if (initialise) {
      this.state = initialState
      return
    }
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

  setGoalState(team_side) {

  }

  render() {
    console.log(`${this.constructor.name} rendered`)
    const {error, isLoaded, isConnected} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!isConnected) {
      return <div>Connecting...</div>;
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {

      let headerGoalClass = ""
      if (this.state.event.team_goal !== "") {
        headerGoalClass = `active-modal`
      }

      this.setMainColor('--team-home-color-primary', this.state.setup.home_team.colors.primary);
      this.setMainColor('--team-home-color-secondary', this.state.setup.home_team.colors.secondary);
      this.setMainColor('--team-away-color-primary', this.state.setup.away_team.colors.primary);
      this.setMainColor('--team-away-color-secondary', this.state.setup.away_team.colors.secondary);
      return <div>
        <header id="lugobot-header" className={`container ${headerGoalClass}`}>
          <Panel
            v={this.state.v}
            setup={this.state.setup}
            setOnNewEventListener={ (cb) => { this.addOnNewEventListener(cb)}}
          />
        </header>

        <main id="lugobot-stadium" className="container">
          <Field
            v={this.state.v}
            setup={this.state.setup}
            setOnNewEventListener={ (cb) => { this.addOnNewEventListener(cb)}}
          />
        </main>
        <Events
          v={this.state.v}
          modal={this.state.modal}
          setOnNewEventListener={ (cb) => { this.addOnNewEventListener(cb)}}
        />
      </div>;
    }
  }
}

export default Stadium;

