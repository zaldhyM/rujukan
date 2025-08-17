-- sisrute.pasien definition
USE `master`;

CREATE TABLE `pasien` (
  `ID` CHAR(36) DEFAULT '' COMMENT 'UID',
  `ID_FAKSES` smallint DEFAULT '0' COMMENT 'db master table faskes column ID',
  `NORM` char(20) DEFAULT '' COMMENT 'norm pasien',
  `NAMA` varchar(150) DEFAULT '',
  `KONTAK` varchar(30) DEFAULT '',
  `ALAMAT` text DEFAULT '',
  `TEMPAT_LAHIR` text DEFAULT '',
  `TANGGAL_LAHIR` DATE DEFAULT NULL,
  `JENIS_KELAMIN` smallint DEFAULT '0' COMMENT 'db master table referensi column ID JENIS = 2',
  `NO_JKN` varchar(20) DEFAULT '',
  `NIK` varchar(20) DEFAULT '',
  `STATUS` tinyint(4) DEFAULT '1' COMMENT '1. AKTIF, 0. TIDAK AKTIF',
  `TANGGAL` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  KEY `ID_FAKSES` (`ID_FAKSES`),
  KEY `NORM` (`NORM`),
  KEY `NAMA` (`NAMA`),
  KEY `STATUS` (`STATUS`),
  KEY `NIK` (`NIK`),
  KEY `NO_JKN` (`NO_JKN`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;