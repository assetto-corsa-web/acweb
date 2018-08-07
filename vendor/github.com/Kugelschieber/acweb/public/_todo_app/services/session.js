var SessionService = {
	data: {
		userId: 0
	},
	init: function(data){
		if(!data){
			return;
		}
		
		this.data.userId = data.user_id;
	},
	login: function(router, data){
		this.init(data.data);
		router.push('/instance');
	},
	logout: function(router){
		this.data.userId = 0;

		Vue.http.put('/api/logout')
		.then(function(resp){
			if(resp.data.code){
				console.log(resp.data.code+': '+resp.data.msg);
			}
			
			router.push('/');
		});
	}
};
