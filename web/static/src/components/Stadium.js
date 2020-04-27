import React from 'react';
import Panel from "./Panel";
import Field from "./Field";


const BackEndPoint = "https://localhost:8080"

class Stadium extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isConnected: false,
      isLoaded: false,
      error: null,
      game: {
        timeRemaining: "00:00",
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
    const evtSource = new EventSource(`${BackEndPoint}/game-state/${gameID}/${uuid}`);

    evtSource.addEventListener("ping", function (event) {
      console.log("ping", event.data);
    });

    // addEventListener version
    evtSource.addEventListener('open', () => {
      me.setState(state => {
        let s = state;
        s.isConnected = true;
        return s;
      })
      me.setup()
    });


    evtSource.onerror = function () {
      console.debug("EventSource failed.");
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
      return <div>
      <header id="lugobot-header" className="container">
        <Panel game={this.state.game}/>
      </header>

      <main id="lugobot-stadium" className="container">
        <Field game={this.state.game}/>
      </main>
    </div>;
    }
  }

  setup() {
    fetch(`${BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(res => res.json())
      .then(
        (result) => {
          console.log(result)
          this.setState({
            isLoaded: true,
            game: result
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

