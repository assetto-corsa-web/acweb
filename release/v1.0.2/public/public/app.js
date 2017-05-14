var VERSION = '1.0.0';
var _MS_PER_DAY = 1000 * 60 * 60 * 24;

// Returns date in format dd.mm.yyyy hh:mi:ss.
Date.prototype.formatDE = function(withTime){
	var day = _addLeadingZero(this.getDate());
	var month = _addLeadingZero(this.getMonth()+1);
	var year = this.getFullYear();
	var hours = _addLeadingZero(this.getHours());
	var minutes = _addLeadingZero(this.getMinutes());
	var seconds = _addLeadingZero(this.getSeconds());

	return [day, month, year].join('.')+' '+[hours, minutes, seconds].join(':');
};

// Returns date in format dd.mm.yyyy.
Date.prototype.formatDDMMYYYY = function(){
	var day = _addLeadingZero(this.getDate());
	var month = _addLeadingZero(this.getMonth()+1);
	var year = this.getFullYear();

	return [day, month, year].join('.');
};

// Returns date in format yyyy-mm-dd.
Date.prototype.formatYYYYMMDD = function(){
	var day = _addLeadingZero(this.getDate());
	var month = _addLeadingZero(this.getMonth()+1);
	var year = this.getFullYear();

	return [year, month, day].join('-');
};

// Returns date in format hh:mi.
Date.prototype.formatHHMI = function(){
	var hours = _addLeadingZero(this.getHours());
	var minutes = _addLeadingZero(this.getMinutes());

	return [hours, minutes].join(':');
};

// Parses time only from given ISO string.
Date.parseTime = function(time){
	var date = new Date(0);
	var hours = parseInt(time.substring(11, 13));
	var minutes = parseInt(time.substring(14, 16));

	date.setHours(hours);
	date.setMinutes(minutes);

	return date;
};

function _addLeadingZero(time){
	if(time < 10){
		return '0'+time
	}

	return time
}
// Extracts hours and minutes in 24h format from string in format hh:mi
// and returns them as object. Throws an error if time is invalid.
function parseTime(time){
	var split = time.split(':');
	var hours = parseInt(split[0]);
	var minutes = parseInt(split[1]);

	if(hours < 0 || hours > 23 || minutes < 0 || minutes > 59){
		throw 'Hours and minutes must be within range 00-23 and 00-59';
	}

	return {hours: hours, minutes: minutes};
}

