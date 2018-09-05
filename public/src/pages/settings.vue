<template>
	<div>
		<hmenu></hmenu>

		<div class="main">
			<h1>Settings</h1>
			
			<div class="box">
				<div class="wrapper">
					<h2>Server Settings</h2>

					<msg :type="'success'" :msg="'The settings have been saved.'" v-if="saved"></msg>
					<msg :type="'error'" :msg="'The AC folder and the executable must be set.'" v-if="err == 1"></msg>
					<msg :type="'error'" :msg="'You have no permission to do this.'" v-if="err == 200"></msg>

					<form v-on:submit.prevent="performSave()">
						<table>
							<tr>
								<td class="w10">AC server folder:</td>
								<td><input type="text" name="path" class="full-width" v-model="folder" /></td>
							</tr>
							<tr>
								<td>Executable:</td>
								<td><input type="text" name="executable" class="full-width" v-model="executable" /></td>
							</tr>
							<tr>
								<td>Arguments:</td>
								<td><input type="text" name="args" class="full-width" v-model="args" /></td>
							</tr>
							<tr>
								<td></td>
								<td><input type="submit" value="Save" /></td>
							</tr>
						</table>
					</form>
				</div>
			</div>

			<div class="box">
				<div class="wrapper">
					The <strong>AC server folder</strong> must be the full path to your AC server installation folder, containing the acServer executable. Example: /home/acuser/steam/steamapps/common/Assetto Corsa Dedicated Server<br />
					<strong>Executable</strong> is the executable file to start a server instance. Example: acServer<br />
					<strong>Arguments</strong> are the arguments passed to the executable to start a server instance, separated by spaces. Example: linux<br /><br />

					Make sure the web interface has the permissions to read, write and execute within the server folder. If you cannot start instances, please test your settings on your server.
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import axios from "axios";
import {hmenu} from "../components";

export default {
	components: {
		hmenu
	},
	data: function(){
		return {
			err: 0,
			folder: '',
			executable: '',
			args: '',
			saved: false
		}
	},
	mounted: function(){
		this._load();
	},
	methods: {
		_load: function(){
			axios.get('/api/settings')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.folder = resp.data.folder;
				this.executable = resp.data.executable;
				this.args = resp.data.args;
			});
		},
		performSave: function(){
			this.saved = false;

			axios.post('/api/settings', {folder: this.folder,
				executable: this.executable,
				args: this.args})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this.saved = true;
			});
		}
	}
};
</script>

<style lang="scss">
</style>
