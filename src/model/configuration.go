package model

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type Configuration struct {
	Id                    int64  `json:"id"`
	Name                  string `json:"name"`
	Pwd                   string `json:"pwd"`
	AdminPwd              string `db:"admin_pwd" json:"admin_pwd"`
	PickupMode            bool   `db:"pickup_mode" json:"pickup_mode"`
	LockEntryList         bool   `db:"lock_entry_list" json:"lock_entry_list"`
	RaceOvertime          int    `db:"race_overtime" json:"race_overtime"`
	MaxSlots              int    `db:"max_slots" json:"max_slots"`
	Welcome               string `json:"welcome"`
	Description           string `json:"description"`
	UDP                   int    `json:"udp"`
	TCP                   int    `json:"tcp"`
	HTTP                  int    `json:"http"`
	PacketsHz             int    `db:"packets_hz" json:"packets_hz"`
	LoopMode              bool   `db:"loop_mode" json:"loop_mode"`
	ShowInLobby           bool   `db:"show_in_lobby" json:"show_in_lobby"`
	Threads               int    `json:"threads"`
	ABS                   int    `json:"abs"`
	TC                    int    `json:"tc"`
	StabilityAid          bool   `db:"stability_aid" json:"stability_aid"`
	AutoClutch            bool   `db:"auto_clutch" json:"auto_clutch"`
	TyreBlankets          bool   `db:"tyre_blankets" json:"tyre_blankets"`
	ForceVirtualMirror    bool   `db:"force_virtual_mirror" json:"force_virtual_mirror"`
	FuelRate              int    `db:"fuel_rate" json:"fuel_rate"`
	DamageRate            int    `db:"damage_rate" json:"damage_rate"`
	TiresWearRate         int    `db:"tires_wear_rate" json:"tires_wear_rate"`
	AllowedTiresOut       int    `db:"allowed_tires_out" json:"allowed_tires_out"`
	MaxBallast            int    `db:"max_ballast" json:"max_ballast"`
	StartRule             int    `db:"start_rule" json:"start_rule"`
	DisableGasCutPenality bool   `db:"disable_gas_cut_penality" json:"disable_gas_cut_penality"`
	TimeOfDayMult         int    `db:"time_of_day_mult" json:"time_of_day_mult"`
	ResultScreenTime      int    `db:"result_screen_time" json:"result_screen_time"`
	DynamicTrack          bool   `db:"dynamic_track" json:"dynamic_track"`
	Condition             string `db:"track_condition" json:"condition"`
	StartValue            int    `db:"start_value" json:"start_value"`
	Randomness            int    `json:"randomness"`
	TransferredGrip       int    `db:"transferred_grip" json:"transferred_grip"`
	LapsToImproveGrip     int    `db:"laps_to_improve_grip" json:"laps_to_improve_grip"`
	KickVoteQuorum        int    `db:"kick_vote_quorum" json:"kick_vote_quorum"`
	SessionVoteQuorum     int    `db:"session_vote_quorum" json:"session_vote_quorum"`
	VoteDuration          int    `db:"vote_duration" json:"vote_duration"`
	Blacklist             int    `json:"blacklist"`
	MaxCollisionsKm       int    `db:"max_collisions_km" json:"max_collisions_km"`
	Booking               bool   `json:"booking"`
	BookingTime           int    `db:"booking_time" json:"booking_time"`
	Practice              bool   `json:"practice"`
	PracticeTime          int    `db:"practice_time" json:"practice_time"`
	CanJoinPractice       bool   `db:"can_join_practice" json:"can_join_practice"`
	Qualify               bool   `json:"qualify"`
	QualifyTime           int    `db:"qualify_time" json:"qualify_time"`
	CanJoinQualify        bool   `db:"can_join_qualify" json:"can_join_qualify"`
	Race                  bool   `json:"race"`
	RaceLaps              int    `db:"race_laps" json:"race_laps"`
	RaceTime              int    `db:"race_time" json:"race_time"`
	RaceWaitTime          int    `db:"race_wait_time" json:"race_wait_time"`
	RaceExtraLap          bool   `db:"race_extra_lap" json:"race_extra_lap"`
	JoinType              int    `db:"join_type" json:"join_type"`
	SunAngle              int    `db:"sun_angle" json:"sun_angle"`
	Track                 string `json:"track"`
	TrackConfig           string `db:"track_config" json:"track_config"`
	LegalTyres            string `db:"legal_tyres" json:"legal_tyres"`
	UdpPluginPort         int    `db:"udp_plugin_local_port" json:"udp_plugin_local_port"`
	UdpPluginAddr         string `db:"udp_plugin_address" json:"udp_plugin_address"`
	RacePitWindowStart    int    `db:"race_pit_window_start" json:"race_pit_window_start"`
	RacePitWindowEnd      int    `db:"race_pit_window_end" json:"race_pit_window_end"`
	ReversedGridRacePos   int    `db:"reversed_grid_race_positions" json:"reversed_grid_race_positions"`

	Weather []Weather `db:"-" json:"weather"`
	Cars    []Car     `db:"-" json:"cars"`
}

