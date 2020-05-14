import React from 'react';
import EventGoal from './EventGoal'
import Modal from "./Modal";

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
    if (this.state.team_goal !== "") {
      style = ["zoom-In", "active-modal", "goal", `goal-${this.state.team_goal}`];
    } else if (this.props.modal !== null) {
      style = ["zoom-In", "active-modal"];
    }

    return <section id="event-view" className={style.join(" ")}>
      <Modal modal={this.props.modal}/>
      <EventGoal team_goal={this.state.team_goal}/>
    </section>;
  }
}

export default Events;

