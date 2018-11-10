import React from "react"
import { connect } from "react-redux"

import "./Users.css"

class Settings extends React.Component {
    constructor(){
        super();
        this.onFieldChange = this.onFieldChange.bind(this)
        this.onSubmit = this.onFieldChange.bind(this)
    }
    
    onFieldChange (event) {
        this.setState({[event.target.name]: event.target.value})
    }

    onSubmit (event) {
        event.preventDefault()
        console.log(document.getElementsByName("username").value)
    }

    render() {
        if(!this.props.user) {
            return <div>Loading...</div>
        }
        return (
            <section className="user_settings bg-grey-lighter">
                <div className="personal_data font-sans text-sm rounded w-full max-w-md mx-auto px-8 pb-8">
                    <h4>Personal data</h4><hr />
                    
                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="username">Username</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="username" value={this.props.user.username} 
                        onChange={this.onFieldChange} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="profile_image"> Profile Image</label>
                        <img src="https://via.placeholder.com/50x50" />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label>Please Select</label>
                        <select className="w-full py-2 px-3 text-grey-darker leading-normal rounded">
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                        </select>
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="first_name">First name</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="first_name" value={this.props.user.first_name} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="last_name">Last name</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="last_name" value={this.props.user.last_name} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="city">City</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="city" value={this.props.user.city} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="country">Country</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="country" value={this.props.user.country} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label>Date of Birth</label>
                        <span>
                            <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                            type="text" name="date" onChange={this.onFieldChange.bind(this)} />
                            <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                            type="text" name="month" onChange={this.onFieldChange.bind(this)} />
                            <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                            type="number" name="year" onChange={this.onFieldChange.bind(this)} />
                        </span>
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="about_me">About Me</label>
                        <textarea className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        onChange={this.onFieldChange.bind(this)}>
                           {this.props.user.about_me}
                        </textarea>
                    </div>

                </div>

                <div className="other_profiles font-sans text-sm rounded w-full max-w-md mx-auto px-8 pb-8">
                    <h4>Online Profiles</h4><hr />
                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="profile_facebook">Facebook</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="profile_facebook" value={this.props.user.online_profile.profile_facebook} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>
                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="profile_instagram">Instagram</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="profile_instagram" value={this.props.user.online_profile.profile_instagram} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>
                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="profile_twitter">Twitter</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="profile_twitter" value={this.props.user.online_profile.profile_twitter} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label htmlFor="profile_website">Website</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" 
                        type="text" name="profile_website" value={this.props.user.online_profile.profile_website} 
                        onChange={this.onFieldChange.bind(this)} />
                    </div>

                    <div className="relative border rounded mb-4 shadow appearance-none label-floating">
                        <label>Email</label>
                        <input className="w-full py-2 px-3 text-grey-darker leading-normal rounded" type="email" 
                        value={this.props.user.email} disabled />
                    </div>
                </div>

                <button className="bg-black hover:bg-black text-white py-2 px-4 mx-auto text-center mb-8 w-1/2" 
                onClick={this.onSubmit.bind(this)}>Save</button>
            </section>
        )
    }
}

function mapStateToProps({user}){
    let userObject = user.toJS()
    return {
        user :userObject.authedUser
    }
}

export default connect(mapStateToProps, null)(Settings);