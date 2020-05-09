import React from 'react';
import EventGoal from './EventGoal'
import Modal from "./Modal";

class Events extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {

    let style = []
    if (this.props.event.team_goal !== "") {
      style = ["zoom-In", "active-modal", "goal", `goal-${this.props.event.team_goal}`];
    } else if (this.props.modal !== null) {
      style = ["zoom-In", "active-modal"];
    }

    return <section id="event-view" className={style.join(" ")}>
      <Modal modal={this.props.modal}/>
      <EventGoal team_goal={this.props.event.team_goal}/>
    </section>;
  }
}

export default Events;

