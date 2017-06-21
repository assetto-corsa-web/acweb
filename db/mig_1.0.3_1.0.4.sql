
ALTER TABLE `cars`
ADD COLUMN `fixed_setup` varchar(100) NOT NULL AFTER `position`;

ALTER TABLE `configurations`
ADD COLUMN `legal_tyres` varchar(100) NOT NULL AFTER `track_config`
ADD COLUMN `udp_plugin_local_port` int(10) NOT NULL AFTER `legal_tyres`,
ADD COLUMN `udp_plugin_address` varchar(100) NOT NULL AFTER `udp_plugin_address`;

ALTER TABLE `weather`
ADD COLUMN `wind_base_speed_min` int(10) NOT NULL AFTER `road_variation`,
ADD COLUMN `wind_base_speed_max` int(10) NOT NULL AFTER `wind_base_speed_min`,
ADD COLUMN `wind_base_direction` int(10) NOT NULL AFTER `wind_base_speed_max`,
ADD COLUMN `wind_variation_direction` int(10) NOT NULL AFTER `wind_base_direction`;
