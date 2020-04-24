import React from 'react';
import Panel from "./Panel";

class Stadium extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      game: {
        homeTeam: {
          name: "Rubens",
          avatar: "src/img/teams/profile-team-home.jpg",
          score: 0,
          colors: {
            a: [0, 250, 0],
            b: [250, 200, 0]
          },
        },
        awayTeam: {
          name: "Outro",
          avatar: "src/img/teams/profile-team-away.jpg",
          score: 3,
          colors: {
            a: [0, 0, 200],
            b: [50, 100, 200]
          },
        },
      }
    };

    document.documentElement.style.setProperty('--team-home-color-primary', this.state.game.homeTeam.colors.a);
    document.documentElement.style.setProperty('--team-home-color-secondary',this.state.game.homeTeam.colors.b);
    document.documentElement.style.setProperty('--team-away-color-primary',this.state.game.awayTeam.colors.a);
    document.documentElement.style.setProperty('--team-away-color-secondary',this.state.game.awayTeam.colors.b);
  }

  render() {
    const me = this;
    setTimeout(function () {

      me.setState(state => {
        let s = state;
        s.game.homeTeam.name = "Championgs";
        return s;
      })

      console.log("mudou")
    }, 3000)
    return <span>
      <header id="lugobot-header" className="container">
        <Panel game={this.state.game}/>
      </header>
    </span>;
  }
}

export default Stadium;

