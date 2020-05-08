import React from 'react';
import Panel from "./Panel";
import Field from "./Field";

import {GameSettings} from '../constants';

const BackEndPoint = "http://localhost:8080"

class Stadium extends React.Component {
  constructor(props) {
    super(props);


    var setup = {
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


    this.initSnapshot = {
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

    this.state = {
      isConnected: false,
      isLoaded: false,
      error: null,
      setup: setup,
      event: {
        type: "",
        time_remaining: "00:00",
        snapshot: this.initSnapshot,
      }
    }
  }

  componentDidMount() {
    const me = this;

    const evtSource = new EventSource(`${BackEndPoint}/game-state/${gameID}/${uuid}/`);

    const eventProcessor = function (event) {
      const g = JSON.parse(event.data);

      // console.log(g.game_event?.game_snapshot)
      me.setState({
        event: {
          time_remaining: g.time_remaining,
          snapshot: g.game_event?.game_snapshot,
        }
      });
    }

    evtSource.addEventListener("state_change", eventProcessor);
    evtSource.addEventListener("new_player", eventProcessor);

    evtSource.addEventListener("ping", function (event) {
      console.log("ping")
    });

    // addEventListener version
    evtSource.addEventListener('open', () => {
      console.log("connection opened")
      me.setState(state => {
        let s = state;
        s.isConnected = true;
        s.event = {
          time_remaining: "",
          snapshot: this.initSnapshot,
        };
        return s;
      })
      me.setup()
    });


    evtSource.onerror = function () {
      console.log("stream connection lost. trying to reconnect...");
      me.setState(state => {
        let s = state;
        s.isConnected = false;
        s.isLoaded = false;
        return s;
      })
    };
  }

  render() {
    const {error, isLoaded, isConnected} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!isConnected) {
      return <div>Connecting...</div>;
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {

      this.setColor('--team-home-color-primary', this.state.setup.home_team.colors.primary);
      this.setColor('--team-home-color-secondary', this.state.setup.home_team.colors.secondary);
      this.setColor('--team-away-color-primary', this.state.setup.away_team.colors.primary);
      this.setColor('--team-away-color-secondary', this.state.setup.away_team.colors.secondary);
      return <div>
        <header id="lugobot-header" className="container">
          <Panel event={this.state.event} setup={this.state.setup}/>
        </header>

        <main id="lugobot-stadium" className="container">
          <Field snapshot={this.state.event.snapshot}/>
        </main>
      </div>;
    }
  }

  setColor(name, colors) {
    const lis = [colors.red ?? 0, colors.green ?? 0, colors.blue ?? 0]
    document.documentElement.style.setProperty(name, lis.toString());
  }

  setup() {
    fetch(`${BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(res => res.json())
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            setup: result
          });
        },
        // Note: it's important to handle errors here
        // instead of a catch() block so that we don't swallow
        // exceptions from actual bugs in components.
        (error) => {
          this.setState({
            isLoaded: true,
            error
          });
        }
      )
  }
}

export default Stadium;

