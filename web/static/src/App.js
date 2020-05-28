import React from 'react';
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import action from "./actions";
import {BackendConfig, EventTypes} from "./constants";

class App extends React.Component {
  componentDidMount() {
    this.evtSource = new EventSource(`${BackendConfig.BackEndPoint}/game-state/${gameID}/${uuid}/`);
    // addEventListener version
    this.evtSource.addEventListener('open', () => {
      this.props.dispatch(action.backConnect())
    });
    this.evtSource.onerror = () => {
      this.props.dispatch(action.backDisconnect())
    };

    this.evtSource.addEventListener("ping", () => {console.debug("ping")});

    this.evtSource.addEventListener(EventTypes.ConnectionLost, (event) => {
      console.error("%cupstream connection lost", "color: #AA0000")
      this.props.dispatch(action.upstreamDisconnect())
    });
    this.evtSource.addEventListener(EventTypes.ConnectionReestablished, (event) => {
      console.log("%cupstream connection reestablished", "color: green")
      this.props.dispatch(action.upstreamConnect())
    });
  }

  render() {
    return <h1>Hello, {this.props.status}</h1>;
  }

}
App.propTypes = {
  dispatch: PropTypes.func.isRequired
}

const mapStateToProps = state => {

  return state
}

export default connect(mapStateToProps)(App)