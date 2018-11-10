import React from 'react'
import axios from 'axios'


export default class PhotoUpload extends React.Component {
    state = {
        selectedFile: null,
        tags: '',
        fileSize: null,
        uploadError: null
    }
 
    fileSelectedHandler = event => { 
        console.dir(event.target.files[0].size);
        
        this.setState({
            selectedFile: event.target.files[0],
            fileSize: event.target.files[0].size
        });
    }

    updateTags = event => {
        console.dir(event.target.value)
        this.setState({tags: event.target.value})
    }

    fileUploadHandler = () => {
        const formData = new FormData(); 
        formData.append('uploads', this.state.selectedFile, this.state.selectedFile.name);
        formData.append("user_id", JSON.parse(localStorage.getItem("userData")).user_id);
        formData.append("tags", this.state.tags);
        formData.append("size", this.state.fileSize);
        
        axios.post("http://localhost:4000/api/uploads", formData, {
            onUploadProgress: progressEvent => {
                //TODO: Show a progress loader to signify upload interval
                console.log('Upload progress ' + Math.round(progressEvent.loaded / progressEvent.total * 100)) 
            }
        })
        .then((item) => console.dir(item)) //TODO: Redirect to homepage for now
        .catch((error) => this.setState({uploadError: "An Error occurred"}))
    }


    render() {
        return (
            <div>  
                <div 
                    className="font-sans text-sm rounded w-full max-w-md mx-auto my-8 px-8 pt-6 pb-8">
                    {this.state.uploadError && (
                        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                            {this.state.uploadError}
                        </div>
                    )}
                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="file">Filenames:</label>
                        <input 
                            style={{display: 'none'}} 
                            ref={fileInput => this.fileInput = fileInput} 
                            onChange={this.fileSelectedHandler} 
                            className="w-full py-2 px-3 text-grey-darker leading-normal rounded" type="file" name="uploads" />
                        <img
                            onClick={() => this.fileInput.click()}
                            src="fileupload.gif" />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <input onChange={this.updateTags} className="w-full py-2 px-3 text-grey-darker leading-normal rounded" placeholder="Tags comma separated" type="text" name="tags" id="tags" />
                    </div>    
                    <button onClick={this.fileUploadHandler} className="bg-black hover:bg-black text-white py-2 px-4">Upload</button>
                </div>
            </div>
        )
    }
}
