<template>
	<div>
		<hmenu></hmenu>

		<div class="main">
			<h1>Status</h1>

			<div class="box" v-if="startInstance">
				<div class="wrapper">
					<h2>Start New Instance</h2>

					<msg :type="'error'" :msg="'Error starting server instance.'" v-if="err != 0 && err != 200" v-on:close="closeMsg"></msg>
					<msg :type="'error'" :msg="'You have no permission to do this.'" v-if="err == 200" v-on:close="closeMsg"></msg>

					<form v-on:submit.prevent="performStart()">
						<table>
							<tr>
								<td class="w20">Name:</td>
								<td><input type="text" name="name" class="full-width" v-model="name" /></td>
							</tr>
							<tr>
								<td>Configuration:</td>
								<td>
									<select name="configuration" v-model="config">
										<option v-for="c in configs" v-bind:value="c.id">{{c.name}}</option>
									</select>
								</td>
							</tr>
							<tr>
								<td>Run script/executable before start:</td>
								<td>
									<input type="text" name="script_before" class="full-width" v-model="script_before">
									<br />
									(full path, or leave empty)
								</td>
							</tr>
							<tr>
								<td>Run script/executable after start:</td>
								<td>
									<input type="text" name="script_after" class="full-width" v-model="script_after">
									<br />
									(full path, or leave empty)
								</td>
							</tr>
							<tr>
								<td></td>
								<td>
									<input type="submit" value="Start" />
									<button v-on:click="startInstance = false">Cancel</button>
								</td>
							</tr>
						</table>
					</form>
				</div>
			</div>

			<div class="box" v-if="stopInstance">
				<div class="wrapper">
					<h2>Stop Instance</h2>
					<msg :type="'error'" :msg="'You have no permission to do this.'" v-if="err == 200" v-on:close="closeMsg"></msg>
					<p>Do you really want to stop this server instance?</p>
					<button v-on:click="performStop()">Yes, stop instance</button>
					<button v-on:click="stopInstance = false">Cancel</button>
				</div>
			</div>

			<div class="box" v-if="showLog">
				<div class="wrapper">
					<h2>Log Output</h2>
					<button v-on:click="refreshLog">Refresh</button>
					<small>(showing the last 256kb)</small>
					<textarea v-model="log" style="min-height: 300px;"></textarea>
					<button v-on:click="showLog = false">Close</button>
				</div>
			</div>

			<div class="box" v-if="showDeleteLog">
				<div class="wrapper">
					<h2>Delete Log File</h2>
					<p>Do you really want to delete the log file "{{activeLogFilename}}"?</p>
					<button v-on:click="deleteLogfile">Yes, delete log file</button>
					<button v-on:click="showDeleteLog = false">Close</button>
				</div>
			</div>

			<div class="box" v-if="showDeleteAllLogs">
				<div class="wrapper">
					<h2>Delete All Logs</h2>
					<p>Do you really want to delete all log files?</p>
					<button v-on:click="deleteAllLogs">Yes, delete all logs</button>
					<button v-on:click="showDeleteAllLogs = false">Close</button>
				</div>
			</div>

			<div class="box">
				<div class="wrapper">
					<h2>Active Server Instances</h2>

					<msg :type="'success'" :msg="'The server instance has been started.'" v-if="started" v-on:close="started = false"></msg>
					<msg :type="'success'" :msg="'The server instance has been stopped.'" v-if="stopped" v-on:close="stopped = false"></msg>

					<button v-on:click="startInstance = true">Start New Instance</button>

					<table>
						<thead>
							<tr>
								<td class="w5">PID</td>
								<td class="w40">Instance</td>
								<td class="w40">Configuration</td>
								<td class="w15"></td>
							</tr>
						</thead>
						<tbody>
							<tr v-for="instance in instances">
								<td>{{instance.pid}}</td>
								<td>{{instance.name}}</td>
								<td>
									{{instance.configuration.name}}<br />
									{{instance.configuration.track}}<br />
									TCP: {{instance.configuration.tcp}} UDP: {{instance.configuration.udp}}
								</td>
								<td>
									<a class="icon-link" v-bind:href="generateInstanceDownloadUrl(instance.configuration.id)">
										<i class="fa fa-cloud-download" aria-hidden="true" title="Download instance configuration"></i>
									</a>
									<i class="fa fa-stop" aria-hidden="true" title="Stop instance" v-on:click="showStopInstance(instance.pid)"></i>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>

			<div class="box">
				<div class="wrapper">
					<h2>Log Files</h2>

					<button v-on:click="showDeleteAllLogs = true">Delete All Logs</button>

					<table>
						<thead>
							<tr>
								<td class="w55">Filename</td>
								<td class="w15">Date</td>
								<td class="w15">Size</td>
								<td class="w15"></td>
							</tr>
						</thead>
						<tbody>
							<tr v-for="log in logs">
								<td>{{log.file}}</td>
								<td>{{log.date}}</td>
								<td>{{log.size}}</td>
								<td>
									<a class="icon-link" v-bind:href="generateConsoleLogDownloadUrl(log.file)">
										<i class="fa fa-cloud-download" aria-hidden="true" title="Download console output"></i>
									</a>
									<i class="fa fa-terminal" aria-hidden="true" title="Show console output" v-on:click="openLog(log.file)"></i>
									<i class="fa fa-trash" aria-hidden="true" title="Delete log file" v-on:click="openDeleteLog(log.file)"></i>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import axios from "axios";
