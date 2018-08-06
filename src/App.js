import React, { Component } from 'react';
import {Switch, BrowserRouter, Route} from 'react-router-dom'
import './App.css';
import LandingPage from './components/LandingPage';
import PhotoDetail from './components/PhotoDetail';

const App = () =>(
  <React.Fragment>
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={LandingPage} />
        <Route exact path="/photos/:id" component={PhotoDetail} />
      </Switch>
    </BrowserRouter>
  </React.Fragment>
);

export default App;
