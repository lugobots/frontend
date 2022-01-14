import React from 'react';
import FieldPlayer from "./FieldPlayer";
import {renderLogger} from '../helpers';
import {GameDefinitions} from "../constants";
import channel from "../channel";
const defaultPost = {x: -1000, y: -10000}

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
      snapshot.home_team?.players?.forEach((player) => {
        this.onNewFrameListeners["home"][`home_${player.number}`](player)
      })
      snapshot.away_team?.players?.forEach((player) => {
        this.onNewFrameListeners["away"][`away_${player.number}`](player)
      })
    })
  }

  render() {
    renderLogger(this.constructor.name)
    const items = []
    const fillPlayer = (p, side) => {
      const a = <FieldPlayer
        setOnNewFrameListener={(cb) => this.onNewFrameListeners[side][`${side}_${p.number}`] = cb}
        key={`${side}-${p.number}`}
        number={p.number}
        team_side={side}
        ang={p.ang}
        position={p.position}
      />
      this.players[side][`${side}_${p.number}`] = a
      items.push(a)
    }

    for (let i = 1; i <= 11; i++) {
      fillPlayer({number: i, ang: 0, position: defaultPost}, "home")
    }

    for (let i = 1; i <= 11; i++) {
      fillPlayer({number: i, ang: 0, position: defaultPost}, "away")
    }

    const gridRows = []
    for (let j = this.grid_rows - 1; j >= 0 ; j--) {
      const lineCols = []
      for (let i = 0; i < this.grid_cols; i++) {
        lineCols.push(<span className="grid_cell" key={i}>{i}x{j}</span>)
      }
      gridRows.push(<span className="grid_lines" key={`line-${j}`}>{lineCols}</span>)
    }


    const vGrid = <div id="grid" key={"grid"}>{gridRows}</div>
    // document.documentElement.style.setProperty("--grid-cols", `${this.grid_cols}`);
    // document.documentElement.style.setProperty("--grid-rows", `${this.grid_rows}`);

    items.push(vGrid)
    return <div id="field">
      <span id="ball" style={{}} ref={this.ballDOM}/>
      {items}
    </div>;
  }
}

export default Field;
