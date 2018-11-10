import React from 'react'
import {Link} from 'react-router-dom'


const NavBar = () => {
    return (
        <nav className="font-sans bg-white text-center flex justify-between my-4 mx-auto container overflow-hidden">
            <a href="/" className="block text-left">
                <img src="https://stitches-cdn.hyperyolo.com/logo.png" className="h-10 sm:h-10 rounded-full" alt="logo" />
            </a>
            <ul className="text-sm text-grey-dark list-reset flex items-center">
                <li>
                    <a href="#" className="inline-block py-2 px-3 text-grey-darkest hover:text-grey-dark no-underline">Products</a>
                    </li>
                <li>
                    <a href="/uploads" className="inline-block py-2 px-3 text-grey-darkest hover:text-grey-dark no-underline">Contribute Photo</a>
                </li>
                <li className="pr-2">
                    <Link className="inline-block py-2 px-3 text-grey-darkest hover:text-grey-dark no-underline" to="/accounts/settings">Settings</Link>
                </li>
                <li className="pl-2 border-l">
                    <Link 
                        to="/auth/login" 
                        className="inline-block py-2 px-3 text-grey-darkest hover:text-grey-dark no-underline">Log In
                    </Link>
                    
                </li>
                <Link
                    to="/auth/signup" 
                    className="bg-black hover:bg-grey-darker text-white ml-4 py-2 px-3 no-underline">Sign Up
                </Link>
            </ul>
        </nav>
    )
}

export default NavBar