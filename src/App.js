import React, { Component } from 'react';
import {Switch, Route, BrowserRouter as Router} from 'react-router-dom'
import './App.css';
import LandingPage from "./components/LandingPage"
import PhotoDetail from "./components/PhotoDetail"

class App extends Component {
  render() {
    return (
      <React.Fragment>
        <Router>
          <Switch>
            <Route exact path="/" component={LandingPage} />
            <Route path="/photos/:photo-slug" component={PhotoDetail} />
          </Switch>
        </Router>

      </React.Fragment>
    );
  }
}

export default App;
