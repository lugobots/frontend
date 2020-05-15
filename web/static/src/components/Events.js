import React from 'react';
import EventGoal from './EventGoal'
import Modal from "./Modal";
import {StadiumStates} from "../constants";

class Events extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      team_goal: "",
    }
  }

  componentDidMount() {
    this.props.setOnNewEventListener(gameEvent => {
      if(gameEvent.team_goal !== this.state.team_goal) {
        this.setState({team_goal: gameEvent.team_goal})
      }
    })
  }

  render() {
    console.log(`${this.constructor.name} rendered`)
    let style = []
    let team_goal = ""
    switch (this.props.stadium_state.mode) {
      case StadiumStates.StadiumStateGoal:
        style = ["zoom-In", "active-modal", "goal", `goal-${this.props.stadium_state.side }`];
        team_goal = this.props.stadium_state.side
        break;
    }
    if(this.props.modal !== null) {
      style = ["zoom-In", "active-modal"];
    }

    return <section id="event-view" className={style.join(" ")}>
      <Modal modal={this.props.modal}/>
      <EventGoal team_goal={team_goal}/>
    </section>;
  }
}

export default Events;

