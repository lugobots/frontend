import React from 'react';
import EventGoal from './EventGoal'


class Events extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {

    let style = ""
    if(this.props.event.team_goal !== "") {
      style = `zoom-In active-modal goal goal-${this.props.event.team_goal}`
    }

    return <section id="event-view" className={style}>
      <EventGoal team_goal={this.props.event.team_goal}/>
    </section>;
  }
}

export default Events;

