import React from 'react';
import PanelTeamsInfo from "./PanelTeamsInfo";
import PanelGameInfo from "./PanelGameInfo";
import {renderLogger} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";
import {ModalModes} from "../constants";

class Panel extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    renderLogger(this.constructor.name)
    let  modal_class = ""
    if(this.props.modal_mode === ModalModes.GOAL) {
      modal_class = "active-modal"
    }

    return <header id="lugobot-header"
                   className={`container ${modal_class}`}>
      <section id="game-panel">
        <PanelTeamsInfo/>
        <PanelGameInfo/>
      </section>
    </header>;
  }
}

Panel.propTypes = {
  modal_mode: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    modal_mode: state.stadium.modal.mode,
  }
}

export default connect(mapStateToProps)(Panel)