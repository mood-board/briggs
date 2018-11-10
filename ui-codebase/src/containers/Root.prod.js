import React from 'react'
import PropTypes from 'prop-types'
import { Provider } from 'react-redux'
import { Router, Route } from 'react-router-dom'
import Register from '../components/Register';
import Login from '../components/Login';
import history from '../history';
import IndexController from '../components/IndexController'


const Root = ({ store }) => (
    <Provider store={store}>
        <Router history={history}>
            <div>
                <Route exact path="/" component={IndexController} />    
                <Route path="/auth/signup" component={Register} />
                <Route path="/auth/login" component={Login} />
                <Route path="/accounts/settings" component={Settings} />
            </div>
        </Router>
    </Provider>
  )
  
  Root.propTypes = {
    store: PropTypes.object.isRequired,
  }
  export default Root