CREATE TABLE IF NOT EXISTS `employee` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`address` varchar(255) DEFAULT NULL,
	`city` varchar(255) DEFAULT NULL,
	`country` varchar(255) DEFAULT NULL,
	`date_of_birth` date DEFAULT NULL,
	`department` varchar(255) DEFAULT NULL,
	`email` varchar(255) DEFAULT NULL,
	`first_name` varchar(255) DEFAULT NULL,
	`hire_date` date DEFAULT NULL,
	`last_name` varchar(255) DEFAULT NULL,
	`phone_number` varchar(255) DEFAULT NULL,
	`position` varchar(255) DEFAULT NULL,
	`postal_code` varchar(255) DEFAULT NULL,
	`salary` varchar(255) DEFAULT NULL,
	`state` varchar(255) DEFAULT NULL,
	PRIMARY KEY (`id`)
  );
  