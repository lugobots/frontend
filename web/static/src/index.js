import React from 'react';
import { Provider } from 'react-redux'
import ReactDOM from 'react-dom';
import { createStore } from 'redux'
import game from './reducers'
import './sass/style.scss';

import {library} from '@fortawesome/fontawesome-svg-core';
import {faEnvelope, faKey} from '@fortawesome/free-solid-svg-icons';
import App from "./App";
library.add(faEnvelope, faKey);

const store = createStore(game)
// Log the initial state
console.log(store.getState())

const unsubscribe = store.subscribe(() => console.log(store.getState()))

// Stop listening to state updates
unsubscribe()

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)