type Weather struct {
	Id                     int64  `json:"id"`
	Configuration          int64  `json:"configuration"`
	Weather                string `json:"weather"`
	BaseAmbientTemp        int    `db:"base_ambient_temp" json:"base_ambient_temp"`
	BaseRoadTemp           int    `db:"base_road_temp" json:"base_road_temp"`
	AmbientVariation       int    `db:"ambient_variation" json:"ambient_variation"`
	RoadVariation          int    `db:"road_variation" json:"road_variation"`
	WindBaseSpeedMin       int    `db:"wind_base_speed_min" json:"wind_base_speed_min"`
	WindBaseSpeedMax       int    `db:"wind_base_speed_max" json:"wind_base_speed_max"`
	WindBaseDirection      int    `db:"wind_base_direction" json:"wind_base_direction"`
	WindVariationDirection int    `db:"wind_variation_direction" json:"wind_variation_direction"`
}

type Car struct {
	Id            int64  `json:"id"`
	Configuration int64  `json:"configuration"`
	Car           string `json:"car"`
	Painting      string `json:"painting"`
	Spectator     bool   `json:"spectator"`
	Driver        string `json:"driver"`
	Team          string `json:"team"`
	GUID          string `json:"guid"`
	Position      int    `json:"position"`
	FixedSetup    string `db:"fixed_setup" json:"fixed_setup"`
}

// Joins weather and cars.
func (m *Configuration) Join() error {
	weather, err := GetWeatherByConfiguration(m.Id)

	if err != nil {
		return err
	}

	cars, err := GetCarsByConfiguration(m.Id)

	if err != nil {
		return err
	}

	m.Weather = weather
	m.Cars = cars

	return nil
}

