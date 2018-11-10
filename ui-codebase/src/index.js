
import React from 'react'
import { render } from 'react-dom'
import jwt_decode  from 'jwt-decode'
import { BrowserRouter as Router } from 'react-router-dom'
import './index.css'
import Root from './containers/Root'
import configureStore from './store/configureStore'
import {authenticateUser, fetchingUser, fetchingUserSuccess, fetchingUserFailure} from './redux/User'
import api from './modules/api';
import UserProfile from './modules/userProfile';


const store = configureStore()

//Check if the user is authenticated initially
const token = localStorage.getItem('token');
if(token) {
  store.dispatch(authenticateUser());
  store.dispatch(fetchingUser());
  var decoded = jwt_decode(token)
  api.GetUserById(decoded.id)
    .then((response) => {
      store.dispatch(fetchingUserSuccess(response.data.data))
      return UserProfile.setUserData(response.data.data)
    })
    .catch((err) => store.dispatch(fetchingUserFailure(err)))
}
render(
  <Router>
    <Root store={store} />
  </Router>,
  document.getElementById('root')
)