<template>
	<div v-bind:class="type">
		{{msg}}
	</div>
</template>

<script>
export default {
	props: ['type', 'msg']
};
</script>

<style lang="scss">
</style>