import {hmenu, msg} from "../components";

export default {
	components: {
		hmenu,
		msg
	},
	data() {
		return {
			_pid: 0,
			instances: [],
			configs: [],
			logs: [],
			err: 0,
			name: '',
			config: 0,
			script_before: '',
			script_after: '',
			log: '',
			activeLogFilename: '',
			showLog: false,
			showDeleteLog: false,
			showDeleteAllLogs: false,
			started: false,
			stopped: false,
			startInstance: false,
			stopInstance: false,
			logFile: ''
		}
	},
	mounted() {
		this._load();
	},
	methods: {
		_load() {
			this._loadConfiguration();
			this._loadLogs();
		},
		_loadConfiguration() {
			axios.get('/api/configuration')
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.configs = resp.data;
				this._loadInstances();
			});
		},
		_loadLogs() {
			axios.get('/api/instance/log')
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.logs = resp.data;

				// reverse
				for(let i = this.logs.length-1; i >= 0; i--){
					this.logs[i].date = new Date(this.logs[i].date).toString();

					if(this.logs[i].size > 1024*1024){
						this.logs[i].size = Math.round(this.logs[i].size/1024/1024*100)/100+' MB';
					}
					else if(this.logs[i].size > 1024){
						this.logs[i].size = Math.round(this.logs[i].size/1024*100)/100+' KB';
					}
					else{
						this.logs[i].size = Math.round(this.logs[i].size*100)/100+' Byte';
					}
				}
			});
		},
		_loadInstances() {
			axios.get('/api/instance')
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.instances = resp.data;

				for(let i = 0; i < this.instances.length; i++){
					this.instances[i].configuration = this._getConfigName(this.instances[i].configuration);
				}
			});
		},
		_getConfigName: function(id){
			for(let i = 0; i < this.configs.length; i++){
				if(id == this.configs[i].id){
					return this.configs[i];
				}
			}

			return null;
		},
		_reset() {
			this.err = 0;
			this.name = '';
			this.config = 0;
			this.log = '';
			this.showLog = false;
			this.started = false;
			this.stopped = false;
			this.startInstance = false;
			this.stopInstance = false;
			this.logFile = '';
		},
		performStart() {
			let config = {
				name: this.name,
				config: this.config,
				script_before: this.script_before,
				script_after: this.script_after
			};
			
			axios.post('/api/instance', config)
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.started = true;
			});
		},
		showStopInstance: function(pid){
			this._reset();

			if(!pid){
				return;
			}

			this._pid = pid;
			this.stopInstance = true;
		},
		performStop() {
			axios.delete('/api/instance', {params: {pid: this._pid}})
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.stopped = true;
			});
		},
		openLog(file) {
			this._reset();
			this.logFile = file;

			axios.get('/api/instance/log', {params: {file}})
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				var log = resp.data.substr(1, resp.data.length-2).split('\\n');
				this.log = log.join('\n');
				this.showLog = true;
			});
		},
		generateConsoleLogDownloadUrl(file) {
			return '/api/instance/log?file=' + file + '&dl=1';
		},
		generateInstanceDownloadUrl(id) {
			return '/api/configuration?id=' + id + '&dl=2';
		},
		openDeleteLog(filename) {
			this.activeLogFilename = filename;
			this.showDeleteLog = true;
		},
		deleteLogfile() {
			axios.delete('/api/instance/log', {params: {filename: this.activeLogFilename}})
			.then(resp => {
				this.activeLogFilename = '';
				this.showDeleteLog = false;
				this._loadLogs();
			});
		},
		deleteAllLogs() {
			axios.delete('/api/instance/log')
			.then(resp => {
				this.showDeleteAllLogs = false;
				this._loadLogs();
			});
		},
		closeMsg() {
			this.err = 0;
		},
		refreshLog() {
			axios.get('/api/instance/log', {params: {file: this.logFile}})
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				var log = resp.data.substr(1, resp.data.length-2).split('\\n');
				this.log = log.join('\n');
			});
		}
	}
}
</script>

<style lang="scss">
</style>
