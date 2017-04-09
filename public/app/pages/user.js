Vue.component('User', {
	template: '{{include=template/pages/user.html}}',
	data: function(){
		return {
			user: [],
			_id: 0,
			login: '',
			email: '',
			pwd1: '',
			pwd2: '',
			admin: false,
			moderator: false,
			err: 0,
			addEditUser: false,
			removeUser: false,
			userSaved: false,
			userRemoved: false
		}
	},
	mounted: function(){
		this._load();
	},
	methods: {
		_load: function(){
			this.$http.get('/api/getAllUser')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.user = resp.data;
			});
		},
		_reset: function(){
			this._id = 0;
			this.login = '';
			this.email = '';
			this.pwd1 = '';
			this.pwd2 = '';
			this.admin = false;
			this.moderator = false;
			this.err = 0;
			this.addEditUser = false;
			this.removeUser = false;
			this.userSaved = false;
			this.userRemoved = false;
		},
		openAddEditUser: function(id){
			this._reset();
			
			if(id){
				this._id = id;

				this.$http.get('/api/getUser', {params: {id: id}})
				.then(function(resp){
					if(resp.data.code){
						console.log(resp.data.code+': '+resp.data.msg);
						return;
					}

					this.login = resp.data.login;
					this.email = resp.data.email;
					this.admin = resp.data.admin;
					this.moderator = resp.data.moderator;
					this.addEditUser = true;
				});
			}
			else{
				this.addEditUser = true;
			}
		},
		openRemoveUser: function(id){
			this._reset();

			if(!id){
				return;
			}

			this._id = id;
			this.removeUser = true;
		},
		performAddEditUser: function(){
			this.$http.post('/api/addEditUser', {id: this._id,
				login: this.login,
				email: this.email,
				pwd1: this.pwd1,
				pwd2: this.pwd2,
				admin: this.admin,
				moderator: this.moderator})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.userSaved = true;
			});
		},
		performRemoveUser: function(){
			this.$http.post('/api/removeUser', {id: this._id})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this._reset();
				this._load();
				this.userRemoved = true;
			});
		}
	}
});
