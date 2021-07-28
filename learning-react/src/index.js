import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

//this says render the content we pass in to <App/> at the 
//certain element (in this case document.getElementById('root')
ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
