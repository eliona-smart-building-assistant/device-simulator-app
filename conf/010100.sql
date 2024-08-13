SET SCHEMA 'device_simulator';

ALTER TABLE "generator"
ALTER "interval_seconds" TYPE double precision;
