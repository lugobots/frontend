import React from 'react';
import EventGoal from './EventGoal'
import Modal from "./Modal";
import {ModalModes, StadiumStates} from "../constants";
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";

class Events extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    let classList = []
    let team_goal = ""
    let displayNone = {display: "none"}
    let modal = null
    if(this.props.activate) {
      displayNone = {}
      classList = ["zoom-In", "active-modal"];
      switch (this.props.mode) {
        case ModalModes.GOAL:
          classList = ["zoom-In", "active-modal", "goal", `goal-${this.props.team_side }`];
          team_goal = this.props.team_side
          break;
        case ModalModes.ALERT:
          modal = {
            title: this.props.title,
            text: this.props.text,
          }
          break;
      }
    }

    return <section id="event-view" style={displayNone} className={classList.join(" ")}>
      {<Modal modal={modal}/>}
      {<EventGoal team_goal={team_goal}/>}
      <svg version="1.1" xmlns="http://www.w3.org/2000/svg" id="filter_blur">
        <filter id="blur"> <feGaussianBlur stdDeviation="6" /></filter>
      </svg>
    </section>;
  }
}

Events.propTypes = {
  activate: PropTypes.bool,
  mode: PropTypes.string,
  team_side: PropTypes.string,
  title: PropTypes.string,
  text: PropTypes.object,
}

const mapStateToProps = state => {
  return state.match.modal

}

export default connect(mapStateToProps)(Events)
