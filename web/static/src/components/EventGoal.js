import React from 'react';

class EventGoal extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const snow = []
    for (let i = 0; i < 100; i++) {
      snow.push(<div key={`snow-${i}`} className="snow"/>)
    }

    let className = ""
    if(this.props.team_goal === "home" || this.props.team_goal === "away") {
      className = `active-modal goal goal-${this.props.team_goal}`
    }

    return <div id="event-goal" className={className}>
            <span id="balls">{snow}</span>
            <h2 className="legend-goal">
              <span key="G">G</span>
              <span key="O0">O</span>
              <span key="O1">O</span>
              <span key="A0">A</span>
              <span key="A1">A</span>
              <span key="Ls">L</span>
              <span key="EX">!</span>
            </h2>
            <span className="soccer-crowd"/>
        </div>;
  }
}

export default EventGoal;

