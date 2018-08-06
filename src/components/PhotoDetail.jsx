import React from "react"
import SiteNavigation from './SiteNavigation'

const PhotoDetail = () => (
  <main>
    <SiteNavigation />

    <div className="pageContainer mt-8 mr-4">
      <div className="container">
        <div className="photo-detail">
          <div className="left image-collage">
            <img className="border-black border-2" src="https://burst.shopifycdn.com/photos/woman-shopping-for-clothes_925x@2x.progressive.jpg" alt="" />
            <div>
              <h3 className="mb-4 mt-12">Related Images</h3>
              <div className="related-images">
                <img className="mr-4" src="https://burst.shopifycdn.com/photos/purple-cropped-wig-with-glittery-drag-fashion_x182.progressive.jpg" />
                <img className="mr-4" src="https://images.pexels.com/photos/53754/pexels-photo-53754.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260" />
                <img className="mr-4" src="https://images.pexels.com/photos/718976/pexels-photo-718976.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260" />
                <img className="mr-4" src="https://cdn.pixabay.com/photo/2014/05/11/13/39/bald-eagles-341898__180.jpg" />
                <img className="mr-4" src="https://images.pexels.com/photos/823223/pexels-photo-823223.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260" />
              </div>
            </div>
            
          </div>

          <div className="user-details px-4 ml-8">
            <h2>Women Collection</h2>
            <p>
              A young couple walks in the rain with their arms around each other. 
              The cold weather cannot stop these two lovebirds, 
              they have their umbrella and each other to keep them warm.
            </p>

            <div className="flex mt-8">
              <a href="/user/@drexender">
                <img className="rounded-full author-dp mr-2" src="https://cdn.pixabay.com/user/2018/05/31/00-43-58-804_48x48.jpg" />
              </a>
              <p className="imageBy">Alexander Dre</p>
            </div>
          </div>
        </div>

        
      </div>
    </div>
  </main>
)

export default PhotoDetail