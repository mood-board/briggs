
import { fromJS } from 'immutable';
import api from '../modules/api'
import history from '../history'
import UserProfile from '../modules/userProfile';

export const AUTH_USER = "AUTH_USER"
export const UNAUTH_USER = "UNAUTH_USER"
export const AUTH_ERROR = "AUTH_ERROR"
export const FETCHING_USER = "FETCHING _USER"
export const FETCHING_USER_FAILURE = "FETCHING_USER_FAILURE"
export const FETCHING_USER_SUCCESS = "FETCHING_USER_SUCCESS"


//Actions creators
export const  authenticateUser = (user) => ({  
	type: AUTH_USER,
	user
});

export const authenticationError = (error) => ({
	type: AUTH_ERROR,
	error
});

export const fetchingUser = () => ({ 
  type: FETCHING_USER
});

export const fetchingUserFailure = (error) => ({  
	type: FETCHING_USER_FAILURE,
	error: 'Error fetching user'
})

export const fetchingUserSuccess = (user) => ({
	type: FETCHING_USER_SUCCESS,
	user
});

export const unauthUser = ()=>({
	type: UNAUTH_USER
});


const initialState = fromJS({
	isFetching: false,
	error: '',
	isAuthenticated: false,
	authedUser: {}
});


//Reducers 

export default function usersReducer ( state = initialState, action) {
	switch( action.type) {
		case AUTH_USER: 
			return state.merge({
				isAuthenticated: true,
				authedUser: action.user
			});
		case UNAUTH_USER:
			return state.merge({
				isAuthenticated: false,
				authedUser: {}
			});
		case AUTH_ERROR: 
			return state.merge({
				error: action.error
			});
		case FETCHING_USER:
			return state.merge({
				isFetching: true
			})
		case FETCHING_USER_FAILURE:
			return state.merge({
				isFetching: false,
				isAuthenticated: false,
				error: action.error
			});
		
		case FETCHING_USER_SUCCESS:
			return action.user == null 
				? state.merge({
					isFetching: false,
					error: ''
				})
				: state.merge({
					isFetching: false,
					error: '',
					authedUser: action.user
				})
		default:
				return state
	}
}

export function logoutUser() {
	return dispatch => {
		localStorage.removeItem('token')
		dispatch(unauthUser());
	};
}

export function login (form) {
	return dispatch => api.Login(form)
		.then((response) => {
			dispatch(authenticateUser(response.data.user))
			localStorage.setItem('token', response.data.token)
			UserProfile.setUserData(response.data.user)
			dispatch(fetchingUserSuccess(response.data.user))
			return history.push('/') //GOTO: Homepage
		})
		.catch((err) => {
			let error = err
			return dispatch(authenticationError(error));
		})
}

export function register(credentials) {
	return dispatch => api.Register(credentials)
		.then((response) => {
			dispatch(authenticateUser(response.data.user))
			localStorage.setItem("token", response.data.token)
			UserProfile.setUserData(response.data.user)
			dispatch(fetchingUserSuccess(response.data.user))
			return history.push('/')
		})
		.catch((err) => {
			let error = err
			return dispatch(authenticationError(error));
		})
}