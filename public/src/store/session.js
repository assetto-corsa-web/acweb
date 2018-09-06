import axios from "axios";
import Vuex from "vuex";

export default function NewSessionStore(){
	return new Vuex.Store({
		state: {
			userId: parseInt(window.localStorage.getItem("user_id"))
		},
		mutations: {
			login(state, data) {
				state.userId = data.user_id;
				window.localStorage.setItem("user_id", data.user_id)
			},
			logout(state) {
				state.userId = 0;
				window.localStorage.removeItem("user_id");

				axios.put('/api/logout')
				.then(resp => {
					if(resp.data.code){
						console.error(resp.data.code+': '+resp.data.msg);
					}
				});
			}
		}
	});
};
