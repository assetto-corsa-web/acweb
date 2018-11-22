ALTER TABLE `configurations`
ADD COLUMN `auth_plugin_address` int(10) NOT NULL AFTER `threads`;

ALTER TABLE `cars`
ADD COLUMN `ballast` int(10) NOT NULL AFTER `fixed_setup`,
ADD COLUMN `restrictor` int(10) NOT NULL AFTER `ballast`;
