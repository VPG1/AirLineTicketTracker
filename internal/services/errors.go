package services

import "errors"

var IncorrectSearchPhrase = errors.New("incorrect search phrase")
var UserNotRegistered = errors.New("user not registered")
var FlightAlreadyTracked = errors.New("flight already tracked")
