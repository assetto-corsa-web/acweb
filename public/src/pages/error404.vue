<template>
	<div class="login">
		<div class="center">
			<img src="static/aclogo.png" alt="" />
		</div>

		<div class="box">
			<div class="wrapper">
				<h1>Error #404</h1>
				<p>The page was not found.</p>
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
import {VERSION, COPYRIGHT, GITHUB_LINK} from "../global.js";

export default {
	data() {
		return {
			version: VERSION,
			copyright: COPYRIGHT,
			github: GITHUB_LINK
		}
	}
}
</script>

<style lang="scss">
</style>
