<template>
	<div>
		<hmenu></hmenu>

		<div class="main">
			<h1>About</h1>
			
			<div class="box">
				<div class="wrapper">
					Assetto Corsa Server Web Interface<br />
					Version {{version}}<br />
					<span v-html="copyright"></span><br />
					View on <span v-html="github"></span>
				</div>
			</div>

			<div class="box">
				<div class="wrapper">
					<h2>License</h2>

					MIT License<br /><br />

					Copyright (c) 2017 Marvin Blum<br /><br />

					Permission is hereby granted, free of charge, to any person obtaining a copy
					of this software and associated documentation files (the "Software"), to deal
					in the Software without restriction, including without limitation the rights
					to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
					copies of the Software, and to permit persons to whom the Software is
					furnished to do so, subject to the following conditions:

					The above copyright notice and this permission notice shall be included in all
					copies or substantial portions of the Software.<br /><br />

					THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
					IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
					FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
					AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
					LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
					OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
					SOFTWARE.
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import {VERSION, COPYRIGHT, GITHUB_LINK} from "../global.js";
import {hmenu} from "../components";

export default {
	components: {
		hmenu
	},
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
