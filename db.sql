CREATE DATABASE `mini_exercise` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `mini_exercise`;

-- mini_exercise.deposits definition

CREATE TABLE `deposits` (
  `id` varchar(36) NOT NULL,
  `deposited_by` varchar(36) NOT NULL,
  `status` varchar(10) NOT NULL,
  `deposited_at` datetime NOT NULL,
  `amount` decimal(10,0) NOT NULL,
  `reference_id` varchar(36) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `reference_id_UN` (`reference_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- mini_exercise.wallets definition

CREATE TABLE `wallets` (
  `id` varchar(36) NOT NULL,
  `owned_by` varchar(36) NOT NULL,
  `status` varchar(50) NOT NULL,
  `enabled_at` datetime DEFAULT NULL,
  `balance` decimal(10,0) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- mini_exercise.withdrawals definition

CREATE TABLE `withdrawals` (
  `id` varchar(36) NOT NULL,
  `withdrawn_by` varchar(36) NOT NULL,
  `status` varchar(50) NOT NULL,
  `withdrawn_at` datetime NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `reference_id` varchar(36) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;