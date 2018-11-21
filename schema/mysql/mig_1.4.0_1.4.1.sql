ALTER TABLE `configurations`
ADD COLUMN `auth_plugin_address` int(10) NOT NULL AFTER `threads`;
