import React from 'react';
import FieldPlayer from "./FieldPlayer";


import {ShouldRender} from '../helpers';
import {GameDefinitions} from "../constants";
const defaultPost = {Y: -1000, X: -10000}

class Field extends React.Component {
  constructor(props) {
    super(props);
    this.ballDOM = React.createRef();
    this.players = {home: {}, away: {}};
    this.onNewFrameListeners = {home: {}, away: {}};
  }

  shouldComponentUpdate(nextProps, nextState) {
    return ShouldRender(this.props, nextProps);
  }

  componentDidMount() {
    this.props.setOnNewEventListener(gameEvent => {
      const snapshot = gameEvent.snapshot;
      const left = 100 * (snapshot.ball.Position.X ?? 0) / GameDefinitions.Field.Width
      const bottom = 100 * (snapshot.ball.Position.Y ?? 0) / GameDefinitions.Field.Height

      this.ballDOM.current.style.left = `${left}%`;
      this.ballDOM.current.style.bottom = `calc(${bottom}%)`;

      snapshot.home_team?.players?.forEach( (player) =>{
        this.onNewFrameListeners["home"][`home_${player.number}`](player)
      })
      snapshot.away_team?.players?.forEach( (player) =>{
        this.onNewFrameListeners["away"][`away_${player.number}`](player)
      })
    })
  }

  render() {
    const items = []
    const fillPlayer = (p, side) => {
      const a = <FieldPlayer
        setOnNewFrameListener={(cb) => this.onNewFrameListeners[side][`${side}_${p.number}`] = cb}
        key={`${side}-${p.number}`}
        number={p.number}
        team_side={side}
        ang={p.ang}
        position={p.Position}
      />
      this.players[side][`${side}_${p.number}`] = a
      items.push(a)
    }

    for (let i = 1; i <= 11; i++) {
        fillPlayer({ number: i, ang: 0, Position: defaultPost}, "home")
    }

    for (let i = 1; i <= 11; i++) {
      fillPlayer({ number: i, ang: 0, Position: defaultPost}, "away")
    }

    return <div id="field">
            <span id="ball" ref={this.ballDOM} />
            {items}
          </div>;
  }
}

export default Field;