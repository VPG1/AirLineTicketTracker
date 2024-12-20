package postgres

import "errors"

var UserAlreadyExistsError = errors.New("User Already Exists")
var UserNotFoundError = errors.New("User Not Found")
var FlightAlreadyStored = errors.New("flight already stored")
