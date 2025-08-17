USE `master`;

CREATE TABLE `jenis_wilayah` (
  `ID` tinyint(4) NOT NULL AUTO_INCREMENT,
  `DESKRIPSI` varchar(20) NOT NULL,
  `DIGIT` tinyint(4) NOT NULL DEFAULT '2',
  `DELIMITER` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO master.jenis_wilayah (ID, DESKRIPSI, DIGIT, DELIMITER)
VALUES
(1, 'Propinsi', 2, ''),
(2, 'Kabupaten / Kota', 2, ''),
(3, 'Kecamatan', 2, ''),
(4, 'Kelurahan / Desa', 4, '');