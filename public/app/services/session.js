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
		router.push('/dashboard');
	},
	logout: function(router){
		
	}
};
