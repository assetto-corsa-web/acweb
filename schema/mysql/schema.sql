SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

CREATE TABLE `cars` (
  `id` int(10) UNSIGNED NOT NULL,
  `configuration` int(10) UNSIGNED NOT NULL,
  `car` varchar(100) NOT NULL,
  `painting` varchar(100) NOT NULL,
  `spectator` tinyint(1) NOT NULL,
  `driver` varchar(40) NOT NULL,
  `team` varchar(40) NOT NULL,
  `guid` varchar(100) NOT NULL,
  `position` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `configurations` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(200) NOT NULL,
  `pwd` varchar(40) NOT NULL,
  `admin_pwd` varchar(40) NOT NULL,
  `pickup_mode` tinyint(1) NOT NULL,
  `lock_entry_list` tinyint(1) NOT NULL,
  `race_overtime` int(11) NOT NULL,
  `max_slots` int(11) NOT NULL,
  `welcome` varchar(200) NOT NULL,
  `description` text NOT NULL,
  `udp` int(11) NOT NULL,
  `tcp` int(11) NOT NULL,
  `http` int(11) NOT NULL,
  `packets_hz` int(11) NOT NULL,
  `loop_mode` tinyint(1) NOT NULL,
  `show_in_lobby` tinyint(1) NOT NULL,
  `threads` int(11) NOT NULL,
  `abs` varchar(40) NOT NULL,
  `tc` varchar(40) NOT NULL,
  `stability_aid` tinyint(1) NOT NULL,
  `auto_clutch` tinyint(1) NOT NULL,
  `tyre_blankets` tinyint(1) NOT NULL,
  `force_virtual_mirror` tinyint(1) NOT NULL,
  `fuel_rate` int(11) NOT NULL,
  `damage_rate` int(11) NOT NULL,
  `tires_wear_rate` int(11) NOT NULL,
  `allowed_tires_out` int(11) NOT NULL,
  `max_ballast` int(11) NOT NULL,
  `disable_gas_cut_penality` tinyint(1) NOT NULL,
  `result_screen_time` int(11) NOT NULL,
  `dynamic_track` tinyint(1) NOT NULL,
  `track_condition` varchar(40) NOT NULL,
  `start_value` int(11) NOT NULL,
  `randomness` int(11) NOT NULL,
  `transferred_grip` int(11) NOT NULL,
  `laps_to_improve_grip` int(11) NOT NULL,
  `kick_vote_quorum` int(11) NOT NULL,
  `session_vote_quorum` int(11) NOT NULL,
  `vote_duration` int(11) NOT NULL,
  `blacklist` varchar(40) NOT NULL,
  `max_collisions_km` int(11) NOT NULL,
  `booking` tinyint(1) NOT NULL,
  `booking_time` int(11) NOT NULL,
  `practice` tinyint(1) NOT NULL,
  `practice_time` int(11) NOT NULL,
  `can_join_practice` tinyint(1) NOT NULL,
  `qualify` tinyint(1) NOT NULL,
  `qualify_time` int(11) NOT NULL,
  `can_join_qualify` tinyint(1) NOT NULL,
  `race` tinyint(1) NOT NULL,
  `race_laps` int(11) NOT NULL,
  `race_time` int(11) NOT NULL,
  `race_wait_time` int(11) NOT NULL,
  `race_extra_lap` tinyint(1) NOT NULL,
  `join_type` varchar(40) NOT NULL,
  `time` varchar(20) NOT NULL,
  `sun_angle` int(11) NOT NULL,
  `track` varchar(100) NOT NULL,
  `track_config` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `settings` (
  `id` int(10) UNSIGNED NOT NULL,
  `folder` text NOT NULL,
  `executable` text NOT NULL,
  `args` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user` (
  `id` int(10) UNSIGNED NOT NULL,
  `login` varchar(40) NOT NULL,
  `email` varchar(200) NOT NULL,
  `password` varchar(64) NOT NULL,
  `admin` tinyint(1) NOT NULL,
  `moderator` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `weather` (
  `id` int(10) UNSIGNED NOT NULL,
  `configuration` int(10) UNSIGNED NOT NULL,
  `weather` varchar(40) NOT NULL,
  `base_ambient_temp` int(11) NOT NULL,
  `realistic_road_temp` int(11) NOT NULL,
  `base_road_temp` int(11) NOT NULL,
  `ambient_variation` int(11) NOT NULL,
  `road_variation` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


ALTER TABLE `cars`
  ADD PRIMARY KEY (`id`),
  ADD KEY `cars_config_fk` (`configuration`);

ALTER TABLE `configurations`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `settings`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `weather`
  ADD PRIMARY KEY (`id`),
  ADD KEY `weather_config_fk` (`configuration`);


ALTER TABLE `cars`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=42;
ALTER TABLE `configurations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
ALTER TABLE `settings`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
ALTER TABLE `user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
ALTER TABLE `weather`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

ALTER TABLE `cars`
  ADD CONSTRAINT `cars_config_fk` FOREIGN KEY (`configuration`) REFERENCES `configurations` (`id`);

ALTER TABLE `weather`
  ADD CONSTRAINT `weather_config_fk` FOREIGN KEY (`configuration`) REFERENCES `configurations` (`id`);

