import React from 'react';
import FieldPlayer from "./FieldPlayer";

import {GameDefinitions} from '../constants';

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
    // this.players = {
    //   home: {},
    //   away: {},
    // };
    //
    // const homePLayers = document.getElementsByClassName("player-home-team");
    // for (let i = 0; i < homePLayers.length; i++) {
    //   this.players.home[`home-${homePLayers[i].attributes.rel.value}`] = homePLayers[i]
    // }
    // const awayPLayers = document.getElementsByClassName("player-away-team");
    // for (let i = 0; i < awayPLayers.length; i++) {
    //   this.players.away[`away-${awayPLayers[i].attributes.rel.value}`] = awayPLayers[i]
    // }
    this.props.setOnNewFrameListener(snapshot => {
      const me = this;
      console.log("RECEBIDO")
      // let list = {home: {}, away: {}}

      // snapshot.home_team?.players.forEach( (player) =>{
      //   list["home"][`home_${player.number}`] = {position: player.Position, ang: player.ang}
      // })
      // snapshot.away_team?.players.forEach( (player) =>{
      //   list["away"][`away_${player.number}`] = {position: player.Position, ang: player.ang}
      // })
      snapshot.home_team.players.forEach( (player) =>{
        this.onNewFrameListeners["home"][`home_${player.number}`](player)
      })
      snapshot.away_team?.players.forEach( (player) =>{
        this.onNewFrameListeners["away"][`away_${player.number}`](player)
      })

    })
  }

  // updatePlayer(side, number, position) {
  //   const ball_left = 100 * (position.X ?? 0) / GameDefinitions.Field.Width
  //   const ball_bottom = 100 * (position.Y ?? 0) / GameDefinitions.Field.Height
  //   let list = this.players.home
  //   if(side === "away") {
  //     list = this.players.away
  //   }
  //   // const direction = {
  //   //   transform: `rotate(${-this.props.ang + 90}deg)`
  //   // }
  //
  //   console.log(`${side}-${number}`)
  //   const p = list[`${side}-${number}`]
  //   p.style.left = `${ball_left}%`
  //   p.style.bottom = `calc(${ball_bottom}%)`
  // }

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