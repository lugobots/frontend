import React from 'react';
import TooBarTabs from './TooBarTabs'

class ToolBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      open_tab: "",
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
          <a onClick={() => this.openTab("a")}>DEBUG</a>
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
        <TooBarTabs active_tab={this.state.open_tab} id="a"> <div>A</div></TooBarTabs>
        <TooBarTabs active_tab={this.state.open_tab} id="b"> <div>B</div></TooBarTabs>
        <TooBarTabs active_tab={this.state.open_tab} id="c"> <div>C</div></TooBarTabs>
      </section>
    </footer>;
  }

}

export default ToolBar;