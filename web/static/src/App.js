import React from 'react';
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import action from "./redux/app/actions";
import {AppStatus, BackendConfig, EventTypes} from "./constants";
import store from "./store";
import Stadium from "./components/Stadium";

class App extends React.Component {

  setup() {
    fetch(`${BackendConfig.BackEndPoint}/setup/${gameID}/${uuid}`)
      .then(res => res.json())
      .then(
        (result) => {
          // console.log("%cSetup", "color: blue")
          // console.log(result)
          this.props.dispatch(action.setup(result))
        },
        (error) => {
          // @todo handle the error
          alert("oops... something went wrong")
        }
      )



  }

  componentDidMount() {
    store.subscribe(() => {
      if(store.getState().status === AppStatus.Setting) {
        this.setup()
      }
    })

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
    return <h1>Hello, {this.props.status}
      <Stadium setup={this.props.setup} />
    </h1>;
  }

}
App.propTypes = {
  dispatch: PropTypes.func.isRequired
}

const mapStateToProps = state => {

  return state
}

export default connect(mapStateToProps)(App)