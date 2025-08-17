-- aplikasi.pengguna definition
USE `aplikasi`;

CREATE TABLE `user` (
  `ID` smallint NOT NULL AUTO_INCREMENT,
  `USERNAME` varchar(50) NOT NULL,
  `PASSWORD` text NOT NULL,
  `NAMA` varchar(150) NOT NULL,
  `NIK` char(16) DEFAULT NULL,
  `ID_FASKES` smallint DEFAULT '0',
  `MASA_AKTIF` date DEFAULT NULL,
  `TOKEN` text NOT NULL DEFAULT '',
  `ROLE` enum('admin', 'faskes', 'dinkeskab', 'dinkesprov') COMMENT 'penentuan alur hak akses berdasarkan role',
  `STATUS_TELEGRAM` tinyint DEFAULT '0' COMMENT '1. telgram aktif 0. tidak aktif',
  `ID_TELEGRAM` char(12) DEFAULT '' COMMENT 'menyimpan id telegram dari user',
  `TERAKHIR_UBAH_PASSWORD` datetime DEFAULT NULL,
  `STATUS` tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`ID`),
  UNIQUE KEY `USERNAME` (`USERNAME`),
  KEY `NIK` (`NIK`),
  KEY `ID_FASKES` (`ID_FASKES`),
  KEY `STATUS` (`STATUS`),
  KEY `MASA_AKTIF` (`MASA_AKTIF`),
  KEY `STATUS_TELEGRAM` (`STATUS_TELEGRAM`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;