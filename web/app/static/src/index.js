import React from 'react';
import { Provider } from 'react-redux'
import ReactDOM from 'react-dom';
import store from './store'
import './sass/style.scss';

import {library} from '@fortawesome/fontawesome-svg-core';
import {faEnvelope, faKey} from '@fortawesome/free-solid-svg-icons';
import App from "./App";
library.add(faEnvelope, faKey);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)