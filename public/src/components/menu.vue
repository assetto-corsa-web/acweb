<template>
	<div class="menu">
		<router-link to="/">Home</router-link>
		<router-link to="/notfound">404</router-link>
	</div>
</template>

<script>
export default {
	
}
</script>

<style lang="scss">
.menu{
	padding: 10px;
	background: #000;
	margin: 0 0 20px 0;

	a{
		color: #fff;
		margin: 0 10px 0 0;
	}
}
</style>
