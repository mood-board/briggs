import React, { Component } from 'react';
import SiteNavigation from "./components/SiteNavigation"
import PageContainer from "./components/PageContainer"
import TopPicks from "./components/TopPicks"
import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <React.Fragment>
        <div class="header">
          <SiteNavigation />
          <PageContainer />
        </div>
        <div className="mt-8">
          <TopPicks />
        </div>
      </React.Fragment>
    );
  }
}

export default App;
