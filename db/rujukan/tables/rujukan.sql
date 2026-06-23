USE `rujukan`;

CREATE TABLE `rujukan` (
  `ID`               int          NOT NULL AUTO_INCREMENT,
  `NO_RUJUKAN`       varchar(30)  NOT NULL,
  `ID_PASIEN`        varchar(36)  NOT NULL,
  `ID_FASKES_ASAL`   char(10)     NOT NULL,
  `ID_FASKES_TUJUAN` char(10)     NOT NULL,
  `KODE_ICD`         varchar(10)  DEFAULT NULL,
  `DIAGNOSA`         text         DEFAULT NULL,
  `CATATAN`          text         DEFAULT NULL,
  `STATUS`           enum('menunggu','diterima','ditolak','selesai') NOT NULL DEFAULT 'menunggu',
  `ID_USER`          smallint     NOT NULL,
  `TANGGAL`          datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `NO_RUJUKAN` (`NO_RUJUKAN`),
  KEY `ID_PASIEN` (`ID_PASIEN`),
  KEY `ID_FASKES_ASAL` (`ID_FASKES_ASAL`),
  KEY `STATUS` (`STATUS`),
  KEY `TANGGAL` (`TANGGAL`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
