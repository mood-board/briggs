import { createStore, applyMiddleware, compose, combineReducers } from 'redux'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'
import DevTools from '../containers/DevTools'


import user from '../redux/User'


const configureStore = preloadedState => {
  const store = createStore(
    combineReducers({
      user
    }),
    preloadedState,
    compose(
      applyMiddleware(thunk, createLogger()),
      DevTools.instrument()
    )
  )
  return store
}

export default configureStore