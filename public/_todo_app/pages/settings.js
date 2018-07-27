Vue.component('Settings', {
	template: '{{include=template/pages/settings.html}}',
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
			this.$http.get('/api/settings')
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

			this.$http.post('/api/settings', {folder: this.folder,
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
});