// Returns hours, minutes for given date in format hh:mi.
function getTimeHHMI(date){
	var minutes = date/1000/60;
	var hours = Math.floor(minutes/60);
	minutes -= hours*60;
	hours = _addLeadingZero(hours);
	minutes = _addLeadingZero(minutes);

	return [hours, minutes].join(':');
}
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
Vue.component('msg', {
	props: ['type', 'msg'],
	template: '<div v-bind:class="type">\
	{{msg}}\
</div>'
});
Vue.component('hmenu', {
	template: '<div class="menu">\
	<div class="wrapper">\
		<div class="logo"></div>\
		<router-link to="/instance">Instances</router-link>\
		<router-link to="/configuration">Configurations</router-link>\
		<router-link to="/settings">Settings</router-link>\
		<router-link to="/user">User</router-link>\
		<router-link to="/about">About</router-link>\
\
		<a href="#" class="right" v-on:click="performLogout()"><i class="fa fa-sign-out" aria-hidden="true" title="Logout"></i></a>\
	</div>\
</div>',
	data: function(){
		return {
			version: VERSION
		}
	},
	methods: {
		performLogout: function(){
			SessionService.logout(this.$router);
		}
	}
});
Vue.component('About', {
	template: '<div>\
	<hmenu></hmenu>\
\
	<div class="main">\
		<h1>About</h1>\
		\
		<div class="box">\
			<div class="wrapper">\
				Assetto Corsa server web interface<br />\
				Version {{version}}<br />\
				&copy; 2017 Marvin Blum<br />\
				<br />\
				<a href="https://github.com/DeKugelschieber/acweb" target="_blank">View on GitHub</a>\
			</div>\
		</div>\
	</div>\
</div>',
	data: function(){
		return {
			version: VERSION
		}
	}
});
Vue.component('Configuration', {
	template: '<div>\
	<hmenu></hmenu>\
\
	<div class="main">\
		<h1>Configurations</h1>\
		\
		<div class="box" v-if="addEditConfig">\
			<div class="wrapper">\
				<h2>Create/Edit Configuration</h2>\
\
				<msg :type="\'error\'" :msg="\'The name must be set.\'" v-if="err == 1"></msg>\
				<msg :type="\'error\'" :msg="\'At least one weather configuration must be added.\'" v-if="err == 2"></msg>\
				<msg :type="\'error\'" :msg="\'At least one car must be added.\'" v-if="err == 3"></msg>\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<form v-on:submit.prevent="performAddEditConfig()">\
					<table>\
						<tr>\
							<td colspan="2"><h3>Basic Settings</h3></td>\
						</tr>\
						<tr>\
							<td class="w20">Name:</td>\
							<td><input type="text" name="name" class="full-width" v-model="name" /></td>\
						</tr>\
						<tr>\
							<td>Password:</td>\
							<td><input type="text" name="pwd" class="full-width" v-model="pwd" /></td>\
						</tr>\
						<tr>\
							<td>Admin password:</td>\
							<td><input type="text" name="admin_pwd" class="full-width" v-model="admin_pwd" /></td>\
						</tr>\
						<tr>\
							<td>Pickup mode:</td>\
							<td><input type="checkbox" name="pickup_mode" v-model="pickup_mode" /></td>\
						</tr>\
						<tr>\
							<td>Lock entry list:</td>\
							<td><input type="checkbox" name="lock_entry_list" v-model="lock_entry_list" /></td>\
						</tr>\
						<tr>\
							<td>Race overtime:</td>\
							<td><input type="number" name="race_overtime" v-model="race_overtime" /></td>\
						</tr>\
						<tr>\
							<td>Max. slots:</td>\
							<td><input type="number" name="max_slots" v-model="max_slots" /></td>\
						</tr>\
						<tr>\
							<td>Result screen time:</td>\
							<td><input type="number" name="result_screen_time" v-model="result_screen_time" /></td>\
						</tr>\
						<tr>\
							<td>Welcome message:</td>\
							<td><input type="text" name="welcome" class="full-width" v-model="welcome" /></td>\
						</tr>\
						<tr>\
							<td>Description:</td>\
							<td><textarea name="description" v-model="description"></textarea></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Advanced Settings</h3></td>\
						</tr>\
						<tr>\
							<td>UDP port:</td>\
							<td><input type="number" name="udp" v-model="udp" /></td>\
						</tr>\
						<tr>\
							<td>TCP port:</td>\
							<td><input type="number" name="tcp" v-model="tcp" /></td>\
						</tr>\
						<tr>\
							<td>HTTP port:</td>\
							<td><input type="number" name="http" v-model="http" /></td>\
						</tr>\
						<tr>\
							<td>Packets Hz:</td>\
							<td><input type="number" name="packets_hz" v-model="packets_hz" /></td>\
						</tr>\
						<tr>\
							<td>Loop mode:</td>\
							<td><input type="checkbox" name="loop_mode" v-model="loop_mode" /></td>\
						</tr>\
						<tr>\
							<td>Show on lobby:</td>\
							<td><input type="checkbox" name="show_in_lobby" v-model="show_in_lobby" /></td>\
						</tr>\
						<tr>\
							<td>Threads:</td>\
							<td>\
								<select name="threads" v-model="threads">\
									<option>2</option>\
									<option>3</option>\
									<option>4</option>\
									<option>5</option>\
									<option>6</option>\
									<option>7</option>\
									<option>8</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Assists</h3></td>\
						</tr>\
						<tr>\
							<td>ABS:</td>\
							<td>\
								<select name="abs" v-model="abs">\
									<option value="0">Denied</option>\
									<option value="1">Factory</option>\
									<option value="2">Forced</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td>TC:</td>\
							<td>\
								<select name="tc" v-model="tc">\
									<option value="0">Denied</option>\
									<option value="1">Factory</option>\
									<option value="2">Forced</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td>Stability aid:</td>\
							<td><input type="checkbox" name="stability_aid" v-model="stability_aid" /></td>\
						</tr>\
						<tr>\
							<td>Auto clutch:</td>\
							<td><input type="checkbox" name="auto_clutch" v-model="auto_clutch" /></td>\
						</tr>\
						<tr>\
							<td>Tyre blankets:</td>\
							<td><input type="checkbox" name="tyre_blankets" v-model="tyre_blankets" /></td>\
						</tr>\
						<tr>\
							<td>Force virtual mirror:</td>\
							<td><input type="checkbox" name="force_virtual_mirror" v-model="force_virtual_mirror" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Realism</h3></td>\
						</tr>\
						<tr>\
							<td>Fuel rate:</td>\
							<td><input type="number" name="fuel_rate" v-model="fuel_rate" /></td>\
						</tr>\
						<tr>\
							<td>Damage rate:</td>\
							<td><input type="number" name="damage_rate" v-model="damage_rate" /></td>\
						</tr>\
						<tr>\
							<td>Tires wear rate:</td>\
							<td><input type="number" name="tires_wear_rate" v-model="tires_wear_rate" /></td>\
						</tr>\
						<tr>\
							<td>Allowed tires out:</td>\
							<td>\
								<select name="allowed_tires_out" v-model="allowed_tires_out">\
									<option>0</option>\
									<option>1</option>\
									<option>2</option>\
									<option>3</option>\
									<option>4</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td>Max ballast:</td>\
							<td><input type="number" name="max_ballast" v-model="max_ballast" /></td>\
						</tr>\
						<tr>\
							<td>Disable gas cut penality:</td>\
							<td><input type="checkbox" name="disable_gas_cut_penality" v-model="disable_gas_cut_penality" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Dynamic Track</h3></td>\
						</tr>\
						<tr>\
							<td>Dynamic track:</td>\
							<td><input type="checkbox" name="dynamic_track" v-model="dynamic_track" /></td>\
						</tr>\
						<tr>\
							<td>Condition:</td>\
							<td>\
								<select name="condition" v-model="condition">\
									<option>CUSTOM</option>\
									<option>DUSTY</option>\
									<option>OLD</option>\
									<option>SLOW</option>\
									<option>GREEN</option>\
									<option>FAST</option>\
									<option>OPTIMUM</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td>Start value:</td>\
							<td><input type="number" name="start_value" v-model="start_value" /></td>\
						</tr>\
						<tr>\
							<td>Randomness:</td>\
							<td><input type="number" name="randomness" v-model="randomness" /></td>\
						</tr>\
						<tr>\
							<td>Transferred grip:</td>\
							<td><input type="number" name="transferred_grip" v-model="transferred_grip" /></td>\
						</tr>\
						<tr>\
							<td>Laps to improve grip:</td>\
							<td><input type="number" name="laps_to_improve_grip" v-model="laps_to_improve_grip" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Voting</h3></td>\
						</tr>\
						<tr>\
							<td>Kick vote quorum:</td>\
							<td><input type="number" name="kick_vote_quorum" v-model="kick_vote_quorum" /></td>\
						</tr>\
						<tr>\
							<td>Session vote quorum:</td>\
							<td><input type="number" name="session_vote_quorum" v-model="session_vote_quorum" /></td>\
						</tr>\
						<tr>\
							<td>Vote duration:</td>\
							<td><input type="number" name="vote_duration" v-model="vote_duration" /></td>\
						</tr>\
						<tr>\
							<td>Blacklist:</td>\
							<td>\
								<select name="blacklist" v-model="blacklist">\
									<option value="0">Kick Player</option>\
									<option value="1">Kick Until Restart</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td>Max. collisions per KM:</td>\
							<td><input type="number" name="max_collisions_km" v-model="max_collisions_km" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Booking</h3></td>\
						</tr>\
						<tr>\
							<td>Booking:</td>\
							<td><input type="checkbox" name="booking" v-model="booking" /></td>\
						</tr>\
						<tr>\
							<td>Booking time:</td>\
							<td><input type="number" name="booking_time" v-model="booking_time" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Practice</h3></td>\
						</tr>\
						<tr>\
							<td>Practice:</td>\
							<td><input type="checkbox" name="practice" v-model="practice" /></td>\
						</tr>\
						<tr>\
							<td>Practice time:</td>\
							<td><input type="number" name="practice_time" v-model="practice_time" /></td>\
						</tr>\
						<tr>\
							<td>Can join:</td>\
							<td><input type="checkbox" name="can_join_practice" v-model="can_join_practice" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Qualify</h3></td>\
						</tr>\
						<tr>\
							<td>Qualify:</td>\
							<td><input type="checkbox" name="qualify" v-model="qualify" /></td>\
						</tr>\
						<tr>\
							<td>Qualify time:</td>\
							<td><input type="number" name="qualify_time" v-model="qualify_time" /></td>\
						</tr>\
						<tr>\
							<td>Can join:</td>\
							<td><input type="checkbox" name="can_join_qualify" v-model="can_join_qualify" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Race</h3></td>\
						</tr>\
						<tr>\
							<td>Race:</td>\
							<td><input type="checkbox" name="race" v-model="race" /></td>\
						</tr>\
						<tr>\
							<td>Race laps:</td>\
							<td><input type="number" name="race_laps" v-model="race_laps" /></td>\
						</tr>\
						<tr>\
							<td>Race time:</td>\
							<td><input type="number" name="race_time" v-model="race_time" /></td>\
						</tr>\
						<tr>\
							<td>Race wait time:</td>\
							<td><input type="number" name="race_wait_time" v-model="race_wait_time" /></td>\
						</tr>\
						<tr>\
							<td>Race extra lap:</td>\
							<td><input type="checkbox" name="race_extra_lap" v-model="race_extra_lap" /></td>\
						</tr>\
						<tr>\
							<td>Join type:</td>\
							<td>\
								<select name="join_type" v-model="join_type">\
									<option value="0">Close</option>\
									<option value="1">Open</option>\
									<option value="2">Close at Start</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Time</h3></td>\
						</tr>\
						<tr>\
							<td>Time:</td>\
							<td>\
								<input type="time" name="time" v-model="time" />\
								(08:00 - 18:00)\
							</td>\
						</tr>\
						<tr>\
							<td>Sun angle:</td>\
							<td><input type="number" name="sun_angle" v-model="sun_angle" /></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Weather</h3></td>\
						</tr>\
						<tr v-for="(w, i) in weather">\
							<td colspan="2">\
								<table>\
									<tr>\
										<td>Weather:</td>\
										<td>\
											<select name="weather" v-model="w.weather">\
												<option value="3_clear">Clear</option>\
												<option value="7_heavy_clouds">Heavy Clouds</option>\
												<option value="1_heavy_fog">Heavy Fog</option>\
												<option value="5_light_clouds">Light Clouds</option>\
												<option value="2_light_fog">Light Fog</option>\
												<option value="4_mid_clear">Mid Clear</option>\
												<option value="6_mid_clouds">Mid Clouds</option>\
											</select>\
										</td>\
										<td>\
											<button v-on:click.prevent="removeWeather(i)">Remove Weather Panel</button>\
										</td>\
									</tr>\
									<tr>\
										<td>Base ambient temp:</td>\
										<td><input type="number" name="base_ambient_temp" v-model="w.base_ambient_temp" /></td>\
										<td></td>\
									</tr>\
									<tr>\
										<td>Realistic road temp:</td>\
										<td><input type="number" name="realistic_road_temp" v-model="w.realistic_road_temp" /></td>\
										<td></td>\
									</tr>\
									<tr>\
										<td>Base road temp:</td>\
										<td><input type="number" name="base_road_temp" v-model="w.base_road_temp" /></td>\
										<td></td>\
									</tr>\
									<tr>\
										<td>Ambient variation:</td>\
										<td><input type="number" name="ambient_variation" v-model="w.ambient_variation" /></td>\
										<td></td>\
									</tr>\
									<tr>\
										<td>Road variation:</td>\
										<td><input type="number" name="road_variation" v-model="w.road_variation" /></td>\
										<td></td>\
									</tr>\
								</table>\
							</td>\
						</tr>\
						<tr>\
							<td colspan="2"><button v-on:click.prevent="addWeather()">Add Weather Panel</button></td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Track</h3></td>\
						</tr>\
						<tr>\
							<td colspan="2">\
								<div class="select">\
									<div class="option" v-for="(track, i) in tracks" v-on:click="selectTrack(i)" v-bind:class="{selected: selectedTrack == i}">\
										<strong>{{track.name}}</strong>\
										({{track.max_slots}})\
										{{track.description}}\
									</div>\
								</div>\
							</td>\
						</tr>\
						<tr>\
							<td colspan="2"><h3>Cars/Entry List</h3></td>\
						</tr>\
						<tr>\
							<td colspan="2">\
								<table>\
									<tr>\
										<td class="w33">\
											<strong>Car</strong>\
\
											<div class="select">\
												<div class="option" v-for="(car, i) in cars" v-on:click="selectCar(i)" v-bind:class="{selected: selectedCar == i}">\
													<strong>{{car.name}}</strong>\
													{{car.description}}\
												</div>\
											</div>\
										</td>\
										<td class="w33">\
											<strong>Painting</strong>\
\
											<div class="select">\
												<div class="option" v-for="(painting, i) in activePaintings" v-on:click="selectPainting(i)" v-bind:class="{selected: selectedPainting == i}">\
													<strong>{{painting}}</strong>\
												</div>\
											</div>\
										</td>\
										<td class="w33 top">\
											<strong>Slot Settings</strong>\
\
											<table>\
												<tr>\
													<td class="w20">Spectator:</td>\
													<td><input type="checkbox" name="spectator" v-model="spectator" /></td>\
												</tr>\
												<tr>\
													<td>Driver:</td>\
													<td><input type="text" name="driver" v-model="driver" /></td>\
												</tr>\
												<tr>\
													<td>Team:</td>\
													<td><input type="text" name="team" v-model="team" /></td>\
												</tr>\
												<tr>\
													<td>GUID:</td>\
													<td><input type="text" name="guid" v-model="guid" /></td>\
												</tr>\
												<tr>\
													<td></td>\
													<td><button v-on:click.prevent="addCar()">Add Car</button></td>\
												</tr>\
											</table>\
										</td>\
									</tr>\
								</table>\
							</td>\
						</tr>\
						<tr>\
							<td colspan="2">\
								<table>\
									<thead>\
										<tr>\
											<td class="w20">Car</td>\
											<td class="w20">Painting</td>\
											<td class="w5">Spectator</td>\
											<td class="w10">Driver</td>\
											<td class="w10">Team</td>\
											<td class="w20">GUID</td>\
											<td></td>\
										</tr>\
									</thead>\
									<tbody>\
										<tr v-for="(car, i) in selectedCars">\
											<td>{{car.car}}</td>\
											<td>{{car.painting}}</td>\
											<td>{{car.spectator}}</td>\
											<td>{{car.driver}}</td>\
											<td>{{car.team}}</td>\
											<td>{{car.guid}}</td>\
											<td>\
												<i class="fa fa-angle-up" aria-hidden="true" title="Move up" v-on:click.prevent="carUp(i)"></i>\
												<i class="fa fa-angle-down" aria-hidden="true" title="Move down" v-on:click.prevent="carDown(i)"></i>\
												<i class="fa fa-trash" aria-hidden="true" title="Remove car" v-on:click.prevent="removeCar(i)"></i>\
											</td>\
										</tr>\
									</tbody>\
								</table>\
							</td>\
						</tr>\
					</table>\
\
					<msg :type="\'error\'" :msg="\'The name must be set.\'" v-if="err == 1"></msg>\
					<msg :type="\'error\'" :msg="\'At least one weather configuration must be added.\'" v-if="err == 2"></msg>\
					<msg :type="\'error\'" :msg="\'At least one car must be added.\'" v-if="err == 3"></msg>\
					<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
					<input type="submit" value="Save" />\
					<button v-on:click.prevent="addEditConfig = false">Cancel</button>\
				</form>\
			</div>\
		</div>\
\
		<div class="box" v-if="removeConfig">\
			<div class="wrapper">\
				<h2>Remove Configuration</h2>\
\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<p>Do you really want to remove this configuration? This won\'t stop any instances currently running.</p>\
\
				<button v-on:click="performRemoveConfig()">Yes, remove configuration</button>\
				<button v-on:click="removeConfig = false">Cancel</button>\
			</div>\
		</div>\
\
		<div class="box">\
			<div class="wrapper">\
				<msg :type="\'success\'" :msg="\'The configuration has been saved.\'" v-if="saved"></msg>\
				<msg :type="\'success\'" :msg="\'The configuration has been removed.\'" v-if="removed"></msg>\
\
				<button v-on:click="openAddEditConfig(0)">Add Configuration</button>\
\
				<table>\
					<thead>\
						<tr>\
							<td class="w5">ID</td>\
							<td class="w40">Name</td>\
							<td class="w40">Info</td>\
							<td></td>\
						</tr>\
					</thead>\
					<tbody>\
						<tr v-for="config in configs">\
							<td>{{config.id}}</td>\
							<td>{{config.name}}</td>\
							<td>{{config.track}}</td>\
							<td>\
								<i class="fa fa-pencil" aria-hidden="true" title="Edit configuration" v-on:click="openAddEditConfig(config.id)"></i>\
								<i class="fa fa-files-o" aria-hidden="true" title="Copy configuration" v-on:click="copyConfig(config.id)"></i>\
								<i class="fa fa-trash" aria-hidden="true" title="Remove configuration" v-on:click="openRemoveConfig(config.id)"></i>\
							</td>\
						</tr>\
					</tbody>\
				</table>\
			</div>\
		</div>\
	</div>\
</div>',
	data: function(){
		return {
			configs: [],
			tracks: [],
			cars: [],
			activePaintings: [],
			selectedTrack: 0,
			selectedCar: 0,
			selectedPainting: 0,
			spectator: false,
			driver: '',
			team: '',
			guid: '',
			// ---
			selectedCars: [],
			weather: [],
			// ---
			_id: 0,
			name: 'Servername',
			pwd: '',
			admin_pwd: '',
			pickup_mode: true,
			lock_entry_list: false,
			race_overtime: 60,
			max_slots: 0,
			result_screen_time: 60,
			welcome: '',
			description: '',
			udp: 9600,
			tcp: 9600,
			http: 8081,
			packets_hz: 18,
			loop_mode: true,
			show_in_lobby: true,
			threads: 2,
			abs: '',
			tc: '',
			stability_aid: true,
			auto_clutch: true,
			tyre_blankets: false,
			force_virtual_mirror: false,
			fuel_rate: 100,
			damage_rate: 50,
			tires_wear_rate: 100,
			allowed_tires_out: 2,
			max_ballast: 150,
			disable_gas_cut_penality: false,
			dynamic_track: true,
			condition: '',
			start_value: 100,
			randomness: 0,
			transferred_grip: 100,
			laps_to_improve_grip: 1,
			kick_vote_quorum: 70,
			session_vote_quorum: 70,
			vote_duration: 15,
			blacklist: '',
			max_collisions_km: 5,
			booking: false,
			booking_time: 0,
			practice: true,
			practice_time: 15,
			can_join_practice: true,
			qualify: true,
			qualify_time: 15,
			can_join_qualify: true,
			race: true,
			race_laps: 10,
			race_time: 0,
			race_wait_time: 60,
			race_extra_lap: false,
			join_type: '',
			time: '08:00',
			sun_angle: 32,
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
			this.$http.get('/api/configuration')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.configs = resp.data;
			});

			this.$http.get('/api/tracks')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.tracks = resp.data;
			});

			this.$http.get('/api/cars')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.cars = resp.data;

				if(this.cars.length){
					this.activePaintings = this.cars[0].paintings;
				}
			});
		},
		_reset: function(){
			this.spectator = false;
			this.driver = '';
			this.team = '';
			this.guid = '';

			this.selectedCars = [];
			this.weather = [];

			this._id = 0;
			this.name = 'Servername';
			this.pwd = '';
			this.admin_pwd = '';
			this.pickup_mode = true;
			this.lock_entry_list = false;
			this.race_overtime = 60;
			this.max_slots = 0;
			this.result_screen_time = 60;
			this.welcome = '';
			this.description = '';
			this.udp = 9600;
			this.tcp = 9600;
			this.http = 8081;
			this.packets_hz = 18;
			this.loop_mode = true;
			this.show_in_lobby = true;
			this.threads = 2;
			this.abs = '';
			this.tc = '';
			this.stability_aid = true;
			this.auto_clutch = true;
			this.tyre_blankets = false;
			this.force_virtual_mirror = false;
			this.fuel_rate = 100;
			this.damage_rate = 50;
			this.tires_wear_rate = 100;
			this.allowed_tires_out = 2;
			this.max_ballast = 150;
			this.disable_gas_cut_penality = false;
			this.dynamic_track = true;
			this.condition = '';
			this.start_value = 100;
			this.randomness = 0;
			this.transferred_grip = 100;
			this.laps_to_improve_grip = 1;
			this.kick_vote_quorum = 70;
			this.session_vote_quorum = 70;
			this.vote_duration = 15;
			this.blacklist = '';
			this.max_collisions_km = 5;
			this.booking = false;
			this.booking_time = 0;
			this.practice = true;
			this.practice_time = 15;
			this.can_join_practice = true;
			this.qualify = true;
			this.qualify_time = 15;
			this.can_join_qualify = true;
			this.race = true;
			this.race_laps = 10;
			this.race_time = 0;
			this.race_wait_time = 60;
			this.race_extra_lap = false;
			this.join_type = '';
			this.time = '08 =00';
			this.sun_angle = 32;
			this.track = '';

			this.err = 0;
			this.addEditConfig = false;
			this.removeConfig = false;
			this.saved = false;
			this.removed = false;
		},
		_openConfig: function(id, copy){
			this.$http.get('/api/configuration', {params: {id: id}})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				// config
				this.name = resp.data.name;
				this.pwd = resp.data.pwd;
				this.admin_pwd = resp.data.admin_pwd;
				this.pickup_mode = resp.data.pickup_mode;
				this.lock_entry_list = resp.data.lock_entry_list;
				this.race_overtime = resp.data.race_overtime;
				this.max_slots = resp.data.max_slots;
				this.welcome = resp.data.welcome;
				this.description = resp.data.description;
				this.udp = resp.data.udp;
				this.tcp = resp.data.tcp;
				this.http = resp.data.http;
				this.packets_hz = resp.data.packets_hz;
				this.loop_mode = resp.data.loop_mode;
				this.show_in_lobby = resp.data.show_in_lobby;
				this.threads = resp.data.threads;
				this.abs = resp.data.abs;
				this.tc = resp.data.tc;
				this.stability_aid = resp.data.stability_aid;
				this.auto_clutch = resp.data.auto_clutch;
				this.tyre_blankets = resp.data.tyre_blankets;
				this.force_virtual_mirror = resp.data.force_virtual_mirror;
				this.fuel_rate = resp.data.fuel_rate;
				this.damage_rate = resp.data.damage_rate;
				this.tires_wear_rate = resp.data.tires_wear_rate;
				this.allowed_tires_out = resp.data.allowed_tires_out;
				this.max_ballast = resp.data.max_ballast;
				this.disable_gas_cut_penality = resp.data.disable_gas_cut_penality;
				this.result_screen_time = resp.data.result_screen_time;
				this.dynamic_track = resp.data.dynamic_track;
				this.condition = resp.data.condition;
				this.start_value = resp.data.start_value;
				this.randomness = resp.data.randomness;
				this.transferred_grip = resp.data.transferred_grip;
				this.laps_to_improve_grip = resp.data.laps_to_improve_grip;
				this.kick_vote_quorum = resp.data.kick_vote_quorum;
				this.session_vote_quorum = resp.data.session_vote_quorum;
				this.vote_duration = resp.data.vote_duration;
				this.blacklist = resp.data.blacklist;
				this.max_collisions_km = resp.data.max_collisions_km;
				this.booking = resp.data.booking;
				this.booking_time = resp.data.booking_time;
				this.practice = resp.data.practice;
				this.practice_time = resp.data.practice_time;
				this.can_join_practice = resp.data.can_join_practice;
				this.qualify = resp.data.qualify;
				this.qualify_time = resp.data.qualify_time;
				this.can_join_qualify = resp.data.can_join_qualify;
				this.race = resp.data.race;
				this.race_laps = resp.data.race_laps;
				this.race_time = resp.data.race_time;
				this.race_wait_time = resp.data.race_wait_time;
				this.race_extra_lap = resp.data.race_extra_lap;
				this.join_type = resp.data.join_type;
				this.time = resp.data.time;
				this.sun_angle = resp.data.sun_angle;

				if(copy){
					this.name += ' (copy)';
				}

				// track
				for(var i = 0; i < this.tracks.length; i++){
					if(this.tracks[i].name == resp.data.track && this.tracks[i].config == resp.data.track_config){
						this.selectTrack(i);
						break;
					}
				}
				
				// weather
				this.weather = resp.data.weather;

				if(copy){
					for(var i = 0; i < this.weather.length; i++){
						this.weather[i].id = 0;
					}
				}

				// cars
				this.selectedCars = resp.data.cars;

				if(copy){
					for(var i = 0; i < this.selectedCars.length; i++){
						this.selectedCars[i].id = 0;
					}
				}

				this.addEditConfig = true;
			});
		},
		openAddEditConfig: function(id){
			this._reset();
			
			if(id){
				this._id = id;
				this._openConfig(id, false);
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
			this._openConfig(id, true);
		},
		performAddEditConfig: function(){
			for(var i = 0; i < this.weather.length; i++){
				this.weather[i].base_ambient_temp = parseInt(this.weather[i].base_ambient_temp);
				this.weather[i].realistic_road_temp = parseInt(this.weather[i].realistic_road_temp);
				this.weather[i].base_road_temp = parseInt(this.weather[i].base_road_temp);
				this.weather[i].ambient_variation = parseInt(this.weather[i].ambient_variation);
				this.weather[i].road_variation = parseInt(this.weather[i].road_variation);
			}

			for(var i = 0; i < this.selectedCars.length; i++){
				this.selectedCars[i].position = i;
			}

			var data = {
				id: this._id,
				name: this.name,
				pwd: this.pwd,
				admin_pwd: this.admin_pwd,
				pickup_mode: this.pickup_mode,
				lock_entry_list: this.lock_entry_list,
				race_overtime: parseInt(this.race_overtime),
				max_slots: parseInt(this.max_slots),
				welcome: this.welcome,
				description: this.description,
				udp: parseInt(this.udp),
				tcp: parseInt(this.tcp),
				http: parseInt(this.http),
				packets_hz: parseInt(this.packets_hz),
				loop_mode: this.loop_mode,
				show_in_lobby: this.show_in_lobby,
				threads: parseInt(this.threads),
				abs: this.abs,
				tc: this.tc,
				stability_aid: this.stability_aid,
				auto_clutch: this.auto_clutch,
				tyre_blankets: this.tyre_blankets,
				force_virtual_mirror: this.force_virtual_mirror,
				fuel_rate: parseInt(this.fuel_rate),
				damage_rate: parseInt(this.damage_rate),
				tires_wear_rate: parseInt(this.tires_wear_rate),
				allowed_tires_out: parseInt(this.allowed_tires_out),
				max_ballast: parseInt(this.max_ballast),
				disable_gas_cut_penality: this.disable_gas_cut_penality,
				result_screen_time: parseInt(this.result_screen_time),
				dynamic_track: this.dynamic_track,
				condition: this.condition,
				start_value: parseInt(this.start_value),
				randomness: parseInt(this.randomness),
				transferred_grip: parseInt(this.transferred_grip),
				laps_to_improve_grip: parseInt(this.laps_to_improve_grip),
				kick_vote_quorum: parseInt(this.kick_vote_quorum),
				session_vote_quorum: parseInt(this.session_vote_quorum),
				vote_duration: parseInt(this.vote_duration),
				blacklist: this.blacklist,
				max_collisions_km: parseInt(this.max_collisions_km),
				booking: this.booking,
				booking_time: parseInt(this.booking_time),
				practice: this.practice,
				practice_time: parseInt(this.practice_time),
				can_join_practice: this.can_join_practice,
				qualify: this.qualify,
				qualify_time: parseInt(this.qualify_time),
				can_join_qualify: this.can_join_qualify,
				race: this.race,
				race_laps: parseInt(this.race_laps),
				race_time: parseInt(this.race_time),
				race_wait_time: parseInt(this.race_wait_time),
				race_extra_lap: this.race_extra_lap,
				join_type: this.join_type,
				time: this.time,
				sun_angle: parseInt(this.sun_angle),
				weather: this.weather,
				track: this.track.name,
				track_config: this.track.config,
				cars: this.selectedCars
			};

			this.$http.post('/api/configuration', data)
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.saved = true;
			});
		},
		performRemoveConfig: function(){
			this.$http.delete('/api/configuration', {params: {id: this._id}})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.removed = true;
			});
		},
		addWeather: function(){
			this.weather.push({
				weather: 'Clear',
				base_ambient_temp: 20,
				realistic_road_temp: 1,
				base_road_temp: 18,
				ambient_variation: 1,
				road_variation: 1
			});
		},
		removeWeather: function(i){
			this.weather.splice(i, 1);
		},
		selectTrack: function(i){
			this.selectedTrack = i;
			this.track = this.tracks[i];
		},
		selectCar: function(i){
			this.selectedCar = i;
			this.selectedPainting = 0;
			this.activePaintings = this.cars[i].paintings;
		},
		selectPainting: function(i){
			this.selectedPainting = i;
		},
		addCar: function(){
			var car = this.cars[this.selectedCar];
			this.selectedCars.push({
				car: car.name,
				painting: car.paintings[this.selectedPainting],
				spectator: this.spectator,
				driver: this.driver,
				team: this.team,
				guid: this.guid,
				position: this.selectedCars.length
			});

			// only reset driver and GUID in case user wants to add multiple similar slots
			this.driver = '';
			this.guid = '';
		},
		carUp: function(i){
			if(i == 0){
				return;
			}

			var car = this.selectedCars[i-1];
			Vue.set(this.selectedCars, i-1, this.selectedCars[i]);
			Vue.set(this.selectedCars, i, car);
		},
		carDown: function(i){
			if(i == this.selectedCars.length-1){
				return;
			}

			var car = this.selectedCars[i+1];
			Vue.set(this.selectedCars, i+1, this.selectedCars[i]);
			Vue.set(this.selectedCars, i, car);
		},
		removeCar: function(i){
			this.selectedCars.splice(i, 1);
		}
	}
});
Vue.component('Instance', {
	template: '<div>\
	<hmenu></hmenu>\
\
	<div class="main">\
		<h1>Status</h1>\
\
		<div class="box" v-if="startInstance">\
			<div class="wrapper">\
				<h2>Start New Instance</h2>\
\
				<msg :type="\'error\'" :msg="\'Error starting server instance.\'" v-if="err != 0 && err != 200"></msg>\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<form v-on:submit.prevent="performStart()">\
					<table>\
						<tr>\
							<td class="w20">Name:</td>\
							<td><input type="text" name="name" class="full-width" v-model="name" /></td>\
						</tr>\
						<tr>\
							<td>Configuration:</td>\
							<td>\
								<select name="configuration" v-model="config">\
									<option v-for="c in configs" v-bind:value="c.id">{{c.name}}</option>\
								</select>\
							</td>\
						</tr>\
						<tr>\
							<td></td>\
							<td>\
								<input type="submit" value="Start" />\
								<button v-on:click="startInstance = false">Cancel</button>\
							</td>\
						</tr>\
					</table>\
				</form>\
			</div>\
		</div>\
\
		<div class="box" v-if="stopInstance">\
			<div class="wrapper">\
				<h2>Stop Instance</h2>\
\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<p>Do you really want to stop this server instance?</p>\
\
				<button v-on:click="performStop()">Yes, stop instance</button>\
				<button v-on:click="stopInstance = false">Cancel</button>\
			</div>\
		</div>\
\
		<div class="box" v-if="showLog">\
			<div class="wrapper">\
				<h2>Log Output</h2>\
\
				<textarea v-model="log"></textarea>\
\
				<button v-on:click="showLog = false">Close</button>\
			</div>\
		</div>\
\
		<div class="box">\
			<div class="wrapper">\
				<h2>Active Server Instances</h2>\
\
				<msg :type="\'success\'" :msg="\'The server instance has been started.\'" v-if="started"></msg>\
				<msg :type="\'success\'" :msg="\'The server instance has been stopped.\'" v-if="stopped"></msg>\
\
				<button v-on:click="startInstance = true">Start New Instance</button>\
\
				<table>\
					<thead>\
						<tr>\
							<td class="w5">PID</td>\
							<td class="w40">Instance</td>\
							<td class="w40">Configuration</td>\
							<td class="w15"></td>\
						</tr>\
					</thead>\
					<tbody>\
						<tr v-for="instance in instances">\
							<td>{{instance.pid}}</td>\
							<td>{{instance.name}}</td>\
							<td>\
								{{instance.configuration.name}}<br />\
								{{instance.configuration.track}}<br />\
								TCP: {{instance.configuration.tcp}} UPD: {{instance.configuration.udp}}\
							</td>\
							<td>\
								<i class="fa fa-stop" aria-hidden="true" title="Stop instance" v-on:click="showStopInstance(instance.pid)"></i>\
							</td>\
						</tr>\
					</tbody>\
				</table>\
			</div>\
		</div>\
\
		<div class="box">\
			<div class="wrapper">\
				<h2>Log Files</h2>\
\
				<table>\
					<thead>\
						<tr>\
							<td class="w55">Filename</td>\
							<td class="w15">Date</td>\
							<td class="w15">Size</td>\
							<td class="w15"></td>\
						</tr>\
					</thead>\
					<tbody>\
						<tr v-for="log in logs">\
							<td>{{log.file}}</td>\
							<td>{{log.date}}</td>\
							<td>{{log.size}}</td>\
							<td>\
								<i class="fa fa-terminal" aria-hidden="true" title="Show console output" v-on:click="openLog(log.file)"></i>\
							</td>\
						</tr>\
					</tbody>\
				</table>\
			</div>\
		</div>\
	</div>\
</div>',
	data: function(){
		return {
			_pid: 0,
			instances: [],
			configs: [],
			logs: [],
			err: 0,
			name: '',
			config: 0,
			log: '',
			showLog: false,
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
			this.$http.get('/api/configuration')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.configs = resp.data;
				this._loadInstances();
			});

			this.$http.get('/api/instance/log')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.logs = resp.data;

				// reverse
				for(var i = this.logs.length-1; i >= 0; i--){
					this.logs[i].date = new Date(this.logs[i].date).formatDE();

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
		_loadInstances: function(){
			this.$http.get('/api/instance')
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
					return this.configs[i];
				}
			}

			return null;
		},
		_reset: function(){
			this.err = 0;
			this.name = '';
			this.config = 0;
			this.log = '';
			this.showLog = false;
			this.started = false;
			this.stopped = false;
			this.startInstance = false;
			this.stopInstance = false;
		},
		performStart: function(){
			this.$http.post('/api/instance', {name: this.name, config: this.config})
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
			this.$http.delete('/api/instance', {params: {pid: this._pid}})
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
		},
		openLog: function(file){
			this._reset();

			this.$http.get('/api/instance/log', {params: {file: file}})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				var log = resp.data.substr(1, resp.data.length-2).split('\\n');
				this.log = log.join('\n');
				this.showLog = true;
			});
		}
	}
});
Vue.component('Login', {
	template: '<div class="login">\
	<div class="center">\
		<img src="img/aclogo.png" alt="" />\
	</div>\
\
	<div class="box">\
		<div class="wrapper">\
			<h1>Login</h1>\
\
			<msg :type="\'error\'" :msg="\'The login and/or password was wrong.\'" v-if="err != 0"></msg>\
\
			<form v-on:submit.prevent="performLogin()">\
				<table>\
					<tr>\
						<td class="w30">Login/E-Mail:</td>\
						<td><input type="text" name="login" v-model="login" class="full-width" /></td>\
					</tr>\
					<tr>\
						<td>Password:</td>\
						<td><input type="password" name="pwd" v-model="pwd" class="full-width" /></td>\
					</tr>\
					<tr>\
						<td></td>\
						<td><input type="submit" value="Login" /></td>\
					</tr>\
				</table>\
			</form>\
		</div>\
	</div>\
\
	<div class="version">\
		<div class="wrapper">\
			Version {{version}} | &copy; 2017 Marvin Blum | <a href="https://github.com/DeKugelschieber/acweb" target="_blank">GitHub</a>\
		</div>\
	</div>\
</div>',
	data: function(){
		return {
			version: VERSION,
			err: 0,
			login: '',
			pwd: ''
		}
	},
	methods: {
		performLogin: function(){
			this.$http.post('/api/login', {login: this.login, pwd: this.pwd})
			.then(function(resp){
				if(resp.data.code){
					this.err = resp.data.code;
					return;
				}

				SessionService.login(this.$router, resp.data);
			});
		}
	}
});
Vue.component('Settings', {
	template: '<div>\
	<hmenu></hmenu>\
\
	<div class="main">\
		<h1>Settings</h1>\
		\
		<div class="box">\
			<div class="wrapper">\
				<h2>Server Settings</h2>\
\
				<msg :type="\'success\'" :msg="\'The settings have been saved.\'" v-if="saved"></msg>\
				<msg :type="\'error\'" :msg="\'The AC folder and the executable must be set.\'" v-if="err == 1"></msg>\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<form v-on:submit.prevent="performSave()">\
					<table>\
						<tr>\
							<td class="w10">AC server folder:</td>\
							<td><input type="text" name="path" class="full-width" v-model="folder" /></td>\
						</tr>\
						<tr>\
							<td>Executable:</td>\
							<td><input type="text" name="executable" class="full-width" v-model="executable" /></td>\
						</tr>\
						<tr>\
							<td>Arguments:</td>\
							<td><input type="text" name="args" class="full-width" v-model="args" /></td>\
						</tr>\
						<tr>\
							<td></td>\
							<td><input type="submit" value="Save" /></td>\
						</tr>\
					</table>\
				</form>\
			</div>\
		</div>\
\
		<div class="box">\
			<div class="wrapper">\
				The <strong>AC server folder</strong> must be the full path to your AC server installation folder, containing the acServer executable. Example: /home/acuser/steam/steamapps/common/Assetto Corsa Dedicated Server<br />\
				<strong>Executable</strong> is the executable file to start a server instance. Example: acServer<br />\
				<strong>Arguments</strong> are the arguments passed to the executable to start a server instance, separated by spaces. Example: linux<br /><br />\
\
				Make sure the web interface has the permissions to read, write and execute within the server folder. If you cannot start instances, please test your settings on your server.\
			</div>\
		</div>\
	</div>\
</div>',
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
Vue.component('User', {
	template: '<div>\
	<hmenu></hmenu>\
\
	<div class="main">\
		<h1>User</h1>\
\
		<div class="box" v-if="addEditUser">\
			<div class="wrapper">\
				<h2>Create/Edit User</h2>\
\
				<msg :type="\'error\'" :msg="\'Login, E-Mail and password must be set.\'" v-if="err == 1"></msg>\
				<msg :type="\'error\'" :msg="\'The passwords must be at least 8 characters long.\'" v-if="err == 2"></msg>\
				<msg :type="\'error\'" :msg="\'The passwords must be equal.\'" v-if="err == 3"></msg>\
				<msg :type="\'error\'" :msg="\'The Login and/or E-Mail is in use already.\'" v-if="err == 4"></msg>\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<form v-on:submit.prevent="performAddEditUser()">\
					<table>\
						<tr>\
							<td class="w20">Login:</td>\
							<td><input type="text" name="login" class="full-width" v-model="login" /></td>\
						</tr>\
						<tr>\
							<td>E-Mail:</td>\
							<td><input type="email" name="email" class="full-width" v-model="email" /></td>\
						</tr>\
						<tr>\
							<td>Password:</td>\
							<td><input type="password" name="pwd1" class="full-width" v-model="pwd1" /></td>\
						</tr>\
						<tr>\
							<td>Repeat password:</td>\
							<td><input type="password" name="pwd2" class="full-width" v-model="pwd2" /></td>\
						</tr>\
						<tr>\
							<td>Administrator:</td>\
							<td><input type="checkbox" name="admin" class="full-width" v-model="admin" /></td>\
						</tr>\
						<tr>\
							<td>Moderator:</td>\
							<td><input type="checkbox" name="moderator" class="full-width" v-model="moderator" /></td>\
						</tr>\
						<tr>\
							<td></td>\
							<td>\
								<input type="submit" value="Save" />\
								<button v-on:click.prevent="addEditUser = false">Cancel</button>\
							</td>\
						</tr>\
					</table>\
				</form>\
			</div>\
		</div>\
\
		<div class="box" v-if="removeUser">\
			<div class="wrapper">\
				<h2>Remove User</h2>\
\
				<msg :type="\'error\'" :msg="\'You have no permission to do this.\'" v-if="err == 200"></msg>\
\
				<p>Do you really want to remove this user?</p>\
\
				<button v-on:click="performRemoveUser()">Yes, remove user</button>\
				<button v-on:click="removeUser = false">Cancel</button>\
			</div>\
		</div>\
\
		<div class="box">\
			<div class="wrapper">\
				<msg :type="\'success\'" :msg="\'The user has been saved.\'" v-if="userSaved"></msg>\
				<msg :type="\'success\'" :msg="\'The user has been removed.\'" v-if="userRemoved"></msg>\
\
				<button v-on:click="openAddEditUser(0)">Add User</button>\
\
				<table>\
					<thead>\
						<tr>\
							<td class="w5">ID</td>\
							<td class="w30">Login</td>\
							<td class="w30">E-Mail</td>\
							<td class="w30">Role</td>\
							<td></td>\
						</tr>\
					</thead>\
					<tbody>\
						<tr v-for="u in user">\
							<td>{{u.id}}</td>\
							<td>{{u.login}}</td>\
							<td>{{u.email}}</td>\
							<td>\
								<span v-if="u.admin">Administrator</span>\
								<span v-if="u.moderator">Moderator</span>\
							</td>\
							<td>\
								<i class="fa fa-pencil" aria-hidden="true" title="Edit user" v-on:click="openAddEditUser(u.id)"></i>\
								<i class="fa fa-trash" aria-hidden="true" title="Remove user" v-on:click="openRemoveUser(u.id)"></i>\
							</td>\
						</tr>\
					</tbody>\
				</table>\
			</div>\
		</div>\
\
		<div class="box">\
			<div class="wrapper">\
				<strong>Administrators</strong> have full access and can change all settings (including changing user passwords).<br />\
				<strong>Moderators</strong> can start/stop server instances and create, change and remove configurations.<br />\
				Other users have read access, except for user E-Mail addresses.\
			</div>\
		</div>\
	</div>\
</div>',
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
			this.$http.get('/api/user')
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

				this.$http.get('/api/user', {params: {id: id}})
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
			this.$http.post('/api/user', {id: this._id,
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
			this.$http.delete('/api/user', {params: {id: this._id}})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this._reset();
				this._load();
				this.userRemoved = true;
			});
		}
	}
});
Vue.use(VueRouter);
Vue.use(VueResource);

var router = new VueRouter({
	routes: [
		{
			path: '/',
			component: Vue.component('Login')
		},
		{
			path: '/instance',
			component: Vue.component('Instance')
		},
		{
			path: '/configuration',
			component: Vue.component('Configuration')
		},
		{
			path: '/settings',
			component: Vue.component('Settings')
		},
		{
			path: '/user',
			component: Vue.component('User')
		},
		{
			path: '/about',
			component: Vue.component('About')
		},
		{
			path: '*',
			component: Vue.component('Dashboard')
		}
	]
});

// Simple login check with redirection when not logged in.
router.beforeEach(function(to, from, next){
	Vue.http.get('/api/session')
	.then(function(resp){
		// if not logged in and not on login page, redirect to login
		if(resp.data.code && to.path != '/'){
			next('/');
			return;
		}

		// if login or password page, but logged in, redirect to start page
		if(!resp.data.code && to.path == '/'){
			next('/instance');
			return;
		}

		SessionService.init(resp.data.data);
		next();
	});
});

window.onload = function(){
	new Vue({el: '#app', router: router});
};
