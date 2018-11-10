import React from 'react'
import "./Listings.css"
import axios from 'axios'


class UploadListings extends React.Component {
    constructor() {
        super()
        this.state = { imageURLS: [] }
    }
    componentWillMount(){
        let URL = "https://api.myjson.com/bins/1bb6as"
        axios.get(URL)
            .then((item) => this.setState({imageURLS: item.data}));
    }
    render() {
        return (
            <div className="grid-container">
                <div className="content-grid" style={{ height: "648px", overflow: "hidden scroll"}}>
                    <div>
                        <div className="results-row">
                            {this.state.imageURLS.map((item) => 
                                <div key={item.largeImageURL} className="item-wrap" style={{width: "155px"}}>
                                    <div className="item-container">
                                        <div className="infinite-list-item grid-item stockfood" style={{height: "241px"}}>
                                            {/* <span><a href="#" style={{backgroundImage: `url(${item.largeImageURL})`}}></a></span> */}
                                            <span><a href="#" style={{backgroundImage: "url(https://images.pexels.com/photos/1036623/pexels-photo-1036623.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260)"}}></a></span>
                                            <div className="source"><i className="source-stockfood"></i></div>
                                            <div className="lightbox-manager">
                                                <div className="basic-button"><button>Login to save image</button></div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            </div>

        )
    }
}

export default UploadListings