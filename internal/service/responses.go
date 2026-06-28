package service

import "github.com/TheDigitalMadness/ai-service-go/internal/domain"

type ValueField[T any] struct {
	Value T `json:"value"`
}

type MultivalueField[T any] struct {
	Value []T `json:"value"`
}

type Distance string

const (
	DistanceCity     Distance = "city"
	DistanceSuburban Distance = "suburban"
	DistanceLong     Distance = "long"
)

type Duration string

const (
	DurationOneDay   Duration = "one-day"
	DurationMultiDay Duration = "multi-day"
)

type Movements string

const (
	MovementsBus     Movements = "bus"
	MovementsWalking Movements = "walking"
	MovementsCar     Movements = "car"
	MovementsWater   Movements = "water"
)

type Events string

const (
	EventsMuseum   Events = "museum"
	EventsHoliday  Events = "holiday"
	EventsFestival Events = "festival"
	EventsQuest    Events = "quest"
	EventsConcert  Events = "concert"
)

type Direction string

const (
	DirectionMoscow                Direction = "Moscow"
	DirectionMoscowRegion          Direction = "Moscow region"
	DirectionSaintPetersburg       Direction = "Saint Petersburg"
	DirectionSaintPetersburgRegion Direction = "Saint Petersburg region"
)

type ResponseProperties struct {
	// Type of answer
	AnswerType ValueField[domain.AnswerType] `json:"answerType"`
	// Short answer for user
	ShortAnswer ValueField[string] `json:"shortAnswer"`

	// Distance type of a tour
	Distance *ValueField[Distance] `json:"distance"`
	// Duration of a tour
	Duration *ValueField[Duration] `json:"duration"`
	// Movements along a tour route. Optional
	Movements *MultivalueField[Movements] `json:"movements"`
	// Events during a tour
	Events *MultivalueField[Events] `json:"events"`
	// Direction of a tour
	Direction *ValueField[Direction] `json:"direction"`
	// If passport required or not
	PassportRequired *ValueField[bool] `json:"passportRequired"`
	// If partial payment available or not
	AvailablePartialPayment *ValueField[bool] `json:"availablePartialPayment"`
	// If route exists or not
	IsRouteExist *ValueField[bool] `json:"isRouteExist"`
	// Keywords which will be needed to find in description
	DescriptionKeyWords *MultivalueField[string] `json:"descriptionKeyWords"`
}

type ResponseFromAI struct {
	Properties ResponseProperties `json:"properties"`
}
