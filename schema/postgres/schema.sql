CREATE SEQUENCE cars_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE cars (
  id bigint NOT NULL,
  configuration integer NOT NULL,
  car character varying(100) NOT NULL,
  painting character varying(100) NOT NULL,
  spectator boolean NOT NULL,
  driver character varying(40) NOT NULL,
  team character varying(40) NOT NULL,
  guid character varying(100) NOT NULL,
  position integer NOT NULL,
  fixed_setup character varying(100) NOT NULL
);

ALTER SEQUENCE cars_id_seq OWNED BY cars.id;

CREATE SEQUENCE configurations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE configurations (
  id bigint NOT NULL,
  name character varying(200) NOT NULL,
  pwd character varying(40) NOT NULL,
  admin_pwd character varying(40) NOT NULL,
  pickup_mode boolean NOT NULL,
  lock_entry_list boolean NOT NULL,
  race_overtime integer NOT NULL,
  max_slots integer NOT NULL,
  welcome character varying(200) NOT NULL,
  description text NOT NULL,
  udp integer NOT NULL,
  tcp integer NOT NULL,
  http integer NOT NULL,
  packets_hz integer NOT NULL,
  loop_mode boolean NOT NULL,
  show_in_lobby boolean NOT NULL,
  threads integer NOT NULL,
  abs character varying(40) NOT NULL,
  tc character varying(40) NOT NULL,
  stability_aid boolean NOT NULL,
  auto_clutch boolean NOT NULL,
  tyre_blankets boolean NOT NULL,
  force_virtual_mirror boolean NOT NULL,
  fuel_rate integer NOT NULL,
  damage_rate integer NOT NULL,
  tires_wear_rate integer NOT NULL,
  allowed_tires_out integer NOT NULL,
  max_ballast integer NOT NULL,
  start_rule integer NOT NULL,
  disable_gas_cut_penality boolean NOT NULL,
  time_of_day_mult integer NOT NULL,
  result_screen_time integer NOT NULL,
  dynamic_track boolean NOT NULL,
  track_condition character varying(40) NOT NULL,
  start_value integer NOT NULL,
  randomness integer NOT NULL,
  transferred_grip integer NOT NULL,
  laps_to_improve_grip integer NOT NULL,
  kick_vote_quorum integer NOT NULL,
  session_vote_quorum integer NOT NULL,
  vote_duration integer NOT NULL,
  blacklist character varying(40) NOT NULL,
  max_collisions_km integer NOT NULL,
  booking boolean NOT NULL,
  booking_time integer NOT NULL,
  practice boolean NOT NULL,
  practice_time integer NOT NULL,
  can_join_practice boolean NOT NULL,
  qualify boolean NOT NULL,
  qualify_time integer NOT NULL,
  can_join_qualify boolean NOT NULL,
  race boolean NOT NULL,
  race_laps integer NOT NULL,
  race_time integer NOT NULL,
  race_wait_time integer NOT NULL,
  race_extra_lap boolean NOT NULL,
  join_type character varying(40) NOT NULL,
  sun_angle integer NOT NULL,
  track character varying(100) NOT NULL,
  track_config character varying(100) NOT NULL,
  legal_tyres character varying(100) NOT NULL,
  udp_plugin_local_port integer NOT NULL,
  udp_plugin_address character varying(100) NOT NULL,
  race_pit_window_start integer NOT NULL,
  race_pit_window_end integer NOT NULL,
  reversed_grid_race_positions integer NOT NULL,
  server_cfg_ini text NOT NULL,
  entry_list_ini text NOT NULL
);

ALTER SEQUENCE configurations_id_seq OWNED BY configurations.id;

CREATE SEQUENCE settings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE settings (
  id bigint NOT NULL,
  folder text NOT NULL,
  executable text NOT NULL,
  args character varying(500) NOT NULL
);

ALTER SEQUENCE settings_id_seq OWNED BY settings.id;

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE "user" (
  id bigint NOT NULL,
  login character varying(40) NOT NULL,
  email character varying(200) NOT NULL,
  password character varying(64) NOT NULL,
  admin boolean NOT NULL,
  moderator boolean NOT NULL
);

ALTER SEQUENCE user_id_seq OWNED BY "user".id;

CREATE SEQUENCE weather_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE weather (
  id bigint NOT NULL,
  configuration integer NOT NULL,
  weather character varying(40) NOT NULL,
  base_ambient_temp integer NOT NULL,
  base_road_temp integer NOT NULL,
  ambient_variation integer NOT NULL,
  road_variation integer NOT NULL,
  wind_base_speed_min integer NOT NULL,
  wind_base_speed_max integer NOT NULL,
  wind_base_direction integer NOT NULL,
  wind_variation_direction integer NOT NULL
);

ALTER SEQUENCE weather_id_seq OWNED BY weather.id;

ALTER TABLE ONLY cars ALTER COLUMN id SET DEFAULT nextval('cars_id_seq'::regclass);
ALTER TABLE ONLY configurations ALTER COLUMN id SET DEFAULT nextval('configurations_id_seq'::regclass);
ALTER TABLE ONLY settings ALTER COLUMN id SET DEFAULT nextval('settings_id_seq'::regclass);
ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);
ALTER TABLE ONLY weather ALTER COLUMN id SET DEFAULT nextval('weather_id_seq'::regclass);

ALTER TABLE ONLY cars
    ADD CONSTRAINT cars_pkey PRIMARY KEY (id);
ALTER TABLE ONLY configurations
    ADD CONSTRAINT configurations_pkey PRIMARY KEY (id);
ALTER TABLE ONLY settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);
ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);
ALTER TABLE ONLY weather
    ADD CONSTRAINT weather_pkey PRIMARY KEY (id);

ALTER TABLE cars
  ADD CONSTRAINT cars_config_fk FOREIGN KEY (configuration) REFERENCES configurations (id);

ALTER TABLE weather
  ADD CONSTRAINT weather_config_fk FOREIGN KEY (configuration) REFERENCES configurations (id);
