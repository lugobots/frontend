import React from 'react';
import Panel from "./Panel";
import Field from "./Field";
import Events from "./Events";

import {GameSettings, GameStates} from '../constants';

const BackEndPoint = "http://localhost:8080"

class Stadium extends React.Component {
  constructor(props) {
    super(props);
    this.reset()
  }

  componentDidMount() {
    const me = this;

    const evtSource = new EventSource(`${BackEndPoint}/game-state/${gameID}/${uuid}/`);

    const eventProcessor = function (event) {
      const g = JSON.parse(event.data);
      let team_goal = ""
      const newState = g.game_event?.game_snapshot?.state || GameStates.WAITING
      if (g.game_event?.game_snapshot?.state === GameStates.GET_READY) {

        const homeScoreOld = me.state.event.snapshot.home_team.Score ?? 0
        const awayScoreOld = me.state.event.snapshot.away_team.Score ?? 0

        const homeScoreNew = g.game_event?.game_snapshot?.home_team?.Score ?? 0
        const awayScoreNew = g.game_event?.game_snapshot?.away_team?.Score ?? 0


        console.log(`homeScoreOld: ${homeScoreOld}`)
        console.log(`awayScoreOld: ${awayScoreOld}`)
        console.log(`homeScoreNew: ${homeScoreNew}`)
        console.log(`awayScoreNew: ${awayScoreNew}`)

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

      me.setState(state);
    }

    evtSource.addEventListener("state_change", eventProcessor);
    evtSource.addEventListener("new_player", eventProcessor);
    evtSource.addEventListener("ping", function (event) {
      console.debug("ping")
    });

    // addEventListener version
    evtSource.addEventListener('open', () => {
      console.log("connection opened")
      me.reset()
      document.getElementById('lugobot-view').classList.remove("loading");
      this.setState({
        isConnected: true,
      });
      me.setup()
    });


    evtSource.onerror = () => {
      console.log("stream connection lost. trying to reconnect...");
      this.setState({
        isConnected: false,
        isLoaded: false,
      });
      me.reset()
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

      let headerGoalClass = ""
      if (this.state.event.team_goal !== "") {
        headerGoalClass = `active-modal`
      }

      this.setColor('--team-home-color-primary', this.state.setup.home_team.colors.primary);
      this.setColor('--team-home-color-secondary', this.state.setup.home_team.colors.secondary);
      this.setColor('--team-away-color-primary', this.state.setup.away_team.colors.primary);
      this.setColor('--team-away-color-secondary', this.state.setup.away_team.colors.secondary);
      return <div>
        <header id="lugobot-header" className={`container ${headerGoalClass}`}>
          <Panel event={this.state.event} setup={this.state.setup}/>
        </header>

        <main id="lugobot-stadium" className="container">
          <Field snapshot={this.state.event.snapshot}/>
        </main>
        <Events event={this.state.event}/>
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

          document.title = `Lugo - ${result.home_team.name} VS ${result.away_team.name}`
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

  reset() {
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
        team_goal: "",
        type: "",
        time_remaining: "00:00",
        snapshot: this.initSnapshot,
      }
    }
  }
}

export default Stadium;

