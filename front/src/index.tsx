import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import SetupRoutingComponent from "./routing/SetupRoutingComponent";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

console.log("Backend: ", process.env.REACT_APP_BACKEND_URL)
console.log("Repository: ", process.env.GITHUB_REPOSITORY, process.env.GITHUB_REPOSITORY_URL)
console.log("Version: ", process.env.GITHUB_SHA)

root.render(
  <React.StrictMode>
      <SetupRoutingComponent />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
