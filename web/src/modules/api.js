import axios from "axios"

let API_URL = "http://localhost:4000/api"

let api = {
    Login : (form) => {
        return axios.post(`${API_URL}/users/authenticate/login`, form, { headers: {'Accept': 'application/json'} })
    },
    Register: (credentials) => {
        return axios.post(`${API_URL}/users/authenticate/signup`, credentials, { headers: {'Accept': 'application/json'} })
    },
    GetUserById: (user_id) => {
        return axios.get(`${API_URL}/users/id/${user_id}`)
    }
}


export default api 