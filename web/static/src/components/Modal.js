import React from 'react';

class Modal extends React.Component {
  render() {
    if(this.props.modal === null) {
      return <span />
    }
    return <div id="modal-game-over" className={`modal modal-bg active-modal zoom-In`}>
      <span key="b" className="modal-content">
        <h2 className="modal-title">{this.props.modal.title}</h2>
        <p className="modal-content">{this.props.modal.text}</p>
      </span>
      <span key="c" className="modal-action">
        {/*<button id="btn-quit-game" className="btn">Quit Game</button>*/}
        {/*<button id="btn-retry" className="btn btn-main">Retry</button>*/}
      </span>
    </div>;
  }

}

export default Modal;