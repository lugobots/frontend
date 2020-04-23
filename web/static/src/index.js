import React from 'react';
import ReactDOM from 'react-dom';
import './sass/apps.scss';
import {library} from '@fortawesome/fontawesome-svg-core';
import {faEnvelope, faKey} from '@fortawesome/free-solid-svg-icons';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
library.add(faEnvelope, faKey);

ReactDOM.render(<h1>Helloworld React!<FontAwesomeIcon icon="envelope"/></h1>, document.getElementById('root'));