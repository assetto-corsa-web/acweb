<template>
	<div v-bind:class="type">
		{{msg}}
		<div class="msg-close" v-on:click="close">
			<i class="fa fa-times" aria-hidden="true"></i>
		</div>
	</div>
</template>

<script>
export default {
	props: ['type', 'msg'],
	methods: {
		close() {
			this.$emit("close");
		}
	}
};
</script>

<style lang="scss">
.msg-close{
	width: 24px;
	height: 24px;
	float: right;
	display: flex;
	justify-content: center;
	align-items: center;
}

.msg-close:hover{
	cursor: pointer;
}
</style>
