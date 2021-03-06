import React from 'react'
import SiteNavigation from "./SiteNavigation"
import PageContainer from "./PageContainer"
import PhotoGallery from "./PhotoGallery"
import TopPicks from './TopPicks'
import "./LP.css"


const  LandingPage = () => (
    <main id="lp">
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
    </main>
)
export default LandingPage