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
    if(this.props.activate) {
      displayNone = {}
      classList = ["zoom-In", "active-modal"];
      switch (this.props.mode) {
        case ModalModes.GOAL:
          classList = ["zoom-In", "active-modal", "goal", `goal-${this.props.team_side }`];
          team_goal = this.props.team_side
          break;
      }
    }

    return <section id="event-view" style={displayNone} className={classList.join(" ")}>
      Gool {this.props.team_side}
      {/*<Modal modal={this.props.modal}/>*/}
      {/*<EventGoal team_goal={team_goal}/>*/}
    </section>;
  }
}

Events.propTypes = {
  activate: PropTypes.bool,
  mode: PropTypes.string,
  team_side: PropTypes.string,
}

const mapStateToProps = state => {
  return state.match.modal

}

export default connect(mapStateToProps)(Events)
