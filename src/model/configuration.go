package model

import (
	"database/sql"
	"db"
	"errors"
	"github.com/DeKugelschieber/go-util"
)

type Configuration struct {
	Id                    int64     `json:"id"`
	Name                  string    `json:"name"`
	Pwd                   string    `json:"pwd"`
	AdminPwd              string    `json:"admin_pwd"`
	PickupMode            bool      `json:"pickup_mode"`
	LockEntryList         bool      `json:"lock_entry_list"`
	RacePitWindowStart    int       `json:"race_pit_window_start"`
	RacePitWindowEnd      int       `json:"race_pit_window_end"`
	ReversedGridRacePos   int       `json:"reversed_grid_race_positions"`
	RaceOvertime          int       `json:"race_overtime"`
	MaxSlots              int       `json:"max_slots"`
	Welcome               string    `json:"welcome"`
	Description           string    `json:"description"`
	UDP                   int       `json:"udp"`
	TCP                   int       `json:"tcp"`
	HTTP                  int       `json:"http"`
	PacketsHz             int       `json:"packets_hz"`
	LoopMode              bool      `json:"loop_mode"`
	ShowInLobby           bool      `json:"show_in_lobby"`
	Threads               int       `json:"threads"`
	ABS                   int       `json:"abs"`
	TC                    int       `json:"tc"`
	StabilityAid          bool      `json:"stability_aid"`
	AutoClutch            bool      `json:"auto_clutch"`
	TyreBlankets          bool      `json:"tyre_blankets"`
	ForceVirtualMirror    bool      `json:"force_virtual_mirror"`
	FuelRate              int       `json:"fuel_rate"`
	DamageRate            int       `json:"damage_rate"`
	TiresWearRate         int       `json:"tires_wear_rate"`
	AllowedTiresOut       int       `json:"allowed_tires_out"`
	MaxBallast            int       `json:"max_ballast"`
	DisableGasCutPenality bool      `json:"disable_gas_cut_penality"`
	ResultScreenTime      int       `json:"result_screen_time"`
	DynamicTrack          bool      `json:"dynamic_track"`
	Condition             string    `json:"condition"`
	StartValue            int       `json:"start_value"`
	Randomness            int       `json:"randomness"`
	TransferredGrip       int       `json:"transferred_grip"`
	LapsToImproveGrip     int       `json:"laps_to_improve_grip"`
	KickVoteQuorum        int       `json:"kick_vote_quorum"`
	SessionVoteQuorum     int       `json:"session_vote_quorum"`
	VoteDuration          int       `json:"vote_duration"`
	Blacklist             int       `json:"blacklist"`
	MaxCollisionsKm       int       `json:"max_collisions_km"`
	Booking               bool      `json:"booking"`
	BookingTime           int       `json:"booking_time"`
	Practice              bool      `json:"practice"`
	PracticeTime          int       `json:"practice_time"`
	CanJoinPractice       bool      `json:"can_join_practice"`
	Qualify               bool      `json:"qualify"`
	QualifyTime           int       `json:"qualify_time"`
	CanJoinQualify        bool      `json:"can_join_qualify"`
	Race                  bool      `json:"race"`
	RaceLaps              int       `json:"race_laps"`
	RaceTime              int       `json:"race_time"`
	RaceWaitTime          int       `json:"race_wait_time"`
	RaceExtraLap          bool      `json:"race_extra_lap"`
	JoinType              int       `json:"join_type"`
	Time                  string    `json:"time"`
	SunAngle              int       `json:"sun_angle"`
	Track                 string    `json:"track"`
	TrackConfig           string    `json:"track_config"`
	LegalTyres            string    `json:"legal_tyres"`
	UdpPluginPort         int       `json:"udp_plugin_local_port"`
	UdpPluginAddr         string    `json:"udp_plugin_address"`
	Weather               []Weather `json:"weather"`
	Cars                  []Car     `json:"cars"`
}

