import React from 'react';
import FieldPlayer from "./FieldPlayer";


import {ShouldRender} from '../helpers';
const defaultPost = {Y: -1000, X: -10000}

class Field extends React.Component {
  constructor(props) {
    super(props);
    this.players = {home: {}, away: {}};
    this.onNewFrameListeners = {home: {}, away: {}};
  }

  shouldComponentUpdate(nextProps, nextState) {
    return ShouldRender(this.props, nextProps);
  }

  componentDidMount() {
    this.props.setOnNewFrameListener(snapshot => {
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
            <span id="ball"/>
            {items}
          </div>;
  }
}

export default Field;