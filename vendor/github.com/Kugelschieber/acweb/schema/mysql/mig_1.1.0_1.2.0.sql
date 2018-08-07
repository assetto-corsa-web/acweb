ALTER TABLE `configurations`
ADD COLUMN `start_rule` int(10) NOT NULL AFTER `max_ballast`,
ADD COLUMN `time_of_day_mult` int(10) NOT NULL AFTER `disable_gas_cut_penality`;
