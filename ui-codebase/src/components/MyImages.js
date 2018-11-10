import React from 'react'
import axios from 'axios'
import {Link} from 'react-router-dom'

class MyImages extends React.Component {
    constructor() {
        super()
        this.state = {
            myImages: []
        }
    }

    componentDidMount() {
        let userID = JSON.parse(localStorage.getItem('userData')).user_id;
        let URL = `http://localhost:4000/api/uploads/media/${userID}`
        axios.get(URL).then((item) => this.setState({myImages: item.data.data}))
    }
    render() {
        const hasUploads = !this.state.myImages || this.state.myImages.length === 0;

        return(
            <div className="container mx-auto">
                <div className="content-grid">
                    <div className="results-row">
                        <h2>My Uploads</h2>
                        { hasUploads ? (
                            <div className="no-uploads">
                                <p>That is an empty slate</p>
                                <p>Increase your presence by uploading images <Link to="/uploads">here</Link></p>
                            </div>
                        ) : (
                            <div style={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)"}}>
                                {this.state.myImages.map((item) => 
                                <div key={item.upload_url} className="item-wrap">
                                    <div className="item-container">
                                        <div className="infinite-list-item grid-item stockfood" style={{height: "250px"}}>
                                            <span><a href="#" style={{backgroundImage: `url(${item.upload_url})`}}></a></span>
                                            <div className="source"><i className="source-stockfood"></i></div>
                                            <div className="lightbox-manager">
                                                <div className="basic-button"><button>Login to save image</button></div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                )}
                            </div>
                        )}
                    </div>
                </div>
            </div>
        )
    }
}


export default MyImages
