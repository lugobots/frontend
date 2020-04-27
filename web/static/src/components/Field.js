import React from 'react';

class Field extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    return <div id="field">
      <span id="ball"/>
    </div>;
  }
}

export default Field;