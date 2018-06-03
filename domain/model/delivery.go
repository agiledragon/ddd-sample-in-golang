package model

import (
    "github.com/agiledragon/ddd-sample-in-golang/domain/model/base"
)

type Delivery struct {
    base.ValueObject
    AfterDays uint
}

