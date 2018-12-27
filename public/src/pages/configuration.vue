<template>
	<div>
		<hmenu></hmenu>

		<div class="main">
			<h1>Configurations</h1>
			
			<div class="box" v-if="addEditConfig">
				<div class="wrapper">
					<h2>Create/Edit Configuration</h2>

					<msg :type="'error'" :msg="'The name must be set.'" v-if="err == 1" v-on:close="closeMsg"></msg>
					<msg :type="'error'" :msg="'At least one weather configuration must be added.'" v-if="err == 2" v-on:close="closeMsg"></msg>
					<msg :type="'error'" :msg="'At least one car must be added.'" v-if="err == 3" v-on:close="closeMsg"></msg>
					<msg :type="'error'" :msg="'You have no permission to do this.'" v-if="err == 200" v-on:close="closeMsg"></msg>

					<form v-on:submit.prevent="performAddEditConfig()">
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseBasic = !collapseBasic">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Basic Settings</h3>
							</div>
							<table v-if="!collapseBasic">
								<tr>
									<td class="w20">Name:</td>
									<td><input type="text" name="name" class="full-width" v-model="name" /></td>
								</tr>
								<tr>
									<td>Password:</td>
									<td><input type="text" name="pwd" class="full-width" v-model="pwd" /></td>
								</tr>
								<tr>
									<td>Admin password:</td>
									<td><input type="text" name="admin_pwd" class="full-width" v-model="admin_pwd" /></td>
								</tr>
								<tr>
									<td>Pickup mode:</td>
									<td><input type="checkbox" name="pickup_mode" v-model="pickup_mode" /></td>
								</tr>
								<tr>
									<td>Lock entry list:</td>
									<td><input type="checkbox" name="lock_entry_list" v-model="lock_entry_list" /></td>
								</tr>
								<tr>
									<td>Race Pit Window Start:</td>
									<td><input type="number" name="race_pit_window_start" v-model="race_pit_window_start" /></td>
								</tr>
								<tr>
									<td>Race Pit Window End:</td>
									<td><input type="number" name="race_pit_window_end" v-model="race_pit_window_end" /></td>
								</tr>
								<tr>
									<td>Reversed Grid Race Postion:</td>
									<td><input type="number" name="reversed_grid_race_positions" v-model="reversed_grid_race_positions" /></td>
								</tr>
								<tr>
									<td>Loop mode:</td>
									<td><input type="checkbox" name="loop_mode" v-model="loop_mode" /></td>
								</tr>
								<tr>
									<td>Show on lobby:</td>
									<td><input type="checkbox" name="show_in_lobby" v-model="show_in_lobby" /></td>
								</tr>
								<tr>
									<td>Max. slots:</td>
									<td><input type="number" name="max_slots" v-model="max_slots" /></td>
								</tr>
								<tr>
									<td>Result screen time:</td>
									<td><input type="number" name="result_screen_time" v-model="result_screen_time" /></td>
								</tr>
								<tr>
									<td>Welcome message:</td>
									<td><input type="text" name="welcome" class="full-width" v-model="welcome" /></td>
								</tr>
								<tr>
									<td>Description:</td>
									<td><textarea name="description" v-model="description"></textarea></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseAdvanced = !collapseAdvanced">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Advanced Settings</h3>
							</div>
							<table v-if="!collapseAdvanced">
								<tr>
									<td>UDP port:</td>
									<td><input type="number" name="udp" v-model="udp" /></td>
								</tr>
								<tr>
									<td>TCP port:</td>
									<td><input type="number" name="tcp" v-model="tcp" /></td>
								</tr>
								<tr>
									<td>HTTP port:</td>
									<td><input type="number" name="http" v-model="http" /></td>
								</tr>
								<tr>
									<td>Packets Hz:</td>
									<td><input type="number" name="packets_hz" v-model="packets_hz" /></td>
								</tr>
								<tr>
									<td>Threads:</td>
									<td>
										<select name="threads" v-model="threads">
											<option>2</option>
											<option>3</option>
											<option>4</option>
											<option>5</option>
											<option>6</option>
											<option>7</option>
											<option>8</option>
										</select>
									</td>
								</tr>
								<tr>
									<td>Auth plugin address:</td>
									<td><input type="number" name="auth_plugin_address" v-model="auth_plugin_address" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseAssists = !collapseAssists">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Assists</h3>
							</div>
							<table v-if="!collapseAssists">
								<tr>
									<td>ABS:</td>
									<td>
										<select name="abs" v-model="abs">
											<option value="0">Denied</option>
											<option value="1">Factory</option>
											<option value="2">Forced</option>
										</select>
									</td>
								</tr>
								<tr>
									<td>TC:</td>
									<td>
										<select name="tc" v-model="tc">
											<option value="0">Denied</option>
											<option value="1">Factory</option>
											<option value="2">Forced</option>
										</select>
									</td>
								</tr>
								<tr>
									<td>Stability aid:</td>
									<td><input type="checkbox" name="stability_aid" v-model="stability_aid" /></td>
								</tr>
								<tr>
									<td>Auto clutch:</td>
									<td><input type="checkbox" name="auto_clutch" v-model="auto_clutch" /></td>
								</tr>
								<tr>
									<td>Tyre blankets:</td>
									<td><input type="checkbox" name="tyre_blankets" v-model="tyre_blankets" /></td>
								</tr>
								<tr>
									<td>Force virtual mirror:</td>
									<td><input type="checkbox" name="force_virtual_mirror" v-model="force_virtual_mirror" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseRealism = !collapseRealism">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Realism</h3>
							</div>
							<table v-if="!collapseRealism">
								<tr>
									<td>Fuel rate:</td>
									<td><input type="number" name="fuel_rate" v-model="fuel_rate" /></td>
								</tr>
								<tr>
									<td>Damage rate:</td>
									<td><input type="number" name="damage_rate" v-model="damage_rate" /></td>
								</tr>
								<tr>
									<td>Tires wear rate:</td>
									<td><input type="number" name="tires_wear_rate" v-model="tires_wear_rate" /></td>
								</tr>
								<tr>
									<td>Allowed tires out:</td>
									<td>
										<select name="allowed_tires_out" v-model="allowed_tires_out">
											<option value="-1">-1 (disabled)</option>
											<option>0</option>
											<option>1</option>
											<option>2</option>
											<option>3</option>
											<option>4</option>
										</select>
									</td>
								</tr>
								<tr>
									<td>Max ballast:</td>
									<td><input type="number" name="max_ballast" v-model="max_ballast" /></td>
								</tr>
								<tr>
									<td>Disable gas cut penality:</td>
									<td><input type="checkbox" name="disable_gas_cut_penality" v-model="disable_gas_cut_penality" /></td>
								</tr>
								<tr>
									<td>Jump Start:</td>
									<select name="start_rule" v-model="start_rule">
										<option value="0">Car Locked</option>
										<option value="1">Teleport To Pit</option>
										<option value="2">Drive-through</option>
									</select>
								</tr>
								<tr>
									<td>Legal Tyres:</td>
									<td><input type="text" name="legal_tyres" v-model="legal_tyres" /></td>
								</tr>
								<tr>
									<td>UDP Plugin Local Port:</td>
									<td><input type="number" name="udp_plugin_local_port" v-model="udp_plugin_local_port" /></td>
								</tr>
								<tr>
									<td>UDP Plugin Address:</td>
									<td><input type="text" name="udp_plugin_address" v-model="udp_plugin_address" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseDynamicTrack = !collapseDynamicTrack">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Dynamic Track</h3>
							</div>
							<table v-if="!collapseDynamicTrack">
								<tr>
									<td>Dynamic track:</td>
									<td><input type="checkbox" name="dynamic_track" v-model="dynamic_track" /></td>
								</tr>
								<tr>
									<td>Condition:</td>
									<td>
										<select name="condition" v-model="condition">
											<option>CUSTOM</option>
											<option>DUSTY</option>
											<option>OLD</option>
											<option>SLOW</option>
											<option>GREEN</option>
											<option>FAST</option>
											<option>OPTIMUM</option>
										</select>
									</td>
								</tr>
								<tr>
									<td>Start value:</td>
									<td><input type="number" name="start_value" v-model="start_value" /></td>
								</tr>
								<tr>
									<td>Randomness:</td>
									<td><input type="number" name="randomness" v-model="randomness" /></td>
								</tr>
								<tr>
									<td>Transferred grip:</td>
									<td><input type="number" name="transferred_grip" v-model="transferred_grip" /></td>
								</tr>
								<tr>
									<td>Laps to improve grip:</td>
									<td><input type="number" name="laps_to_improve_grip" v-model="laps_to_improve_grip" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseVoting = !collapseVoting">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Voting</h3>
							</div>
							<table v-if="!collapseVoting">
								<tr>
									<td>Kick vote quorum:</td>
									<td><input type="number" name="kick_vote_quorum" v-model="kick_vote_quorum" /></td>
								</tr>
								<tr>
									<td>Session vote quorum:</td>
									<td><input type="number" name="session_vote_quorum" v-model="session_vote_quorum" /></td>
								</tr>
								<tr>
									<td>Vote duration:</td>
									<td><input type="number" name="vote_duration" v-model="vote_duration" /></td>
								</tr>
								<tr>
									<td>Blacklist:</td>
									<td>
										<select name="blacklist" v-model="blacklist">
											<option value="0">Kick Player</option>
											<option value="1">Kick Until Restart</option>
										</select>
									</td>
								</tr>
								<tr>
									<td>Max. collisions per KM:</td>
									<td><input type="number" name="max_collisions_km" v-model="max_collisions_km" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseBooking = !collapseBooking">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Booking</h3>
							</div>
							<table v-if="!collapseBooking">
								<tr>
									<td>Booking:</td>
									<td><input type="checkbox" name="booking" v-model="booking" /></td>
								</tr>
								<tr>
									<td>Booking time:</td>
									<td><input type="number" name="booking_time" v-model="booking_time" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapsePractice = !collapsePractice">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Practice</h3>
							</div>
							<table v-if="!collapsePractice">
								<tr>
									<td>Practice:</td>
									<td><input type="checkbox" name="practice" v-model="practice" /></td>
								</tr>
								<tr>
									<td>Practice time:</td>
									<td><input type="number" name="practice_time" v-model="practice_time" /></td>
								</tr>
								<tr>
									<td>Can join:</td>
									<td><input type="checkbox" name="can_join_practice" v-model="can_join_practice" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseQualify = !collapseQualify">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Qualify</h3>
							</div>
							<table v-if="!collapseQualify">
								<tr>
									<td>Qualify:</td>
									<td><input type="checkbox" name="qualify" v-model="qualify" /></td>
								</tr>
								<tr>
									<td>Qualify time:</td>
									<td><input type="number" name="qualify_time" v-model="qualify_time" /></td>
								</tr>
								<tr>
									<td>Can join:</td>
									<td><input type="checkbox" name="can_join_qualify" v-model="can_join_qualify" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseRace = !collapseRace">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Race</h3>
							</div>
							<table v-if="!collapseRace">
								<tr>
									<td>Race:</td>
									<td><input type="checkbox" name="race" v-model="race" /></td>
								</tr>
								<tr>
									<td>Race laps:</td>
									<td><input type="number" name="race_laps" v-model="race_laps" /></td>
								</tr>
								<tr>
									<td>Race time:</td>
									<td><input type="number" name="race_time" v-model="race_time" /></td>
								</tr>
								<tr>
									<td>Race overtime:</td>
									<td><input type="number" name="race_overtime" v-model="race_overtime" /></td>
								</tr>
								<tr>
									<td>Race wait time:</td>
									<td><input type="number" name="race_wait_time" v-model="race_wait_time" /></td>
								</tr>
								<tr>
									<td>Race extra lap:</td>
									<td><input type="checkbox" name="race_extra_lap" v-model="race_extra_lap" /></td>
								</tr>
								<tr>
									<td>Join type:</td>
									<td>
										<select name="join_type" v-model="join_type">
											<option value="0">Close</option>
											<option value="1">Open</option>
											<option value="2">Close at Start</option>
										</select>
									</td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseTime = !collapseTime">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Time</h3>
							</div>
							<table v-if="!collapseTime">
								<tr>
									<td>Time:</td>
									<td>
										<input type="time" name="time" v-model="time" step="1800" min="08:00" max="18:00" />
										(08:00 - 18:00, 30 minute steps)
									</td>
								</tr>
								<tr>
									<td>Sun angle:</td>
									<td><input type="number" name="sun_angle" v-model="sun_angle" readonly /></td>
								</tr>
								<tr>
									<td>Multiplier:</td>
									<td><input type="number" name="time_of_day_mult" v-model="time_of_day_mult" min="1" max="10" step="1" /></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseWeather = !collapseWeather">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Weather</h3>
							</div>
							<table v-if="!collapseWeather">
								<tr v-for="(w, i) in weather">
									<td colspan="2">
										<table>
											<tr>
												<td>Weather:</td>
												<td>
													<select name="weather" v-model="w.weather">
														<option value="3_clear">Clear</option>
														<option value="7_heavy_clouds">Heavy Clouds</option>
														<option value="1_heavy_fog">Heavy Fog</option>
														<option value="5_light_clouds">Light Clouds</option>
														<option value="2_light_fog">Light Fog</option>
														<option value="4_mid_clear">Mid Clear</option>
														<option value="6_mid_clouds">Mid Clouds</option>
													</select>
												</td>
												<td>
													<button v-on:click.prevent="removeWeather(i)">Remove Weather Panel</button>
												</td>
											</tr>
											<tr>
												<td>Base ambient temp:</td>
												<td><input type="number" name="base_ambient_temp" v-model="w.base_ambient_temp" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Base road temp:</td>
												<td><input type="number" name="base_road_temp" v-model="w.base_road_temp" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Ambient variation:</td>
												<td><input type="number" name="ambient_variation" v-model="w.ambient_variation" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Road variation:</td>
												<td><input type="number" name="road_variation" v-model="w.road_variation" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Wind Base Speed Min:</td>
												<td><input type="number" name="wind_base_speed_min" v-model="w.wind_base_speed_min" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Wind Base Speed Max:</td>
												<td><input type="number" name="wind_base_speed_max" v-model="w.wind_base_speed_max" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Wind Base Direction:</td>
												<td><input type="number" name="wind_base_direction" v-model="w.wind_base_direction" /></td>
												<td></td>
											</tr>
											<tr>
												<td>Wind Variation Direction:</td>
												<td><input type="number" name="wind_variation_direction" v-model="w.wind_variation_direction" /></td>
												<td></td>
											</tr>
										</table>
									</td>
								</tr>
								<tr>
									<td colspan="2"><button v-on:click.prevent="addWeather()">Add Weather Panel</button></td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseTrack = !collapseTrack">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Track</h3>
							</div>
							<table v-if="!collapseTrack">
								<tr>
									<td colspan="2">
										<div class="select">
											<div class="option" v-for="(track, i) in tracks" v-on:click="selectTrack(i)" v-bind:class="{selected: selectedTrack == i}">
												<strong>{{track.name}}</strong>
												({{track.max_slots}})
												{{track.description}}
											</div>
										</div>
									</td>
								</tr>
							</table>
						</div>
						<div class="box no-border">
							<div class="collapse" v-on:click="collapseCars = !collapseCars">
								<h3><i class="fa fa-plus" aria-hidden="true"></i> Cars/Entry List</h3>
							</div>
							<table v-if="!collapseCars">
								<tr>
									<td colspan="2">
										<table>
											<tr>
												<td class="w33">
													<strong>Car</strong>

													<div class="select">
														<div class="option" v-for="(car, i) in cars" v-on:click="selectCar(i)" v-bind:class="{selected: selectedCar == i}">
															<strong>{{car.name}}</strong>
															{{car.description}}
														</div>
													</div>
												</td>
												<td class="w33">
													<strong>Painting</strong>

													<div class="select">
														<div class="option" v-for="(painting, i) in activePaintings" v-on:click="selectPainting(i)" v-bind:class="{selected: selectedPainting == i}">
															<strong>{{painting}}</strong>
														</div>
													</div>
												</td>
												<td class="w33 top">
													<strong>Slot Settings</strong>

													<table>
														<tr>
															<td class="w20">Spectator:</td>
															<td><input type="checkbox" name="spectator" v-model="spectator" /></td>
														</tr>
														<tr>
															<td>Driver:</td>
															<td><input type="text" name="driver" v-model="driver" /></td>
														</tr>
														<tr>
															<td>Team:</td>
															<td><input type="text" name="team" v-model="team" /></td>
														</tr>
														<tr>
															<td>GUID:</td>
															<td><input type="text" name="guid" v-model="guid" /></td>
														</tr>
														<tr>
															<td>Fixed Setup:</td>
															<td><input type="text" name="fixed_setup" v-model="fixed_setup" /></td>
														</tr>
														<tr>
															<td>Ballast (kg):</td>
															<td><input type="text" name="ballast" v-model="ballast" /></td>
														</tr>
														<tr>
															<td>Restrictor (0-100, %):</td>
															<td><input type="text" name="restrictor" v-model="restrictor" /></td>
														</tr>
														<tr>
															<td></td>
															<td><button v-on:click.prevent="addCar()">Add Car</button></td>
														</tr>
													</table>
												</td>
											</tr>
										</table>
									</td>
								</tr>
								<tr>
									<td colspan="2">
										<table>
											<thead>
												<tr>
													<td class="w20">Car</td>
													<td class="w20">Painting</td>
													<td class="w5">Spectator</td>
													<td class="w10">Driver</td>
													<td class="w10">Team</td>
													<td class="w10">GUID</td>
													<td class="w20">Fixed Setup</td>
													<td class="w10">Ballast</td>
													<td class="w10">Restrictor</td>
													<td></td>
												</tr>
											</thead>
											<tbody>
												<tr v-for="(car, i) in selectedCars">
													<td>{{car.car}}</td>
													<td>{{car.painting}}</td>
													<td>{{car.spectator}}</td>
													<td>{{car.driver}}</td>
													<td>{{car.team}}</td>
													<td>{{car.guid}}</td>
													<td>{{car.fixed_setup}}</td>
													<td>{{car.ballast}}kg</td>
													<td>{{car.restrictor}}%</td>
													<td>
														<i class="fa fa-angle-up" aria-hidden="true" title="Move up" v-on:click.prevent="carUp(i)"></i>
														<i class="fa fa-angle-down" aria-hidden="true" title="Move down" v-on:click.prevent="carDown(i)"></i>
														<i class="fa fa-trash" aria-hidden="true" title="Remove car" v-on:click.prevent="removeCar(i)"></i>
													</td>
												</tr>
											</tbody>
										</table>
									</td>
								</tr>
							</table>
						</div>
						<table>
							<tr>
								<td colspan="2"><hr /></td>
							</tr>
							<tr>
								<td>
									Custom server_cfg.ini:<br />
									(will be appended to the end of server_cfg.ini)
								</td>
								<td><textarea v-model="server_cfg_ini"></textarea></td>
							</tr>
							<tr>
								<td>
									Custom entry_list.ini:<br />
									(will be appended to the end of entry_list.ini)
								</td>
								<td><textarea v-model="entry_list_ini"></textarea></td>
							</tr>
						</table>

						<msg :type="'error'" :msg="'The name must be set.'" v-if="err == 1" v-on:close="closeMsg"></msg>
						<msg :type="'error'" :msg="'At least one weather configuration must be added.'" v-if="err == 2" v-on:close="closeMsg"></msg>
						<msg :type="'error'" :msg="'At least one car must be added.'" v-if="err == 3" v-on:close="closeMsg"></msg>
						<msg :type="'error'" :msg="'You have no permission to do this.'" v-if="err == 200" v-on:close="closeMsg"></msg>

						<input type="submit" value="Save" />
						<button v-on:click.prevent="addEditConfig = false">Cancel</button>
					</form>
				</div>
			</div>

			<div class="box" v-if="removeConfig">
				<div class="wrapper">
					<h2>Remove Configuration</h2>

					<msg :type="'error'" :msg="'You have no permission to do this.'" v-if="err == 200" v-on:close="closeMsg"></msg>

					<p>Do you really want to remove this configuration? This won't stop any instances currently running.</p>

					<button v-on:click="performRemoveConfig()">Yes, remove configuration</button>
					<button v-on:click="removeConfig = false">Cancel</button>
				</div>
			</div>

			<div class="box" v-if="importConfig">
				<div class="wrapper">
					<h2>Import Configuration</h2>
					<p>Please select the following config files in the file upload dialog (you can select multiple files):</p>
					<p>
						<i class="fa fa-check-square-o" v-if="tmpSrvCfg !== null"></i>
						<i class="fa fa-square-o" v-else></i>
						<code>server_cfg.ini</code>
						<br/>
						<i class="fa fa-check-square-o" v-if="tmpEntryListCfg !== null"></i>
						<i class="fa fa-square-o" v-else></i>
						<code>entry_list.ini</code>
					</p>

					<p>
						<input type="file" multiple @change="importConfigs($event.target.files)">
					</p>

					<button v-on:click="closeImportConfig()">Cancel</button>
				</div>
			</div>

			<div class="box">
				<div class="wrapper">
					<msg :type="'success'" :msg="'The configuration has been saved.'" v-if="saved" v-on:close="saved = false"></msg>
					<msg :type="'success'" :msg="'The configuration has been removed.'" v-if="removed" v-on:close="removed = false"></msg>

					<button v-on:click="openAddEditConfig(0)">Add Configuration</button>
					<button v-on:click="openImportConfig()">Import Configuration</button>

					<table>
						<thead>
							<tr>
								<td class="w5">ID</td>
								<td class="w40">Name</td>
								<td class="w40">Info</td>
								<td></td>
							</tr>
						</thead>
						<tbody>
							<tr v-for="config in configs">
								<td>{{config.id}}</td>
								<td>{{config.name}}</td>
								<td>
									{{config.track}}<br />
									Ports: {{config.tcp}} TCP/{{config.udp}} UDP
								</td>
								<td>
									<i class="fa fa-pencil" aria-hidden="true" title="Edit configuration" v-on:click="openAddEditConfig(config.id)"></i>
									<i class="fa fa-files-o" aria-hidden="true" title="Copy configuration" v-on:click="copyConfig(config.id)"></i>
									<a class="icon-link" v-bind:href="generateCfgDownloadUrl(config.id)">
										<i class="fa fa-cloud-download" aria-hidden="true" title="Download configuration"></i>
									</a>
									<i class="fa fa-trash" aria-hidden="true" title="Remove configuration" v-on:click="openRemoveConfig(config.id)"></i>
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
			collapseBasic: true,
			collapseAdvanced: true,
			collapseAssists: true,
			collapseRealism: true,
			collapseDynamicTrack: true,
			collapseVoting: true,
			collapseBooking: true,
			collapsePractice: true,
			collapseQualify: true,
			collapseRace: true,
			collapseTime: true,
			collapseWeather: true,
			collapseTrack: true,
			collapseCars: true,
			// ---
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
			fixed_setup: '',
			ballast: 0,
			restrictor: 0,
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
			auth_plugin_address: 0,
			abs: 1,
			tc: 1,
			stability_aid: false,
			auto_clutch: false,
			tyre_blankets: true,
			force_virtual_mirror: true,
			fuel_rate: 100,
			damage_rate: 50,
			tires_wear_rate: 100,
			allowed_tires_out: 2,
			max_ballast: 150,
			start_rule: 1,
			disable_gas_cut_penality: false,
			dynamic_track: true,
			condition: 'CUSTOM',
			start_value: 100,
			randomness: 0,
			transferred_grip: 100,
			laps_to_improve_grip: 1,
			kick_vote_quorum: 70,
			session_vote_quorum: 70,
			vote_duration: 15,
			blacklist: 0,
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
			join_type: 1,
			time: '14:00',
			sun_angle: 16,
			time_of_day_mult: 1,
			track: '',
			legal_tyres: '',
			udp_plugin_local_port: 0,
			udp_plugin_address: '',
			race_pit_window_start: 0,
			race_pit_window_end: 0,
			reversed_grid_race_positions: 0,
			server_cfg_ini: '',
			entry_list_ini: '',
			// ---
			err: 0,
			addEditConfig: false,
			removeConfig: false,
			importConfig: false,
			saved: false,
			removed: false,
			// ---
			tmpSrvCfg: null,
			tmpEntryListCfg: null
		}
	},
	mounted() { 
		this._load();
	},
	watch: {
		condition: function (value) {
			this.populateDynamicTrackWithPreset(value);
		},
		time: function (value) {
			this.calculateSunAngleByTime(value);
		}
	},
	methods: {
		_load() { 
			axios.get('/api/configuration')
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.configs = resp.data;
			});

			axios.get('/api/tracks')
			.then(resp => {
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.tracks = resp.data;
			});

			axios.get('/api/cars')
			.then(resp => {
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
		_reset() { 
			this.spectator = false;
			this.driver = '';
			this.team = '';
			this.guid = '';
			this.fixed_setup = '';
			this.ballast = 0;
			this.restrictor = 0;

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
			this.abs = 1;
			this.tc = 1;
			this.stability_aid = false;
			this.auto_clutch = false;
			this.tyre_blankets = true;
			this.force_virtual_mirror = true;
			this.fuel_rate = 100;
			this.damage_rate = 50;
			this.tires_wear_rate = 100;
			this.allowed_tires_out = 2;
			this.max_ballast = 150;
			this.start_rule = 1;
			this.disable_gas_cut_penality = false;
			this.dynamic_track = true;
			this.condition = 'CUSTOM';
			this.start_value = 100;
			this.randomness = 0;
			this.transferred_grip = 100;
			this.laps_to_improve_grip = 1;
			this.kick_vote_quorum = 70;
			this.session_vote_quorum = 70;
			this.vote_duration = 15;
			this.blacklist = 0;
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
			this.join_type = 1;
			this.time = '14:00';
			this.sun_angle = 16;
			this.time_of_day_mult = 1;
			this.track = '';
			this.legal_tyres = '';
			this.udp_plugin_local_port = 0;
			this.udp_plugin_address = '';
			this.race_pit_window_start = 0;
			this.race_pit_window_end = 0;
			this.reversed_grid_race_positions = 0;
			this.server_cfg_ini = '';
			this.entry_list_ini = '';
			
			this.err = 0;
			this.addEditConfig = false;
			this.removeConfig = false;
			this.importConfig = false;
			this.saved = false;
			this.removed = false;
		},
		_openConfig: function(id, copy){
			axios.get('/api/configuration', {params: {id: id}})
			.then(resp => {
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
				this.start_rule = resp.data.start_rule;
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
				this.sun_angle = resp.data.sun_angle;
				this.time_of_day_mult = resp.data.time_of_day_mult;
				this.time = this.calculateTimeBySunAngle(this.sun_angle);
				this.legal_tyres = resp.data.legal_tyres;
				this.udp_plugin_local_port = resp.data.udp_plugin_local_port;
				this.udp_plugin_address = resp.data.udp_plugin_address;
				this.race_pit_window_start = resp.data.race_pit_window_start;
				this.race_pit_window_end = resp.data.race_pit_window_end;
				this.reversed_grid_race_positions = resp.data.reversed_grid_race_positions;
				this.server_cfg_ini = resp.data.server_cfg_ini;
				this.entry_list_ini = resp.data.entry_list_ini;

				if(copy){
					this.name += ' (copy)';
				}

				// track
				this.findAndSelectTrack(resp.data.track, resp.data.track_config);

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

		importConfigs: function (files) {
			// https://stackoverflow.com/a/12452845
			var parseINIString = function (data) {
				var regex = {
					section: /^\s*\[\s*([^\]]*)\s*\]\s*$/,
					param: /^\s*([^=]+?)\s*=\s*(.*?)\s*$/,
					comment: /^\s*;.*$/
				};
				var value = {};
				var lines = data.split(/[\r\n]+/);
				var section = null;
				lines.forEach(function (line) {
					var match;
					if (regex.comment.test(line)) {
						// Ignore comments
					} else if (regex.param.test(line)) {
						match = line.match(regex.param);
						if (section) {
							value[section][match[1]] = match[2];
						} else {
							value[match[1]] = match[2];
						}
					} else if (regex.section.test(line)) {
						match = line.match(regex.section);
						value[match[1]] = {};
						section = match[1];
					} else if (line.length === 0 && section) {
						section = null;
					}
				});

				return value;
			};

			for (var i = 0; i < files.length; i++) {
				var reader = new FileReader();

				reader.onerror = function () {
					alert('Error while loading the config files');
					this.closeImportConfig();
				}.bind(this);

				reader.onload = function (e) {
					var ini = parseINIString(e.target.result);

					if ('SERVER' in ini) {
						this.tmpSrvCfg = ini;
					} else if ('CAR_0' in ini) {
						this.tmpEntryListCfg = ini;
					}

					this.checkImportableConfigs();
				}.bind(this);

				reader.readAsText(files[i]);
			}
		},
		checkImportableConfigs: function () {
			if (!this.tmpSrvCfg || !this.tmpEntryListCfg) {
				return;
			}

			this.openAddEditConfig(0);

			sToB = function (s) { return s === '1' };

			this.name = this.tmpSrvCfg.SERVER.NAME;
			this.pwd = this.tmpSrvCfg.SERVER.PASSWORD;
			this.admin_pwd = this.tmpSrvCfg.SERVER.ADMIN_PASSWORD;
			this.pickup_mode = sToB(this.tmpSrvCfg.SERVER.PICKUP_MODE_ENABLED);
			this.max_slots = this.tmpSrvCfg.SERVER.MAX_CLIENTS;
			this.udp = this.tmpSrvCfg.SERVER.UDP_PORT;
			this.tcp = this.tmpSrvCfg.SERVER.TCP_PORT;
			this.http = this.tmpSrvCfg.SERVER.HTTP_PORT;
			this.packets_hz = this.tmpSrvCfg.SERVER.CLIENT_SEND_INTERVAL_HZ;
			this.show_in_lobby = sToB(this.tmpSrvCfg.SERVER.REGISTER_TO_LOBBY);
			this.loop_mode = sToB(this.tmpSrvCfg.SERVER.LOOP_MODE);
			this.threads = this.tmpSrvCfg.SERVER.NUM_THREADS;
			this.abs = this.tmpSrvCfg.SERVER.ABS_ALLOWED;
			this.tc = this.tmpSrvCfg.SERVER.TC_ALLOWED;
			this.welcome = this.tmpSrvCfg.SERVER.WELCOME_MESSAGE;
			this.fuel_rate = this.tmpSrvCfg.SERVER.FUEL_RATE;
			this.damage_rate = this.tmpSrvCfg.SERVER.DAMAGE_MULTIPLIER;
			this.tires_wear_rate = this.tmpSrvCfg.SERVER.TYRE_WEAR_RATE;
			this.auto_clutch = sToB(this.tmpSrvCfg.SERVER.AUTOCLUTCH_ALLOWED);
			this.stability_aid = sToB(this.tmpSrvCfg.SERVER.STABILITY_ALLOWED);
			this.tyre_blankets = sToB(this.tmpSrvCfg.SERVER.TYRE_BLANKETS_ALLOWED);
			this.force_virtual_mirror = sToB(this.tmpSrvCfg.SERVER.FORCE_VIRTUAL_MIRROR);
			this.result_screen_time = this.tmpSrvCfg.SERVER.RESULT_SCREEN_TIME;
			this.kick_vote_quorum = this.tmpSrvCfg.SERVER.KICK_QUORUM;
			this.session_vote_quorum = this.tmpSrvCfg.SERVER.VOTING_QUORUM;
			this.vote_duration = this.tmpSrvCfg.SERVER.VOTE_DURATION;
			this.max_ballast = this.tmpSrvCfg.SERVER.MAX_BALLAST_KG;
			this.start_rule = this.tmpSrvCfg.SERVER.START_RULE;
			this.disable_gas_cut_penality = sToB(this.tmpSrvCfg.SERVER.RACE_GAS_PENALTY_DISABLED);
			this.sun_angle = this.tmpSrvCfg.SERVER.SUN_ANGLE;
			this.time_of_day_mult = this.tmpSrvCfg.SERVER.TIME_OF_DAY_MULT;
			this.legal_tyres = this.tmpSrvCfg.SERVER.LEGAL_TYRES;
			this.udp_plugin_local_port = this.tmpSrvCfg.SERVER.UDP_PLUGIN_LOCAL_PORT;
			this.udp_plugin_address = this.tmpSrvCfg.SERVER.UDP_PLUGIN_ADDRESS;
			this.allowed_tires_out = this.tmpSrvCfg.SERVER.ALLOWED_TYRES_OUT;
			this.max_collisions_km = this.tmpSrvCfg.SERVER.MAX_CONTACTS_PER_KM;
			this.blacklist = this.tmpSrvCfg.SERVER.BLACKLIST_MODE;
			this.description = this.tmpSrvCfg.DATA.DESCRIPTION;

			if (this.booking = this.tmpSrvCfg.BOOK !== undefined) {
				this.booking_time = this.tmpSrvCfg.BOOK.TIME;
			}
			if (this.practice = this.tmpSrvCfg.PRACTICE !== undefined) {
				this.practice_time = this.tmpSrvCfg.PRACTICE.TIME;
				this.can_join_practice = sToB(this.tmpSrvCfg.PRACTICE.IS_OPEN);
			}
			if (this.qualify = this.tmpSrvCfg.QUALIFY !== undefined) {
				this.qualify_time = this.tmpSrvCfg.QUALIFY.TIME;
				this.can_join_qualify = sToB(this.tmpSrvCfg.QUALIFY.IS_OPEN);
			}
			if (this.race = this.tmpSrvCfg.RACE !== undefined) {
				this.race_time = this.tmpSrvCfg.RACE.TIME;
				this.race_laps = this.tmpSrvCfg.RACE.LAPS;
				this.race_wait_time = this.tmpSrvCfg.RACE.WAIT_TIME;
				this.join_type = this.tmpSrvCfg.RACE.IS_OPEN;
			}
			this.race_extra_lap = this.tmpSrvCfg.RACE_EXTRA_LAP;
			this.race_overtime = this.tmpSrvCfg.SERVER.RACE_OVER_TIME;
			this.race_pit_window_start = this.tmpSrvCfg.SERVER.RACE_PIT_WINDOW_START;
			this.race_pit_window_end = this.tmpSrvCfg.SERVER.RACE_PIT_WINDOW_END;
			this.reversed_grid_race_positions = this.tmpSrvCfg.SERVER.REVERSED_GRID_RACE_POSITIONS;

			if (this.dynamic_track = this.tmpSrvCfg.DYNAMIC_TRACK !== undefined) {
				this.start_value = this.tmpSrvCfg.DYNAMIC_TRACK.SESSION_START;
				this.randomness = this.tmpSrvCfg.DYNAMIC_TRACK.RANDOMNESS;
				this.transferred_grip = this.tmpSrvCfg.DYNAMIC_TRACK.SESSION_TRANSFER;
				this.laps_to_improve_grip = this.tmpSrvCfg.DYNAMIC_TRACK.LAP_GAIN;
			}

			for (var wi = 0; wi < 25; wi++) {
				var key = 'WEATHER_' + wi;
				if (!(key in this.tmpSrvCfg)) {
					break;
				}

				this.weather.push({
					weather: this.tmpSrvCfg[key].GRAPHICS,
					base_ambient_temp: this.tmpSrvCfg[key].BASE_TEMPERATURE_AMBIENT,
					base_road_temp: this.tmpSrvCfg[key].BASE_TEMPERATURE_ROAD,
					ambient_variation: this.tmpSrvCfg[key].VARIATION_AMBIENT,
					road_variation: this.tmpSrvCfg[key].VARIATION_ROAD,
					wind_base_speed_min: this.tmpSrvCfg[key].WIND_BASE_SPEED_MIN,
					wind_base_speed_max: this.tmpSrvCfg[key].WIND_BASE_SPEED_MAX,
					wind_base_direction: this.tmpSrvCfg[key].WIND_BASE_DIRECTION,
					wind_variation_direction: this.tmpSrvCfg[key].WIND_VARIATION_DIRECTION
				});
			}

			var unknownTrack = null;
			if (!this.findAndSelectTrack(this.tmpSrvCfg.SERVER.TRACK, this.tmpSrvCfg.SERVER.CONFIG_TRACK)) {
				unknownTrack = this.tmpSrvCfg.SERVER.TRACK;
				if (this.tmpSrvCfg.SERVER.CONFIG_TRACK) {
					unknownTrack += ' (' + this.tmpSrvCfg.SERVER.CONFIG_TRACK + ')';
				}
			}

			var unknownCars = [];
			Object.keys(this.tmpEntryListCfg).forEach(function (car_idx) {
				var carIsUnknown = unknownCars.indexOf(this.tmpEntryListCfg[car_idx].MODEL) > -1;

				if (carIsUnknown) {
					return;
				}

				for (var ci = 0; ci < this.cars.length; ci++) {
					if (this.cars[ci].name === this.tmpEntryListCfg[car_idx].MODEL) {
						this.selectedCars.push({
							car: this.tmpEntryListCfg[car_idx].MODEL,
							painting: this.tmpEntryListCfg[car_idx].SKIN,
							spectator: sToB(this.tmpEntryListCfg[car_idx].SPECTATOR_MODE),
							driver: this.tmpEntryListCfg[car_idx].DRIVERNAME,
							team: this.tmpEntryListCfg[car_idx].TEAM,
							guid: this.tmpEntryListCfg[car_idx].GUID,
							fixed_setup: this.tmpEntryListCfg[car_idx].FIXED_SETUP,
							ballast: this.tmpEntryListCfg[car_dx].BALLAST,
							restrictor: this.tmpEntryListCfg[car_dx].RESTRICTOR,
							position: this.selectedCars.length
						});

						return;
					}
				}

				if (!carIsUnknown) {
					unknownCars.push(this.tmpEntryListCfg[car_idx].MODEL);
				}
			}.bind(this));

			this.closeImportConfig();

			if (unknownTrack) {
				alert('Unknown track: ' + unknownTrack);
			}
			if (unknownCars.length > 0) {
				alert('Unknown cars: ' + unknownCars.join(', '));
			}
		},
		openImportConfig: function () {
			this.importConfig = true;
			this.tmpSrvCfg = this.tmpEntryListCfg = null;
		},
		closeImportConfig: function () {
			this.importConfig = false;
			this.tmpSrvCfg = this.tmpEntryListCfg = null;
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
		performAddEditConfig() { 
			for(var i = 0; i < this.weather.length; i++){
				this.weather[i].base_ambient_temp = parseInt(this.weather[i].base_ambient_temp);
				this.weather[i].base_road_temp = parseInt(this.weather[i].base_road_temp);
				this.weather[i].ambient_variation = parseInt(this.weather[i].ambient_variation);
				this.weather[i].road_variation = parseInt(this.weather[i].road_variation);
				this.weather[i].wind_base_speed_min = parseInt(this.weather[i].wind_base_speed_min);
				this.weather[i].wind_base_speed_max = parseInt(this.weather[i].wind_base_speed_max);
				this.weather[i].wind_base_direction = parseInt(this.weather[i].wind_base_direction);
				this.weather[i].wind_variation_direction = parseInt(this.weather[i].wind_variation_direction);
			}

			for(var i = 0; i < this.selectedCars.length; i++){
				this.selectedCars[i].position = i;
				this.selectedCars[i].ballast = parseInt(this.selectedCars[i].ballast);
				this.selectedCars[i].restrictor = parseInt(this.selectedCars[i].restrictor);
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
				abs: parseInt(this.abs),
				tc: parseInt(this.tc),
				stability_aid: this.stability_aid,
				auto_clutch: this.auto_clutch,
				tyre_blankets: this.tyre_blankets,
				force_virtual_mirror: this.force_virtual_mirror,
				fuel_rate: parseInt(this.fuel_rate),
				damage_rate: parseInt(this.damage_rate),
				tires_wear_rate: parseInt(this.tires_wear_rate),
				allowed_tires_out: parseInt(this.allowed_tires_out),
				max_ballast: parseInt(this.max_ballast),
				start_rule: parseInt(this.start_rule),
				disable_gas_cut_penality: this.disable_gas_cut_penality,
				time_of_day_mult: parseInt(this.time_of_day_mult),
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
				blacklist: parseInt(this.blacklist),
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
				join_type: parseInt(this.join_type),
				sun_angle: parseInt(this.sun_angle),
				time: this.calculateTimeBySunAngle(parseInt(this.sun_angle)),
				weather: this.weather,
				track: this.track.name,
				track_config: this.track.config,
				legal_tyres: this.legal_tyres,
				udp_plugin_local_port: parseInt(this.udp_plugin_local_port),
				udp_plugin_address: this.udp_plugin_address,
				race_pit_window_start: parseInt(this.race_pit_window_start),
				race_pit_window_end: parseInt(this.race_pit_window_end),
				reversed_grid_race_positions: parseInt(this.reversed_grid_race_positions),
				server_cfg_ini: this.server_cfg_ini,
				entry_list_ini: this.entry_list_ini,
				cars: this.selectedCars
			};

			axios.post('/api/configuration', data)
			.then(resp => {
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
		performRemoveConfig() { 
			axios.delete('/api/configuration', {params: {id: this._id}})
			.then(resp => {
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
		addWeather() { 
			this.weather.push({
				weather: 'Clear',
				base_ambient_temp: 20,
				base_road_temp: 18,
				ambient_variation: 1,
				road_variation: 1,
				wind_base_speed_min: 0,
				wind_base_speed_max: 0,
				wind_base_direction: 0,
				wind_variation_direction: 0
			});
		},
		removeWeather: function(i){
			this.weather.splice(i, 1);
		},
		selectTrack: function(i){
			this.selectedTrack = i;
			this.track = this.tracks[i];
		},
		findAndSelectTrack: function (name, config_name) {
			for (var i = 0; i < this.tracks.length; i++) {
				if (this.tracks[i].name === name && (!config_name || this.tracks[i].config === config_name)) {
					this.selectTrack(i);
					return true;
				}
			}

			return false;
		},
		selectCar: function(i){
			this.selectedCar = i;
			this.selectedPainting = 0;
			this.activePaintings = this.cars[i].paintings;
		},
		selectPainting: function(i){
			this.selectedPainting = i;
		},
		addCar() { 
			var car = this.cars[this.selectedCar];
			this.selectedCars.push({
				car: car.name,
				painting: car.paintings[this.selectedPainting],
				spectator: this.spectator,
				driver: this.driver,
				team: this.team,
				guid: this.guid,
				position: this.selectedCars.length,
				fixed_setup: this.fixed_setup,
				ballast: this.ballast,
				restrictor: this.restrictor
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
		},
		populateDynamicTrackWithPreset: function(preset) {
			switch(preset) {
				case 'DUSTY':
					this.start_value = 86;
					this.randomness = 1;
					this.transferred_grip = 50;
					this.laps_to_improve_grip = 30;
					break;
				case 'OLD':
					this.start_value = 89;
					this.randomness = 3;
					this.transferred_grip = 80;
					this.laps_to_improve_grip = 50;
					break;
				case 'SLOW':
					this.start_value = 96;
					this.randomness = 1;
					this.transferred_grip = 80;
					this.laps_to_improve_grip = 300;
					break;
				case 'GREEN':
					this.start_value = 95;
					this.randomness = 2;
					this.transferred_grip = 90;
					this.laps_to_improve_grip = 132;
					break;
				case 'FAST':
					this.start_value = 98;
					this.randomness = 2;
					this.transferred_grip = 80;
					this.laps_to_improve_grip = 700;
					break;
				case 'OPTIMUM':
					this.start_value = 100;
					this.randomness = 0;
					this.transferred_grip = 100;
					this.laps_to_improve_grip = 1;
					break;
			}
		},
		generateCfgDownloadUrl: function (id) {
			return '/api/configuration?id=' + id + '&dl=1';
		},
		calculateSunAngleByTime: function(time) {
			var totalHours = parseInt(time.replace(':', ''), 10);

			if (isNaN(totalHours) || totalHours < 800 || totalHours > 1800) {
				// Invalid date, default to 0 (13:00)
				this.sun_angle = 0;
				return;
			}

			var timeParts = time.split(':');
			var hours = parseInt(timeParts[0], 10);
			var minutes = parseInt(timeParts[1], 10);

			// Calculate the time in minutes and subtract 13 hours
			// The sun angle can range from -80 to 80 in the time frame 08:00h - 18:00h
			// e.g. 08:00 = 480min. - 780min. = -300 / 30 = -10 * 8 = -80
			// e.g. 13:00 = 780min. - 780min. = 0 / 30 = 0 * 8 = 0
			// e.g. 16:30 = 990min. - 780min. = 210 / 30 = 7 * 8 = 56
			var totalMinutes = (hours * 60 + minutes) - 780;

			this.sun_angle = Math.floor(totalMinutes / 30) * 8;
		},
		calculateTimeBySunAngle: function(sun_angle) {
			var totalHours = ((sun_angle / 8) * 30 + 780) / 60;
			var hrs = totalHours.toString();

			return Math.floor(totalHours).toString() + ':' + (hrs[hrs.length - 1] === '5' ? '30' : '00');
		},
		closeMsg() {
			this.err = 0;
		}
	}
};
</script>

<style lang="scss">
.no-border{
	box-shadow: none;
}

.collapse:hover{
	cursor: pointer;
}
</style>
