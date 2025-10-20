SET SCHEMA 'device_simulator';

ALTER TABLE device_simulator.generator
    ADD IF NOT EXISTS tenant_id uuid NOT NULL DEFAULT '5a01cd84-b514-4101-ac5f-2a6fe55b6b99'; -- compatibility reasons
