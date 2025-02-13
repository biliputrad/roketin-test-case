package challenge_one

import (
	"fmt"
	"math"
	"net/http"
	"test-case-roketin/common/constants"
	challengeOneDto "test-case-roketin/common/dto/challenge-one-dto"
	responseMessage "test-case-roketin/common/response-message"
)

type ChallengeOneService interface {
	ConvertEarthTimeToRoketinTime(dto challengeOneDto.ChallengeOne) responseMessage.Response
}

type challengeOneService struct{}

func NewChallengeOneService() *challengeOneService {
	return &challengeOneService{}
}

func (service *challengeOneService) ConvertEarthTimeToRoketinTime(dto challengeOneDto.ChallengeOne) responseMessage.Response {
	/*
			on earth 1 day = 24 hours, 1 hour = 60 minutes, 1 minute = 60 seconds
			on Roketin Planet 1 day consists of 10 hours, 1 hour 100 minutes, 1 minute 100 seconds.
		    so 1 day on Roketin Planet = 10 * 100 * 100 = 10000 seconds on earth

	*/

	if dto.Hour > 24 || dto.Minute > 60 || dto.Second > 60 {
		return responseMessage.Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid time format",
			Data:       nil,
		}
	}

	// check if the time is 00:00:00
	if dto.Hour == 0 {
		dto.Hour = 24
	}

	// convert earth time to seconds
	totalEarthSeconds := dto.Hour*3600 + dto.Minute*60 + dto.Second

	// convert earth seconds to roketin seconds
	totalRoketinSeconds := float64(totalEarthSeconds) * (float64(10*100*100) / float64(24*60*60))

	// convert roketin seconds to roketin time
	roketinHours := int(totalRoketinSeconds) / (100 * 100)
	roketinMinutes := (int(totalRoketinSeconds) % (100 * 100)) / 100
	roketinSeconds := int(math.Round(totalRoketinSeconds)) % 100

	// change the time to 00:00:00 if the time is 24:00:00
	if dto.Hour == 24 {
		dto.Hour = 0
	}

	result := fmt.Sprintf("On Earth %02d:%02d:%02d, on Roketin Planet %02d:%02d:%02d", dto.Hour, dto.Minute, dto.Second, roketinHours, roketinMinutes, roketinSeconds)

	// %02d is used to format the output to have 2 digits with leading zero if the number is less than 10 and remove decimal point
	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}

}
