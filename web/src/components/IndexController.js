import React, {Component} from 'react'
import UploadListings from './UploadListings'



class IndexController extends Component {
    render() {
        return (
            <main>
                <UploadListings />
                {/* <BrowseByCategory /> */}
            </main>
        )
    }
}

export default IndexController