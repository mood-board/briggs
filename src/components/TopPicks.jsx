import React from 'react'
import "./TopPicks.css"

const TopPicks = () => {
  return (
    <section className="top-picks">
      <div className="tp-container container">
          <div>
            <h1>Top free pics this week</h1>
            <h3 className="mb-8">Stunning stock images, perfect for blogs and websites</h3>
          </div>

          <div className="flex scrolling-carousel">
            <div className="scroll-item mr-2 ml-2">
              <img className="rounded-lg" style={{height: "285px", width: "428px"}} src="https://burst.shopifycdn.com/photos/motorcycle-caravan-rides-hillside-in-black-and-white_x285.progressive.jpg" />
            </div>


            <div className="scroll-item mr-2 ml-2 ">
              <img className="rounded-lg" style={{height: "285px", width: "428px"}} src="https://burst.shopifycdn.com/photos/robot-components_x285.progressive.jpg" />
            </div>

            <div className="scroll-item mr-2 ml-2">
              <img className="rounded-lg" style={{height: "285px", width: "428px"}} src="https://burst.shopifycdn.com/photos/horizontal-flatlay-of-marijuana-bud-and-concentrates_x285.progressive.jpg" />
            </div>

            {/**<div className="scroll-item mr-2 ml-2">
              <img className="rounded-lg" style={{height: "285px", width: "428px"}} src="https://burst.shopifycdn.com/photos/streetcar-into-the-light_x285.progressive.jpg" />
            </div>

            <div className="scroll-item mr-2 ml-2 ">
              <img className="rounded-lg" style={{height: "285px", width: "428px"}} src="https://burst.shopifycdn.com/photos/robot-components_x285.progressive.jpg" />
            </div> **/}

          </div>
      </div>
    </section>
  )
}

export default TopPicks;
