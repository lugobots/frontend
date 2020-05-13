import React from 'react';

class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }

  setName(name) {
    this.setState({name})
  }
}

export default Welcome;