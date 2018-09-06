<template>
	<div class="login">
		<div class="center">
			<img src="static/aclogo.png" alt="" />
		</div>

		<div class="box">
			<div class="wrapper">
				<h1>Login</h1>
				<msg :type="'error'" :msg="'The login and/or password was wrong.'" v-if="err != 0"></msg>

				<form v-on:submit.prevent="performLogin()">
					<table>
						<tr>
							<td class="w30">Login/E-Mail:</td>
							<td><input type="text" name="login" v-model="login" class="full-width" autofocus /></td>
						</tr>
						<tr>
							<td>Password:</td>
							<td><input type="password" name="pwd" v-model="pwd" class="full-width" /></td>
						</tr>
						<tr>
							<td></td>
							<td><input type="submit" value="Login" /></td>
						</tr>
					</table>
				</form>
			</div>
		</div>

		<div class="version">
			<div class="wrapper">
				Version {{version}} | <span v-html="copyright"></span> | <span v-html="github"></span>
			</div>
		</div>
	</div>
</template>

<script>
import axios from "axios";
import {VERSION, COPYRIGHT, GITHUB_LINK} from "../global.js";
import {msg} from "../components";

export default {
	components: {
		msg
	},
	data() {
		return {
			version: VERSION,
			copyright: COPYRIGHT,
			github: GITHUB_LINK,
			err: 0,
			login: '',
			pwd: ''
		}
	},
	methods: {
		performLogin() {
			axios.post('/api/login', {login: this.login, pwd: this.pwd})
			.then(resp => {
				if(resp.data.code){
					this.err = resp.data.code;
					return;
				}

				this.$store.commit("login", resp.data);
				this.$router.push('/instance');
			});
		}
	}
}
</script>

<style lang="scss">
</style>
