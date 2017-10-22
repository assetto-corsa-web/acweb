ALTER TABLE `configurations`
ADD COLUMN `server_cfg_ini` TEXT NOT NULL AFTER `reversed_grid_race_positions`,
ADD COLUMN `entry_list_ini` TEXT NOT NULL AFTER `server_cfg_ini`;
