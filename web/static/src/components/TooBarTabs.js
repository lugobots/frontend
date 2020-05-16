import React from "react";

class TooBarTabs extends React.Component {

  render() {
    if(this.props.active_tab === this.props.id) {
      return <div className="tab-content debug-tab active-tab-content">
        {this.props.children}
      </div>
    }
    return false
  }
}

export default TooBarTabs;
