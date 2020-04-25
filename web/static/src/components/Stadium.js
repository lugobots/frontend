import React from 'react';
import Panel from "./Panel";
import axios from 'axios';

class Stadium extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isConnected: false,
      isLoaded: false,
      error: null,
      game: {
        homeTeam: {
          name: "Rubens",
          avatar: "external/profile-team-home.jpg",
          score: 0,
          colors: {
            a: [0, 250, 0],
            b: [250, 200, 0]
          },
        },
        awayTeam: {
          name: "Outro",
          avatar: "external/profile-team-away.jpg",
          score: 3,
          colors: {
            a: [0, 0, 200],
            b: [50, 100, 200]
          },
        },
      }
    };
  }

  componentDidMount() {

    const me = this;
    const evtSource = new EventSource("https://localhost:8080/stream");

    evtSource.addEventListener("setup", function (event) {
      console.log("PING.", event.data);
      me.setState(state => {
        let s = state;
        s.game = JSON.parse(event.data)
        s.isLoaded = true;
        return s;
      })
    });

    // addEventListener version
    evtSource.addEventListener('open', (e) => {
      console.log("Connection to server opened.");
      me.setState(state => {
        let s = state;
        s.isConnected = true;
        return s;
      })
    });


    evtSource.onerror = function () {
      console.log("EventSource failed.");
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
      document.documentElement.style.setProperty('--team-home-color-primary', this.state.game.homeTeam.colors.a);
      document.documentElement.style.setProperty('--team-home-color-secondary', this.state.game.homeTeam.colors.b);
      document.documentElement.style.setProperty('--team-away-color-primary', this.state.game.awayTeam.colors.a);
      document.documentElement.style.setProperty('--team-away-color-secondary', this.state.game.awayTeam.colors.b);
      return <span>
      <header id="lugobot-header" className="container">
        <Panel game={this.state.game}/>
      </header>
    </span>;
    }
  }
}

export default Stadium;

