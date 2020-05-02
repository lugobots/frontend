import React from 'react';

class FieldPlayer extends React.Component {
  render() {

    

    AQUI!!!! COlocar posicaodo jogador dinamicament

    return <span id={"player-"+this.props.team_side+"-team-1"} className={"player player-"+this.props.team_side+"-team"}>
        <span className="player-direction"/>
        <span className="player-number">{this.props.number}</span>
      </span>
  }
}

export default FieldPlayer;