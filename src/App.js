import React, { Component } from 'react';
import SiteNavigation from "./components/SiteNavigation"
import PageContainer from "./components/PageContainer"
import TopPicks from "./components/TopPicks"
import logo from './logo.svg';
import './App.css';
import PhotoGallery from './components/PhotoGallery';

class App extends Component {
  render() {
    return (
      <React.Fragment>
        <div className="header">
          <SiteNavigation />
          <PageContainer />
        </div>
        <div className="mt-8">
          <TopPicks />
        </div>

        <section className="mt-8">
          <PhotoGallery />
        </section>

      </React.Fragment>
    );
  }
}

export default App;
