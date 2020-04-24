import React from 'react';
import ReactDOM from 'react-dom';
import './sass/style.scss';
import {library} from '@fortawesome/fontawesome-svg-core';
import {faEnvelope, faKey} from '@fortawesome/free-solid-svg-icons';
library.add(faEnvelope, faKey);

import Stadium from "./components/Stadium"

ReactDOM.render(<Stadium  />, document.getElementById('app'));