import React from 'react';
import EventGoal from './EventGoal'
import EventAlert from "./EventAlert";
import {StadiumStatus} from "../constants";
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";
import EventGameOver from "./EventGameOver";

class Events extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    let classList = []
    let team_goal = ""
    let displayNone = {}


    let modal = null
    let is_over = false
    switch (this.props.stadium_status) {
      case StadiumStatus.ALERT:
        classList = ["zoom-In", "active-modal"]
        modal = {
          title: this.props.title,
          text: this.props.text,
        }
        break;
      case StadiumStatus.GOAL:
        classList = ["zoom-In", "active-modal", "goal", `goal-${this.props.team_goal}`];
        team_goal = this.props.team_goal
        break;
      case StadiumStatus.OVER:
        classList = ["zoom-In", "active-modal"]
        is_over = true
        break
      default:
        displayNone = {display: "none"}
    }

    return <section id="event-view" style={displayNone} className={classList.join(" ")}>
      <EventAlert modal={modal}/>
      <EventGoal team_goal={team_goal}/>
      <EventGameOver show={is_over} />
      <svg version="1.1" xmlns="http://www.w3.org/2000/svg" id="filter_blur">
        <filter id="blur">
          <feGaussianBlur stdDeviation="6"/>
        </filter>
      </svg>
    </section>;
  }
}

Events.propTypes = {
  stadium_status: PropTypes.string,
  team_goal: PropTypes.string,
  title: PropTypes.string,
  text: PropTypes.object,
}

const mapStateToProps = state => {
  return {
    stadium_status: state.stadium.status,
    team_goal: state.stadium.event_data?.team_goal,
    title: state.stadium.event_data?.modal?.title,
    text: state.stadium.event_data?.modal?.text,
  }
}

export default connect(mapStateToProps)(Events)
