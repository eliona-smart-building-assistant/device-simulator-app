SET SCHEMA 'device_simulator';

ALTER TABLE device_simulator.generator
ALTER interval_seconds TYPE double precision;
