package domain

import (
	"context"
	"errors"
)

type PaymentProvider string

const (
	ProviderStripe PaymentProvider = "stripe"
	ProviderCrypto PaymentProvider = "crypto"
)

type Currency string

const (
	CurrencyUSD Currency = "usd"
	CurrencyETH Currency = "eth"
)

var (
	ErrPriceMustBePositive              = errors.New("price must be positive")
	ErrMaxDeviceLimitMustBePositive     = errors.New("max device limit must be positive")
	ErrOldPriceMustBeGreaterThanCurrent = errors.New("old price must be greater than current")
	ErrTitleIsRequired                  = errors.New("title is required")
	ErrAtLeastOnePriceRequired          = errors.New("at least one price required")
	ErrInvalidPaymentProvider           = errors.New("invalid payment provider")
	ErrInvalidCurrency                  = errors.New("invalid currency")
	ErrIDIsRequired                     = errors.New("id is required")
	ErrDeviceLimitReached               = errors.New("device limit reached")
)

type Price struct {
	Value    int64
	OldValue int64
	Provider PaymentProvider
	Currency Currency
}

type Tariff struct {
	ID         string
	Title      string
	MaxDevices uint8
	Prices     []Price
	IsVisible  bool
}

type BillingRepository interface {
	CreateTariff(ctx context.Context, tariff Tariff) (newTariff Tariff, err error)
	UpdateTariff(ctx context.Context, tariff Tariff) (updatedTariff Tariff, err error)
	DeleteTariff(ctx context.Context, id string) (err error)

	GetByID(ctx context.Context, id string) (tariff Tariff, err error)
	GetVisible(ctx context.Context) (tariffs []Tariff, err error)
	GetAll(ctx context.Context) (tariffs []Tariff, err error)
}

type BillingService interface {
	CreateTariff(ctx context.Context, title string, prices []Price) (tariff Tariff, err error)

	HideTariff(ctx context.Context, id string) (err error)
	ShowTariff(ctx context.Context, id string) (err error)

	UpdateTariff(ctx context.Context, id string, title string, prices []Price) (tariff Tariff, err error)
}

func NewPrice(value, oldValue int64, provider PaymentProvider, currency Currency) (price Price, err error) {
	if value <= 0 {
		return Price{}, ErrPriceMustBePositive
	}

	if oldValue != 0 && oldValue <= value {
		return Price{}, ErrOldPriceMustBeGreaterThanCurrent
	}

	if !provider.IsValid() {
		return Price{}, ErrInvalidPaymentProvider
	}

	if !currency.IsValid() {
		return Price{}, ErrInvalidCurrency
	}

	return Price{
		Value:    value,
		OldValue: oldValue,
		Provider: provider,
		Currency: currency,
	}, nil
}

func (p Price) HasDiscount() (hasDiscount bool) {
	return p.OldValue > 0
}

func NewTariff(id string, title string, maxDevices uint8, prices []Price) (tariff Tariff, err error) {
	if id == "" {
		return Tariff{}, ErrIDIsRequired
	}

	if title == "" {
		return Tariff{}, ErrTitleIsRequired
	}

	if len(prices) == 0 {
		return Tariff{}, ErrAtLeastOnePriceRequired
	}

	if maxDevices <= 0 {
		return Tariff{}, ErrMaxDeviceLimitMustBePositive
	}

	return Tariff{
		ID:         id,
		Title:      title,
		MaxDevices: maxDevices,
		Prices:     prices,
		IsVisible:  true,
	}, nil
}

func (t *Tariff) Hide() {
	t.IsVisible = false
}

func (t *Tariff) Show() {
	t.IsVisible = true
}

func (t Tariff) CanAddDevice(current int) (err error) {
	if current <= int(t.MaxDevices) {
		return nil
	}

	return ErrDeviceLimitReached
}

func (p PaymentProvider) IsValid() (isValid bool) {
	switch p {
	case ProviderStripe, ProviderCrypto:
		return true
	}
	return false
}

func (c Currency) IsValid() (isValid bool) {
	switch c {
	case CurrencyUSD, CurrencyETH:
		return true
	}
	return false
}
