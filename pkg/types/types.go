package types

const (
	Meta_Extension  = "meta"
	Fitbit_Type     = "health.fitbit.type"
	ID_Type         = "health.id.type"
	Strava_Type     = "health.strava.type"
	Runkeeper_Type  = "health.runkeeper.type"
	Step_Type       = "health.step.type"
	Elliptical_Type = "health.elliptical.type"
	Walk_Type       = "health.walk.type"
	Run_Type        = "health.run.type"
	Ride_Type       = "health.ride.type"
)

type Steps struct {
	Steps int  `json:"steps,omitempty"`
	ID    Meta `json:"id,omitempty"`
}

type Activity struct {
	ActivityDuration int  `json:"activityDuration,omitempty"`
	ID               Meta `json:"id,omitempty"`
}

type Meta struct {
	Date string `json:"date,omitempty"`
	Name string `json:"name,omitempty"`
}

type Fitbit struct {
	Summary    FitbitSummary      `json:"summary,omitempty"`
	Activities []FitbitActivities `json:"activities,omitempty"`
}

type FitbitActivities struct {
	ActivityName string `json:"name,omitempty"`
	Duration     int    `json:"duration,omitempty"`
}

type FitbitSummary struct {
	Steps int `json:"steps,omitempty"`
}

type Strava struct {
	StartDate    string        `json:"start_date,omitempty"`
	Duration     int           `json:"elapsed_time,omitempty"`
	ActivityName string        `json:"type,omitempty"`
	Athlete      StravaAthlete `json:"athlete,omitempty"`
}

type StravaAthlete struct {
	ID int `json:"id,omitempty"`
}

type Runkeeper struct {
	Data string `json:"data,omitempty"`
}
