import React from 'react';
import ToolBarTabDebug from './ToolBarTabDebug'

class ToolBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      open_tab: "debug",
    }
  }

  openTab(tab) {
    this.setState({
      open_tab: tab,
    })
  }

  render() {
    return <footer id="lugobot-admin-panel" className="container debug-mode">
      <nav id="tabs-panel-link">
        <li className="tab-link active-tab-link">
          <a onClick={() => this.openTab("debug")}>DEBUG</a>
        </li>
        <li className="tab-link active-tab-link">
          <a onClick={() => this.openTab("b")}>EVENTS</a>
        </li>
        <a onClick={() => this.openTab("c")}>OTHERS</a>
        <li className="tab-link active-tab-link">
        </li>
        <li className="bg-tab"><a>BG TAB</a></li>
      </nav>
      <section id="tabs-panel-content">
        {
          {
            'debug': <ToolBarTabDebug className="tab-content active-tab-content" />,
            'a': <div>B</div>,
            'b': <div>C</div>,
          }[this.state.open_tab]
        }
      </section>
    </footer>;
  }

}

export default ToolBar;