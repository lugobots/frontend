import React from 'react';
import {ToolBarTabDebug} from './ToolBarTabDebug'
import {renderLogger} from "../helpers";

class ToolBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      open_tab: "debug",
      // debugOn: false,
    }
  }

  openTab(tab) {
    this.setState({
      open_tab: tab,
    })
  }

  componentDidMount() {
    // this.props.setOnNewEventListener(gameEvent => {
    //
    //   // if(this.state.debugOn !== gameEvent.debug_mode) {
    //   //   this.setState({
    //   //     debugOn: gameEvent.debug_mode
    //   //   })
    //   // }
    //
    // })
  }

  render() {
    renderLogger(this.constructor.name)
    return <footer id="lugobot-admin-panel" className="container debug-mode">
      <nav id="tabs-panel-link">
        <li className="tab-link active-tab-link">
          <a onClick={() => this.openTab("debug")}>DEBUG</a>
        </li>
        {/*<li className="tab-link">*/}
        {/*  <a onClick={() => this.openTab("b")}>EVENTS</a>*/}
        {/*</li>*/}
        {/*<li className="tab-link">*/}
        {/*  <a onClick={() => this.openTab("c")}>OTHERS</a>*/}
        {/*</li>*/}
        {/*<li className="bg-tab"><a>BG TAB</a></li>*/}
      </nav>
      <section id="tabs-panel-content">
        {
          {
            'debug': <ToolBarTabDebug
              // setup={this.state.setup}
              // stadium_state={this.props.stadium_state}
              // gotoStateDebugging={this.props.gotoStateDebugging}
              className="tab-content active-tab-content"
            />,
            // 'a': <div>B</div>,
          }[this.state.open_tab]
        }
      </section>
    </footer>;
  }

}

export default ToolBar;