type Weather struct {
	Id                     int64  `json:"id"`
	Configuration          int64  `json:"configuration"`
	Weather                string `json:"weather"`
	BaseAmbientTemp        int    `json:"base_ambient_temp"`
	RealisticRoadTemp      int    `json:"realistic_road_temp"`
	BaseRoadTemp           int    `json:"base_road_temp"`
	AmbientVariation       int    `json:"ambient_variation"`
	RoadVariation          int    `json:"road_variation"`
	WindBaseSpeedMin       int    `json:"wind_base_speed_min"`
	WindBaseSpeedMax       int    `json:"wind_base_speed_max"`
	WindBaseDirection      int    `json:"wind_base_direction"`
	WindVariationDirection int    `json:"wind_variation_direction"`
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
	FixedSetup    string `json:"fixed_setup"`
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
	tx, err := db.Get().Begin()

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

func (m *Configuration) saveConfiguration(tx *sql.Tx) error {
	if m.Id == 0 {
		res, err := tx.Exec(`INSERT INTO configurations (name,
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
			disable_gas_cut_penality,
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
			time,
			sun_angle,
			track,
			track_config,
			legal_tyres,
			udp_plugin_local_port,
			udp_plugin_address,
			race_pit_window_start,
			race_pit_window_end,
			reversed_grid_race_positions) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?)`, m.Name,
			m.Pwd,
			m.AdminPwd,
			m.PickupMode,
			m.LockEntryList,
			m.RaceOvertime,
			m.MaxSlots,
			m.Welcome,
			m.Description,
			m.UDP,
			m.TCP,
			m.HTTP,
			m.PacketsHz,
			m.LoopMode,
			m.ShowInLobby,
			m.Threads,
			m.ABS,
			m.TC,
			m.StabilityAid,
			m.AutoClutch,
			m.TyreBlankets,
			m.ForceVirtualMirror,
			m.FuelRate,
			m.DamageRate,
			m.TiresWearRate,
			m.AllowedTiresOut,
			m.MaxBallast,
			m.DisableGasCutPenality,
			m.ResultScreenTime,
			m.DynamicTrack,
			m.Condition,
			m.StartValue,
			m.Randomness,
			m.TransferredGrip,
			m.LapsToImproveGrip,
			m.KickVoteQuorum,
			m.SessionVoteQuorum,
			m.VoteDuration,
			m.Blacklist,
			m.MaxCollisionsKm,
			m.Booking,
			m.BookingTime,
			m.Practice,
			m.PracticeTime,
			m.CanJoinPractice,
			m.Qualify,
			m.QualifyTime,
			m.CanJoinQualify,
			m.Race,
			m.RaceLaps,
			m.RaceTime,
			m.RaceWaitTime,
			m.RaceExtraLap,
			m.JoinType,
			m.Time,
			m.SunAngle,
			m.Track,
			m.TrackConfig,
			m.LegalTyres,
			m.UdpPluginPort,
			m.UdpPluginAddr,
			m.RacePitWindowStart,
			m.RacePitWindowEnd,
			m.ReversedGridRacePos)

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

	_, err := tx.Exec(`UPDATE configurations SET name = ?,
		pwd = ?,
		admin_pwd = ?,
		pickup_mode = ?,
		lock_entry_list = ?,
		race_overtime = ?,
		max_slots = ?,
		welcome = ?,
		description = ?,
		udp = ?,
		tcp = ?,
		http = ?,
		packets_hz = ?,
		loop_mode = ?,
		show_in_lobby = ?,
		threads = ?,
		abs = ?,
		tc = ?,
		stability_aid = ?,
		auto_clutch = ?,
		tyre_blankets = ?,
		force_virtual_mirror = ?,
		fuel_rate = ?,
		damage_rate = ?,
		tires_wear_rate = ?,
		allowed_tires_out = ?,
		max_ballast = ?,
		disable_gas_cut_penality = ?,
		result_screen_time = ?,
		dynamic_track = ?,
		track_condition = ?,
		start_value = ?,
		randomness = ?,
		transferred_grip = ?,
		laps_to_improve_grip = ?,
		kick_vote_quorum = ?,
		session_vote_quorum = ?,
		vote_duration = ?,
		blacklist = ?,
		max_collisions_km = ?,
		booking = ?,
		booking_time = ?,
		practice = ?,
		practice_time = ?,
		can_join_practice = ?,
		qualify = ?,
		qualify_time = ?,
		can_join_qualify = ?,
		race = ?,
		race_laps = ?,
		race_time = ?,
		race_wait_time = ?,
		race_extra_lap = ?,
		join_type = ?,
		time = ?,
		sun_angle = ?,
		track = ?,
		track_config = ?,
		legal_tyres = ?,
		udp_plugin_local_port = ?,
		udp_plugin_address = ?,
		race_pit_window_start = ?,
		race_pit_window_end = ?,
		reversed_grid_race_positions = ?
		WHERE id = ?`, m.Name,
		m.Pwd,
		m.AdminPwd,
		m.PickupMode,
		m.LockEntryList,
		m.RaceOvertime,
		m.MaxSlots,
		m.Welcome,
		m.Description,
		m.UDP,
		m.TCP,
		m.HTTP,
		m.PacketsHz,
		m.LoopMode,
		m.ShowInLobby,
		m.Threads,
		m.ABS,
		m.TC,
		m.StabilityAid,
		m.AutoClutch,
		m.TyreBlankets,
		m.ForceVirtualMirror,
		m.FuelRate,
		m.DamageRate,
		m.TiresWearRate,
		m.AllowedTiresOut,
		m.MaxBallast,
		m.DisableGasCutPenality,
		m.ResultScreenTime,
		m.DynamicTrack,
		m.Condition,
		m.StartValue,
		m.Randomness,
		m.TransferredGrip,
		m.LapsToImproveGrip,
		m.KickVoteQuorum,
		m.SessionVoteQuorum,
		m.VoteDuration,
		m.Blacklist,
		m.MaxCollisionsKm,
		m.Booking,
		m.BookingTime,
		m.Practice,
		m.PracticeTime,
		m.CanJoinPractice,
		m.Qualify,
		m.QualifyTime,
		m.CanJoinQualify,
		m.Race,
		m.RaceLaps,
		m.RaceTime,
		m.RaceWaitTime,
		m.RaceExtraLap,
		m.JoinType,
		m.Time,
		m.SunAngle,
		m.Track,
		m.TrackConfig,
		m.LegalTyres,
		m.UdpPluginPort,
		m.UdpPluginAddr,
		m.RacePitWindowStart,
		m.RacePitWindowEnd,
		m.ReversedGridRacePos,
		m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (m *Configuration) saveWeather(tx *sql.Tx) error {
	for _, weather := range m.Weather {
		if weather.Id == 0 {
			_, err := tx.Exec(`INSERT INTO weather (configuration,
				weather,
				base_ambient_temp,
				realistic_road_temp,
				base_road_temp,
				ambient_variation,
				road_variation,
				wind_base_speed_min,
				wind_base_speed_max,
				wind_base_direction,
				wind_variation_direction
				) VALUES (?, ?, ?, ?, ?, ?, ?)`, m.Id,
				weather.Weather,
				weather.BaseAmbientTemp,
				weather.RealisticRoadTemp,
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
				realistic_road_temp = ?,
				base_road_temp = ?,
				ambient_variation = ?,
				road_variation = ?,
				wind_base_speed_min = ?,
				wind_base_speed_max = ?,
				wind_base_direction = ?,
				wind_variation_direction = ?
				WHERE id = ?`, weather.Weather,
				weather.BaseAmbientTemp,
				weather.RealisticRoadTemp,
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

func (m *Configuration) saveCars(tx *sql.Tx) error {
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
				fixed_setup) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, m.Id,
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
				fixed_setup = ? WHERE id = ?`, car.Car,
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

	tx, err := db.Get().Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM weather WHERE configuration = ?", m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM cars WHERE configuration = ?", m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM configurations WHERE id = ?", m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (m *Weather) Remove() error {
	_, err := db.Get().Exec("DELETE FROM weather WHERE id = ?", m.Id)
	return err
}

func (m *Car) Remove() error {
	_, err := db.Get().Exec("DELETE FROM cars WHERE id = ?", m.Id)
	return err
}

func GetWeatherByConfiguration(id int64) ([]Weather, error) {
	rows, err := db.Get().Query("SELECT * FROM weather WHERE configuration = ?", id)

	if err != nil {
		return nil, err
	}

	return scanWeather(rows)
}

func GetCarsByConfiguration(id int64) ([]Car, error) {
	rows, err := db.Get().Query("SELECT * FROM cars WHERE configuration = ? ORDER BY position ASC", id)

	if err != nil {
		return nil, err
	}

	return scanCars(rows)
}

func GetAllConfigurations() ([]Configuration, error) {
	rows, err := db.Get().Query("SELECT * FROM configurations ORDER BY name ASC")

	if err != nil {
		return nil, err
	}

	return scanConfigurations(rows)
}

func GetConfigurationById(id int64) (*Configuration, error) {
	row := db.Get().QueryRow("SELECT * FROM configurations WHERE id = ?", id)
	return scanConfiguration(row)
}

func scanConfigurations(rows *sql.Rows) ([]Configuration, error) {
	config := make([]Configuration, 0)

	for rows.Next() {
		c, err := scanConfiguration(rows)

		if err != nil {
			return nil, err
		}

		config = append(config, *c)
	}

	return config, nil
}

func scanConfiguration(row util.RowScanner) (*Configuration, error) {
	config := Configuration{}

	if err := row.Scan(&config.Id,
		&config.Name,
		&config.Pwd,
		&config.AdminPwd,
		&config.PickupMode,
		&config.LockEntryList,
		&config.RaceOvertime,
		&config.MaxSlots,
		&config.Welcome,
		&config.Description,
		&config.UDP,
		&config.TCP,
		&config.HTTP,
		&config.PacketsHz,
		&config.LoopMode,
		&config.ShowInLobby,
		&config.Threads,
		&config.ABS,
		&config.TC,
		&config.StabilityAid,
		&config.AutoClutch,
		&config.TyreBlankets,
		&config.ForceVirtualMirror,
		&config.FuelRate,
		&config.DamageRate,
		&config.TiresWearRate,
		&config.AllowedTiresOut,
		&config.MaxBallast,
		&config.DisableGasCutPenality,
		&config.ResultScreenTime,
		&config.DynamicTrack,
		&config.Condition,
		&config.StartValue,
		&config.Randomness,
		&config.TransferredGrip,
		&config.LapsToImproveGrip,
		&config.KickVoteQuorum,
		&config.SessionVoteQuorum,
		&config.VoteDuration,
		&config.Blacklist,
		&config.MaxCollisionsKm,
		&config.Booking,
		&config.BookingTime,
		&config.Practice,
		&config.PracticeTime,
		&config.CanJoinPractice,
		&config.Qualify,
		&config.QualifyTime,
		&config.CanJoinQualify,
		&config.Race,
		&config.RaceLaps,
		&config.RaceTime,
		&config.RaceWaitTime,
		&config.RaceExtraLap,
		&config.JoinType,
		&config.Time,
		&config.SunAngle,
		&config.Track,
		&config.TrackConfig,
		&config.LegalTyres,
		&config.UdpPluginPort,
		&config.UdpPluginAddr,
		&config.RacePitWindowStart,
		&config.RacePitWindowEnd,
		&config.ReversedGridRacePos); err != nil {
		return nil, err
	}

	return &config, nil
}

func scanWeather(rows *sql.Rows) ([]Weather, error) {
	weather := make([]Weather, 0)

	for rows.Next() {
		w, err := scanOneWeather(rows)

		if err != nil {
			return nil, err
		}

		weather = append(weather, *w)
	}

	return weather, nil
}

func scanOneWeather(row util.RowScanner) (*Weather, error) {
	weather := Weather{}

	if err := row.Scan(&weather.Id,
		&weather.Configuration,
		&weather.Weather,
		&weather.BaseAmbientTemp,
		&weather.RealisticRoadTemp,
		&weather.BaseRoadTemp,
		&weather.AmbientVariation,
		&weather.RoadVariation,
		&weather.WindBaseSpeedMin,
		&weather.WindBaseSpeedMax,
		&weather.WindBaseDirection,
		&weather.WindVariationDirection); err != nil {
		return nil, err
	}

	return &weather, nil
}

func scanCars(rows *sql.Rows) ([]Car, error) {
	cars := make([]Car, 0)

	for rows.Next() {
		car, err := scanCar(rows)

		if err != nil {
			return nil, err
		}

		cars = append(cars, *car)
	}

	return cars, nil
}

func scanCar(row util.RowScanner) (*Car, error) {
	car := Car{}
	var spectator string

	if err := row.Scan(&car.Id,
		&car.Configuration,
		&car.Car,
		&car.Painting,
		&spectator,
		&car.Driver,
		&car.Team,
		&car.GUID,
		&car.Position,
		&car.FixedSetup); err != nil {
		return nil, err
	}

	if spectator == "1" {
		car.Spectator = true
	}

	return &car, nil
}
