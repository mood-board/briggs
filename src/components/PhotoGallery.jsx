import React, {Component} from "react"
import axios from "axios"
import {Link} from "react-router-dom"
import "./PhotoGallery.css"


class PhotoGallery extends Component {
    constructor(){
        super();

        this.state = {
            photos: []
        }
    }
    componentWillMount() {
       let freeImagesURL = "https://api.myjson.com/bins/1e2fdc"
       axios.get(freeImagesURL)
        .then(res => this.setState({ photos: res.data}))
    }

    render() {
        if(!this.state.photos) {
            return(
                <div>Loading...</div>
            )
        }
        let freeImageListing = this.state.photos.map(item => {
            return (
                <Link key={item.slug} className="grid_item" to={`/photos/${item.slug}`}>
                    <img alt="" src={item.url} className="rounded-sm" />
                </Link>
            )
        })

        return(
            <div className="freestock_lp container">
                <h3 className="py-4 mb-4">Free Stock Photos</h3>
                <div className="grid">
                    {freeImageListing}
                </div>
            </div>
        )
    }
}

export default PhotoGallery;
