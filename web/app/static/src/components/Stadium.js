import React from 'react';
import Panel from "./Panel";
import Field from "./Field";
import Events from "./Events";
import ToolBar from "./ToolBar";

import {StadiumStatus} from '../constants';
import {renderLogger, updateRatio} from "../helpers";
import PropTypes from "prop-types";
import {connect} from "react-redux";


class Stadium extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {

    updateRatio()

  }

  setMainColor(name, colors) {
    const lis = [colors.red ?? 0, colors.green ?? 0, colors.blue ?? 0]
    document.documentElement.style.setProperty(name, lis.toString());
  }

  render() {
    renderLogger(this.constructor.name)
    this.setMainColor('--team-home-color-primary', this.props.setup.home_team.colors.primary);
    this.setMainColor('--team-home-color-secondary', this.props.setup.home_team.colors.secondary);
    this.setMainColor('--team-away-color-primary', this.props.setup.away_team.colors.primary);
    this.setMainColor('--team-away-color-secondary', this.props.setup.away_team.colors.secondary);

    let stadium_class = this.props.stadium_status.toLowerCase()
    if (this.props.stadium_status === StadiumStatus.ALERT ||
      this.props.stadium_status === StadiumStatus.OVER) {
      stadium_class = " active-modal"
    }
    return <div id="stadium" className={stadium_class}>
      <Panel/>
      <main id="lugobot-stadium" className="container">
        <Field/>
      </main>
      {this.props.setup.dev_mode ? <ToolBar/> : null}
      <Events/>
    </div>;
  }
}

Stadium.propTypes = {
  setup: PropTypes.object,
  stadium_status: PropTypes.string,
}

const mapStateToProps = state => {
  return {
    setup: state.app.setup,
    stadium_status: state.stadium.status,
  }
}

export default connect(mapStateToProps)(Stadium)