func (m *Configuration) Save() error {
	tx, err := session.Beginx()

	if err != nil {
		return err
	}

	// get an ID first
	if err := m.saveConfiguration(tx); err != nil {
		return err
	}

	// save weather and cars
	if err := m.saveWeather(tx); err != nil {
		return err
	}

	if err := m.saveCars(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (m *Configuration) saveConfiguration(tx *sqlx.Tx) error {
	if m.Id == 0 {
		res, err := tx.NamedExec(`INSERT INTO configurations (name,
			pwd,
			admin_pwd,
			pickup_mode,
			lock_entry_list,
			race_overtime,
			max_slots,
			welcome,
			description,
			udp,
			tcp,
			http,
			packets_hz,
			loop_mode,
			show_in_lobby,
			threads,
			abs,
			tc,
			stability_aid,
			auto_clutch,
			tyre_blankets,
			force_virtual_mirror,
			fuel_rate,
			damage_rate,
			tires_wear_rate,
			allowed_tires_out,
			max_ballast,
			start_rule,
			disable_gas_cut_penality,
			time_of_day_mult,
			result_screen_time,
			dynamic_track,
			track_condition,
			start_value,
			randomness,
			transferred_grip,
			laps_to_improve_grip,
			kick_vote_quorum,
			session_vote_quorum,
			vote_duration,
			blacklist,
			max_collisions_km,
			booking,
			booking_time,
			practice,
			practice_time,
			can_join_practice,
			qualify,
			qualify_time,
			can_join_qualify,
			race,
			race_laps,
			race_time,
			race_wait_time,
			race_extra_lap,
			join_type,
			sun_angle,
			track,
			track_config,
			legal_tyres,
			udp_plugin_local_port,
			udp_plugin_address,
			race_pit_window_start,
			race_pit_window_end,
			reversed_grid_race_positions
			) VALUES (
			:name,
			:pwd,
			:admin_pwd,
			:pickup_mode,
			:lock_entry_list,
			:race_overtime,
			:max_slots,
			:welcome,
			:description,
			:udp,
			:tcp,
			:http,
			:packets_hz,
			:loop_mode,
			:show_in_lobby,
			:threads,
			:abs,
			:tc,
			:stability_aid,
			:auto_clutch,
			:tyre_blankets,
			:force_virtual_mirror,
			:fuel_rate,
			:damage_rate,
			:tires_wear_rate,
			:allowed_tires_out,
			:max_ballast,
			:start_rule,
			:disable_gas_cut_penality,
			:time_of_day_mult,
			:result_screen_time,
			:dynamic_track,
			:track_condition,
			:start_value,
			:randomness,
			:transferred_grip,
			:laps_to_improve_grip,
			:kick_vote_quorum,
			:session_vote_quorum,
			:vote_duration,
			:blacklist,
			:max_collisions_km,
			:booking,
			:booking_time,
			:practice,
			:practice_time,
			:can_join_practice,
			:qualify,
			:qualify_time,
			:can_join_qualify,
			:race,
			:race_laps,
			:race_time,
			:race_wait_time,
			:race_extra_lap,
			:join_type,
			:sun_angle,
			:track,
			:track_config,
			:legal_tyres,
			:udp_plugin_local_port,
			:udp_plugin_address,
			:race_pit_window_start,
			:race_pit_window_end,
			:reversed_grid_race_positions)`, m)

		if err != nil {
			tx.Rollback()
			return err
		}

		id, err := res.LastInsertId()

		if err != nil {
			tx.Rollback()
			return err
		}

		m.Id = id
		return nil
	}

	_, err := tx.NamedExec(`UPDATE configurations SET name = :name,
		pwd = :pwd,
		admin_pwd = :admin_pwd,
		pickup_mode = :pickup_mode,
		lock_entry_list = :lock_entry_list,
		race_overtime = :race_overtime,
		max_slots = :max_slots,
		welcome = :welcome,
		description = :description,
		udp = :udp,
		tcp = :tcp,
		http = :http,
		packets_hz = :packets_hz,
		loop_mode = :loop_mode,
		show_in_lobby = :show_in_lobby,
		threads = :threads,
		abs = :abs,
		tc = :tc,
		stability_aid = :stability_aid,
		auto_clutch = :auto_clutch,
		tyre_blankets = :tyre_blankets,
		force_virtual_mirror = :force_virtual_mirror,
		fuel_rate = :fuel_rate,
		damage_rate = :damage_rate,
		tires_wear_rate = :tires_wear_rate,
		allowed_tires_out = :allowed_tires_out,
		max_ballast = :max_ballast,
		start_rule = :start_rule,
		disable_gas_cut_penality = :disable_gas_cut_penality,
		time_of_day_mult = :time_of_day_mult,
		result_screen_time = :result_screen_time,
		dynamic_track = :dynamic_track,
		track_condition = :track_condition,
		start_value = :start_value,
		randomness = :randomness,
		transferred_grip = :transferred_grip,
		laps_to_improve_grip = :laps_to_improve_grip,
		kick_vote_quorum = :kick_vote_quorum,
		session_vote_quorum = :session_vote_quorum,
		vote_duration = :vote_duration,
		blacklist = :blacklist,
		max_collisions_km = :max_collisions_km,
		booking = :booking,
		booking_time = :booking_time,
		practice = :practice,
		practice_time = :practice_time,
		can_join_practice = :can_join_practice,
		qualify = :qualify,
		qualify_time = :qualify_time,
		can_join_qualify = :can_join_qualify,
		race = :race,
		race_laps = :race_laps,
		race_time = :race_time,
		race_wait_time = :race_wait_time,
		race_extra_lap = :race_extra_lap,
		join_type = :join_type,
		sun_angle = :sun_angle,
		track = :track,
		track_config = :track_config,
		legal_tyres = :legal_tyres,
		udp_plugin_local_port = :udp_plugin_local_port,
		udp_plugin_address = :udp_plugin_address,
		race_pit_window_start = :race_pit_window_start,
		race_pit_window_end = :race_pit_window_end,
		reversed_grid_race_positions = :reversed_grid_race_positions
		WHERE id = :id`, m)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (m *Configuration) saveWeather(tx *sqlx.Tx) error {
	for _, weather := range m.Weather {
		if weather.Id == 0 {
			_, err := tx.Exec(`INSERT INTO weather (configuration,
				weather,
				base_ambient_temp,
				base_road_temp,
				ambient_variation,
				road_variation,
				wind_base_speed_min,
				wind_base_speed_max,
				wind_base_direction,
				wind_variation_direction
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				m.Id,
				weather.Weather,
				weather.BaseAmbientTemp,
				weather.BaseRoadTemp,
				weather.AmbientVariation,
				weather.RoadVariation,
				weather.WindBaseSpeedMin,
				weather.WindBaseSpeedMax,
				weather.WindBaseDirection,
				weather.WindVariationDirection)

			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			_, err := tx.Exec(`UPDATE weather SET weather = ?,
				base_ambient_temp = ?,
				base_road_temp = ?,
				ambient_variation = ?,
				road_variation = ?,
				wind_base_speed_min = ?,
				wind_base_speed_max = ?,
				wind_base_direction = ?,
				wind_variation_direction = ?
				WHERE id = ?`,
				weather.Weather,
				weather.BaseAmbientTemp,
				weather.BaseRoadTemp,
				weather.AmbientVariation,
				weather.RoadVariation,
				weather.WindBaseSpeedMin,
				weather.WindBaseSpeedMax,
				weather.WindBaseDirection,
				weather.WindVariationDirection,
				weather.Id)

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (m *Configuration) saveCars(tx *sqlx.Tx) error {
	for _, car := range m.Cars {
		if car.Id == 0 {
			_, err := tx.Exec(`INSERT INTO cars (configuration,
				car,
				painting,
				spectator,
				driver,
				team,
				guid,
				position,
				fixed_setup
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				m.Id,
				car.Car,
				car.Painting,
				car.Spectator,
				car.Driver,
				car.Team,
				car.GUID,
				car.Position,
				car.FixedSetup)

			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			_, err := tx.Exec(`UPDATE cars SET car = ?,
				painting = ?,
				spectator = ?,
				driver = ?,
				team = ?,
				guid = ?,
				position = ?,
				fixed_setup = ? WHERE id = ?`,
				car.Car,
				car.Painting,
				car.Spectator,
				car.Driver,
				car.Team,
				car.GUID,
				car.Position,
				car.FixedSetup,
				car.Id)

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (m *Configuration) Remove() error {
	if m.Id == 0 {
		return errors.New("ID must be set")
	}

	tx, err := session.Beginx()

	if err != nil {
		return err
	}

	_, err = tx.NamedExec("DELETE FROM weather WHERE configuration = :id", m)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NamedExec("DELETE FROM cars WHERE configuration = :id", m)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NamedExec("DELETE FROM configurations WHERE id = :id", m)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (m *Weather) Remove() error {
	_, err := session.NamedExec("DELETE FROM weather WHERE id = :id", m)
	return err
}

func (m *Car) Remove() error {
	_, err := session.NamedExec("DELETE FROM cars WHERE id = :id", m)
	return err
}

func GetWeatherByConfiguration(id int64) ([]Weather, error) {
	weather := make([]Weather, 0)

	if err := session.Select(&weather, "SELECT * FROM weather WHERE configuration = ?", id); err != nil {
		return nil, err
	}

	return weather, nil
}

func GetCarsByConfiguration(id int64) ([]Car, error) {
	cars := make([]Car, 0)

	if err := session.Select(&cars, "SELECT * FROM cars WHERE configuration = ? ORDER BY position ASC", id); err != nil {
		return nil, err
	}

	return cars, nil
}

func GetAllConfigurations() ([]Configuration, error) {
	configs := make([]Configuration, 0)

	if err := session.Select(&configs, "SELECT * FROM configurations ORDER BY name ASC"); err != nil {
		return nil, err
	}

	return configs, nil
}

func GetConfigurationById(id int64) (*Configuration, error) {
	config := new(Configuration)

	if err := session.Get(config, "SELECT * FROM configurations WHERE id = ?", id); err != nil {
		return nil, err
	}

	return config, nil
}
