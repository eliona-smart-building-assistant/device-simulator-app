--  This file is part of the Eliona project.
--  Copyright © 2024 IoTEC AG. All Rights Reserved.
--  ______ _ _
-- |  ____| (_)
-- | |__  | |_  ___  _ __   __ _
-- |  __| | | |/ _ \| '_ \ / _` |
-- | |____| | | (_) | | | | (_| |
-- |______|_|_|\___/|_| |_|\__,_|
--
--  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
--  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
--  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
--  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
--  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

create schema if not exists device_simulator;

-- Should be editable by eliona frontend.
CREATE TABLE IF NOT EXISTS device_simulator.generator
(
    id                  BIGSERIAL PRIMARY KEY,
    asset_id            INTEGER NOT NULL,
    attribute           TEXT NOT NULL,
    subtype             TEXT NOT NULL,
    asset_type          TEXT NOT NULL,
    function_type       VARCHAR(50) NOT NULL,      -- Type of function (e.g., "boolean", "sin_wave", "sawtooth_wave")
    min_value           DOUBLE PRECISION NOT NULL, -- Minimum value for the generated data
    max_value           DOUBLE PRECISION NOT NULL, -- Maximum value for the generated data
    interval_seconds    INTEGER NOT NULL,          -- Interval in seconds for data generation
    frequency           DOUBLE PRECISION NOT NULL  -- Frequency for wave functions or duty cycle
);

-- There is a transaction started in app.Init(). We need to commit to make the
-- new objects available for all other init steps.
-- Chain starts the same transaction again.
commit and chain;
