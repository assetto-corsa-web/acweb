<template>
	<div class="menu">
		<div class="wrapper">
			<div class="logo"><img src="static/aclogo_small.png" alt="" /></div>
			<router-link to="/instance">Instances</router-link>
			<router-link to="/configuration">Configurations</router-link>
			<router-link to="/settings">Settings</router-link>
			<router-link to="/user">User</router-link>
			<router-link to="/about">About</router-link>

			<a href="#" class="right" v-on:click="performLogout()"><i class="fa fa-sign-out" aria-hidden="true" title="Logout"></i></a>
		</div>
	</div>
</template>

<script>
import * as global from "../global.js";

export default {
	data() {
		return {
			version: global.VERSION
		}
	},
	methods: {
		performLogout() {
			this.$store.commit("logout");
			this.$router.push('/');
		}
	}
};
</script>

<style lang="scss">
</style>
