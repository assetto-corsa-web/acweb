Vue.component('Instance', {
	template: '{{include=template/pages/instance.html}}',
	data: function(){
		return {
			_pid: 0,
			instances: [],
			configs: [],
			err: 0,
			name: '',
			config: 0,
			started: false,
			stopped: false,
			startInstance: false,
			stopInstance: false
		}
	},
	mounted: function(){
		this._load();
	},
	methods: {
		_load: function(){
			this.$http.get('/api/getAllConfigurations')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.configs = resp.data;
				this._loadInstances();
			});
		},
		_loadInstances: function(){
			this.$http.get('/api/getAllInstances')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.instances = resp.data;

				for(var i = 0; i < this.instances.length; i++){
					this.instances[i].configuration = this._getConfigName(this.instances[i].configuration);
				}
			});
		},
		_getConfigName: function(id){
			for(i in this.configs){
				if(id == this.configs[i].id){
					return this.configs[i].name;
				}
			}

			return '';
		},
		_reset: function(){
			this.err = 0;
			this.name = '';
			this.config = 0;
			this.started = false;
			this.stopped = false;
			this.startInstance = false;
			this.stopInstance = false;
		},
		performStart: function(){
			this.$http.post('/api/startInstance', {name: this.name, config: this.config})
			.then(function(resp){
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
		performStop: function(){
			this.$http.post('/api/stopInstance', {pid: this._pid})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.stopped = true;
			});
		}
	}
});
