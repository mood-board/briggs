import React from 'react'


const SiteNavigation = () => {
  return (
    <div className="flex py-2 pin-t pin-x z-100 h-16 items-center transparent">
      <div className="container">
      <div className="w-full max-w-screen-xl relative mx-auto px-6">
        <div className="flex items-center -mx-6">
          <div className="lg:w-1/4 xl:w-1/5 pl-6 pr-6 lg:pr-8">
            <div className="flex items-center">
              <a href="/" className="hidden brand lg:block font-bold text-lg text-white no-underline">
                Seemars
              </a>
            </div>
          </div>

          <div className="flex flex-grow items-center lg:w-3/4 xl:w-4/5">
            <div className="w-full lg:px-6 lg:w-3/4 xl:px-12">
              <div className="relative">
                <span className="algolia-autocomplete">
                  <input id="docsearch" className="transition focus:outline-0 border border-transparent
                    focus:bg-white focus:border-grey-light placeholder-grey-darkest rounded bg-grey-lighter
                    py-2 pr-4 pl-10 block w-full appearance-none leading-normal ds-input" type="text"
                    placeholder="Search the docs" autoComplete="off" spellCheck="false" role="combobox"
                    aria-autocomplete="list" aria-expanded="false" aria-owns="algolia-autocomplete-listbox-0"
                    dir="auto" />
                    <pre aria-hidden="true"></pre><span className="ds-dropdown-menu" role="listbox" id="algolia-autocomplete-listbox-0"><div className="ds-dataset-1"></div></span></span>
                <div className="pointer-events-none absolute pin-y pin-l pl-3 flex items-center">
                  <svg className="fill-current pointer-events-none text-grey-dark w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M12.9 14.32a8 8 0 1 1 1.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12z"></path></svg>
                </div>
              </div>
            </div>
          </div>

          <div className="w-3/5">
            <ul className="flex list-reset nav-listings">
              <li className="mr-12">
                <a href="#" className="text-white no-underline">Collections</a>
              </li>
              <li className="mr-4">
                <a href="#" className="text-white no-underline">All Photos</a>
              </li>

              <li className="mr-4">
                <a className="bg-blue text-white p-4 no-underline" href='#'>Contribute Photos</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>

  )
}

export default SiteNavigation;
