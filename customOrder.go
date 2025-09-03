package storeApi

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"os"
)

func PaymentEvent(sum int) error {
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeKey == "" {
		return fmt.Errorf("STRIPE_SECRET_KEY не задана")
	}
	stripe.Key = stripeKey

	amount := int64(sum)

	// Тестовая карта (payment method)
	paymentMethodID := "pm_card_visa"

	pi, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Currency:      stripe.String(string(stripe.CurrencyUSD)), // "usd"
		PaymentMethod: stripe.String(paymentMethodID),
		Confirm:       stripe.Bool(true),
	})

	if err != nil {
		return err
	}

	fmt.Println("✅ PaymentIntent создан:")
	fmt.Println("ID:", pi.ID)
	fmt.Println("Статус:", pi.Status)
	fmt.Println("Сумма:", pi.Amount)
	fmt.Println("Клиентский секрет:", pi.ClientSecret)

	return nil
}
