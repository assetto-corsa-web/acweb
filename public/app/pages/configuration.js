Vue.component('Configuration', {
	template: '{{include=template/pages/configuration.html}}',
	data: function(){
		return {
			configs: [],
			tracks: [],
			cars: [],
			_id: 0,
			// ---
			name: '',
			pwd: '',
			admin_pwd: '',
			pickup_mode: false,
			race_overtime: 0,
			max_slots: 0,
			description: '',
			udp: 0,
			tcp: 0,
			http: 0,
			packets_hz: 0,
			loop_mode: false,
			show_in_lobby: false,
			abs: '',
			tc: '',
			stability_aid: false,
			auto_clutch: false,
			tyre_blankets: false,
			force_virtual_mirror: false,
			fuel_rate: 0,
			damage_rate: 0,
			tires_wear_rate: 0,
			allowed_tires_out: 0,
			max_ballast: 0,
			dynamic_track: false,
			condition: '',
			start_value: 0,
			randomness: 0,
			transferred_grip: 0,
			laps_to_improve_grip: 0,
			kick_vote_quorum: 0,
			session_vote_quorum: 0,
			vote_duration: 0,
			blacklist: '',
			booking: false,
			booking_time: 0,
			practice: false,
			practice_time: 0,
			can_join_practice: false,
			qualify: false,
			qualify_time: 0,
			can_join_qualify: false,
			race: false,
			race_time: 0,
			race_wait_time: 0,
			join_type: '',
			time: '',
			track: '',
			// ---
			err: 0,
			addEditConfig: false,
			removeConfig: false,
			saved: false,
			removed: false
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
			});

			this.$http.get('/api/getAvailableTracks')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.tracks = resp.data;
			});

			this.$http.get('/api/getAvailableCars')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.cars = resp.data;
			});
		},
		_reset: function(){
			this._id = 0;
			this.err = 0;
			this.addEditConfig = false;
			this.removeConfig = false;
			this.saved = false;
			this.removed = false;
		},
		openAddEditConfig: function(id){
			this._reset();
			
			if(id){
				this._id = id;

				this.$http.get('/api/getConfiguration', {params: {id: id}})
				.then(function(resp){
					if(resp.data.code){
						console.log(resp.data.code+': '+resp.data.msg);
						return;
					}

					// TODO
					this.addEditConfig = true;
				});
			}
			else{
				this.addEditConfig = true;
			}
		},
		openRemoveConfig: function(id){
			this._reset();

			if(!id){
				return;
			}

			this._id = id;
			this.removeConfig = true;
		},
		copyConfig: function(id){

		},
		performAddEditConfig: function(){
			var weather = [];
			var cars = [];

			var data = {
				id: this._id,
				name: this.name,
				pwd: this.pwd,
				admin_pwd: this.admin_pwd,
				pickup_mode: this.pickup_mode,
				race_overtime: this.race_overtime,
				max_slots: this.max_slots,
				description: this.description,
				udp: this.udp,
				tcp: this.tcp,
				http: this.http,
				packets_hz: this.packets_hz,
				loop_mode: this.loop_mode,
				show_in_lobby: this.show_in_lobby,
				abs: this.abs,
				tc: this.tc,
				stability_aid: this.stability_aid,
				auto_clutch: this.auto_clutch,
				tyre_blankets: this.tyre_blankets,
				force_virtual_mirror: this.force_virtual_mirror,
				fuel_rate: this.fuel_rate,
				damage_rate: this.damage_rate,
				tires_wear_rate: this.tires_wear_rate,
				allowed_tires_out: this.allowed_tires_out,
				max_ballast: this.max_ballast,
				dynamic_track: this.dynamic_track,
				condition: this.condition,
				start_value: this.start_value,
				randomness: this.randomness,
				transferred_grip: this.transferred_grip,
				laps_to_improve_grip: this.laps_to_improve_grip,
				kick_vote_quorum: this.kick_vote_quorum,
				session_vote_quorum: this.session_vote_quorum,
				vote_duration: this.vote_duration,
				blacklist: this.blacklist,
				booking: this.booking,
				booking_time: this.booking_time,
				practice: this.practice,
				practice_time: this.practice_time,
				can_join_practice: this.can_join_practice,
				qualify: this.qualify,
				qualify_time: this.qualify_time,
				can_join_qualify: this.can_join_qualify,
				race: this.race,
				race_time: this.race_time,
				race_wait_time: this.race_wait_time,
				join_type: this.join_type,
				time: this.time,
				weather: weather,
				track: this.track,
				cars: cars
			};

			console.log(data);

			/*this.$http.post('/api/addEditConfiguration', data)
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.saved = true;
			});*/
		},
		performRemoveConfig: function(){
			this.$http.post('/api/removeConfiguration', {id: this._id})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this._reset();
				this._load();
				this.removed = true;
			});
		}
	}
});
