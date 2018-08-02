import React, {Component} from "react"
import axios from "axios"
import "./PhotoGallery.css"
import "../api/free-images.json"


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
        let freeImageListing = this.state.photos.map(item => {
            return (
                <a className="grid_item" href={`/photos/${item.slug}`}>
                    <img alt="" src={item.url} />
                </a>
            )
        })

        return(
            <div className="freestock_lp container">
                <h3 className="py-4 mb-4">Free Stock Photos</h3>
                <div className="grid">

                    {freeImageListing}
                    {/* <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>

                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/1275310/pexels-photo-1275310.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>

                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/1280730/pexels-photo-1280730.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>

                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/1278566/pexels-photo-1278566.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>

                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/439862/pexels-photo-439862.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/1261408/pexels-photo-1261408.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545004/pexels-photo-545004.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/1278620/pexels-photo-1278620.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div>
                    <div className="grid_item">
                        <img alt="" src="https://images.pexels.com/photos/545063/pexels-photo-545063.jpeg?auto=compress&cs=tinysrgb&h=350" />
                    </div> */}
                </div>
            </div>
        )
    }
}

export default PhotoGallery;
