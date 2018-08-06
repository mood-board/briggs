import React from 'react'

const PageContainer = () => {
  return (
    <div className="page-container">
      <div className="container">
        <div className="call-to-a w-3/5 py-4 px-2">
          <h1 className="cta-h1 mb-4">Free stock photos for websites and commercial use</h1>
          <h2 className="cta-h2 mb-4">Download free, high-resolution images</h2>

          <input id="docsearch" className="transition focus:outline-0 border border-transparent
            focus:bg-white focus:border-grey-light placeholder-grey-darkest rounded bg-grey-lighter
            py-4 pr-4 pl-10 block w-full appearance-none leading-normal ds-input" type="text"
            placeholder="Search high resolution images" />

          <div className="cta-popular-categories mt-8">
            <div className="flex">
              <ul className="flex">
                <li>
                  <span className="bold popular-cats">Popular Categories: </span>
                </li>
                <li className="ml-4 mr-4">
                  <a href="/shoes">Shoes</a>
                </li>

                <li className="mr-4">
                  <a href="/food">Food</a>
                </li>

                <li className="mr-4">
                  <a href="/travel">Travel</a>
                </li>

                <li className="mr-4">
                  <a href="/nature">Nature</a>
                </li>

                <li className="mr-4">
                  <a href="/photography">Photography</a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default PageContainer;
