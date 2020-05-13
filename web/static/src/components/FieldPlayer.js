import React from 'react';

import {GameDefinitions} from '../constants'
import {ShouldRender} from "../helpers";

class FieldPlayer extends React.Component {

  constructor(props) {
    super(props);
    this.myDOM = React.createRef();
    this.myDirectionDOM = React.createRef();
  }

  shouldComponentUpdate(nextProps, nextState) {
    return ShouldRender(this.props, nextProps);
  }

  onNewFrame(player) {


    // this.onNewFrameListener(player)
  }
  // setOnNewFrameListener(cb) {
  //   this.onNewFrameListener = cb
  // }

  // onNewFrameListener(cb) {
  //   console.log(`${this.constructor.name}: onNewFrameListener not listened`)
  // }

  componentDidMount() {
    this.props.setOnNewFrameListener(player => {
      console.log(`Player ${this.props.team_side} -> ${this.props.number}`)
      // const me = player[this.props.team_side][`${this.props.team_side}_${this.props.number}`]
      const left = 100 * (player.Position.X ?? 0) / GameDefinitions.Field.Width
      const bottom = 100 * (player.Position.Y ?? 0) / GameDefinitions.Field.Height

      this.myDOM.current.style.left = `${left}%`;
      this.myDOM.current.style.bottom = `calc(${bottom}%)`;

      console.log(this.myDirectionDOM)
      this.myDirectionDOM.current.style.transform = `rotate(${-player.velocity.direction.ang + 90}deg)`;

    })
  }

  render() {
    return <span ref={this.myDOM}
      id={`player-${this.props.team_side}-team-${this.props.number}`}
      className={"player player-"+this.props.team_side+"-team"}
      rel={this.props.number}
      style={{left: 0, bottom: 0}}
    >
        <span className="player-direction" ref={this.myDirectionDOM} style={{transform: null}} />
        <span className="player-number">{this.props.number}</span>
      </span>
  }

}
export default FieldPlayer;