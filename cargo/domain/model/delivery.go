package model

import (
    "github.com/agiledragon/ddd-sample-in-golang/cargo/domain/model/base"
)

type Delivery struct {
    base.ValueObject
    AfterDays uint
}

