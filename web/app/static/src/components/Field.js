import React from 'react';
import FieldPlayer from "./FieldPlayer";
import {renderLogger} from '../helpers';
import {GameDefinitions} from "../constants";
import channel from "../channel";

const defaultPost = {x: -1000, y: -10000}
const presentPlayers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]

class Field extends React.Component {
  constructor(props) {
    super(props);
    this.ballDOM = React.createRef();
    this.players = {home: {}, away: {}};
    this.onNewFrameListeners = {home: {}, away: {}};
    const params = new URLSearchParams(window.location.search);
    this.grid_cols = parseInt(params.get("c")) ?? 1
    this.grid_rows = parseInt(params.get("r")) ?? 1
  }

  componentDidMount() {
    channel.subscribe((snapshot) => {
      const left = 100 * (snapshot.ball.position.x ?? 0) / GameDefinitions.Field.Width
      const bottom = 100 * (snapshot.ball.position.y ?? 0) / GameDefinitions.Field.Height

      this.ballDOM.current.style.left = `${left}%`;
      this.ballDOM.current.style.bottom = `calc(${bottom}%)`;

      if (snapshot.turns_ball_in_goal_zone > 0) {
        this.ballDOM.current.className = 'within_goal_zone'
      } else {
        this.ballDOM.current.className = ''
      }
      let homeMissed = presentPlayers.slice()
      snapshot.home_team?.players?.forEach((player) => {
        homeMissed.splice(homeMissed.indexOf(player.number), 1)
        this.onNewFrameListeners["home"][`home_${player.number}`](player)
      })

      homeMissed.forEach(playerNumber => {
        this.onNewFrameListeners["home"][`home_${playerNumber}`](false)
      })

      let awayMissed = presentPlayers.slice()
      snapshot.away_team?.players?.forEach((player) => {
        awayMissed.splice(awayMissed.indexOf(player.number), 1)
        this.onNewFrameListeners["away"][`away_${player.number}`](player)
      })

      awayMissed.forEach(playerNumber => {
        this.onNewFrameListeners["away"][`away_${playerNumber}`](false)
      })
    })
  }

  render() {
    renderLogger(this.constructor.name)
    const items = []

    const gridRows = []
    for (let j = this.grid_rows - 1; j >= 0; j--) {
      const lineCols = []
      for (let i = 0; i < this.grid_cols; i++) {
        lineCols.push(<span className="grid_cell" key={i}>{i}x{j}</span>)
      }
      gridRows.push(<span className="grid_lines" key={`line-${j}`}>{lineCols}</span>)
    }
    const vGrid = <div id="grid" key={"grid"}>{gridRows}</div>
    items.push(vGrid)

    const fillPlayer = (p, side) => {
      const a = <FieldPlayer
        setOnNewFrameListener={(cb) => this.onNewFrameListeners[side][`${side}_${p.number}`] = cb}
        key={`${side}-${p.number}`}
        number={p.number}
        team_side={side}
        ang={0}
        defaultPosition={this.defaultPosition(side, p.number)}
      />
      this.players[side][`${side}_${p.number}`] = a
      items.push(a)
    }

    for (let i = 1; i <= 11; i++) {
      fillPlayer({number: i}, "home")
    }

    for (let i = 1; i <= 11; i++) {
      fillPlayer({number: i}, "away")
    }

    return <div id="field">
      <span id="ball" style={{}} ref={this.ballDOM}/>
      {items}
    </div>;
  }

  defaultPosition(side, playerNumber) {
    let p = {x: (800 * (playerNumber -1)) + 500, y: -GameDefinitions.Player.Size}
    if(side === "away") {
      p.x = GameDefinitions.Field.Width - p.x
    }
    // console.log(`${side}: `, p)
    return p
  }
}

export default Field;
