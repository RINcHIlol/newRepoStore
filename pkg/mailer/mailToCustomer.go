package mailer

import (
	"fmt"
	"gopkg.in/mail.v2"
	"os"
	"storeApi/models"
	"time"
)

func MailToCustomer(products []models.ProductCount, order models.Order, finalPrice float64) (*mail.Message, []*os.File, error) {
	cids := make([]string, len(products))

	tmpFiles := make([]*os.File, len(products))

	for i, product := range products {
		tmpFile, err := os.CreateTemp("", fmt.Sprintf("product-%d-*.jpg", i))
		if err != nil {
			return nil, nil, err
		}

		_, err = tmpFile.Write(product.Product.Image)
		if err != nil {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
			return nil, nil, err
		}

		tmpFiles[i] = tmpFile
		cids[i] = fmt.Sprintf("product-image-%d", i)
	}

	buyer := mail.NewMessage()
	buyer.SetHeader("From", "galimatron229@gmail.com")
	buyer.SetHeader("To", order.CustomerEmail)
	buyer.SetHeader("Subject", "Shopping in Svarka_Shop")

	for i, tmpFile := range tmpFiles {
		buyer.Embed(tmpFile.Name(), mail.SetHeader(map[string][]string{
			"Content-ID": {"<" + cids[i] + ">"},
		}))
	}

	productsHtml := ""
	for i, product := range products {
		productsHtml += fmt.Sprintf(`
		<table style="width: 100%%; margin-bottom: 20px; border: 1px solid #ddd; border-radius: 8px; padding: 10px;">
			<tr>
				<td style="width: 150px; padding: 10px; text-align: left; vertical-align: top;">
					<img src="cid:%s" alt="Product Image" style="max-width: 140px; border-radius: 8px; border: 2px solid #ccc; padding: 4px;"/>
				</td>
				<td style="padding: 10px; vertical-align: top; font-family: Arial, sans-serif;">
					<p><strong>Name:</strong> %s</p>
					<p><strong>Count:</strong> %d</p>
					<p><strong>Price:</strong> %.2f</p>
					<p><strong>Description:</strong> %s</p>
				</td>
			</tr>
		</table>
		`, cids[i], product.Product.Name, product.Count, product.Price, product.Product.Description)
	}

	body := fmt.Sprintf(`
	<div style="font-family: Arial, sans-serif; max-width: 600px; padding: 20px;">
		<div style="position: absolute; top: 10px; right: 20px; font-size: 14px; color: #888;">%s</div>
		<h2 style="color: #2196F3;">🎉 Thank You for Your Purchase!</h2>
		<p><strong>Order Number:</strong> %d</p>
		<p><strong>Your Email:</strong> %s</p>
		<p><strong>Shipping Address:</strong> %s</p>
		<p><strong>Finale Price:</strong> %0.2f</p>
		<hr style="margin: 20px 0;">
		%s
	</div>
	`, time.Now().Format("02 Jan 2006 15:04"), order.ID, order.CustomerEmail, order.Address, finalPrice, productsHtml)

	buyer.SetBody("text/html", body)

	return buyer, tmpFiles, nil
}

func MailToSeller(products []models.ProductCount, order models.Order, finalPrice float64) (*mail.Message, []*os.File, error) {
	cids := make([]string, len(products))
	tmpFiles := make([]*os.File, len(products))

	for i, product := range products {
		tmpFile, err := os.CreateTemp("", fmt.Sprintf("product-%d-*.jpg", i))
		if err != nil {
			return nil, nil, err
		}

		_, err = tmpFile.Write(product.Product.Image)
		if err != nil {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
			return nil, nil, err
		}

		tmpFiles[i] = tmpFile
		cids[i] = fmt.Sprintf("product-image-%d", i)
	}

	owner := mail.NewMessage()
	owner.SetHeader("From", "galimatron229@gmail.com")
	owner.SetHeader("To", order.CustomerEmail)
	owner.SetHeader("Subject", "Shopping in Svarka_Shop")

	for i, tmpFile := range tmpFiles {
		owner.Embed(tmpFile.Name(), mail.SetHeader(map[string][]string{
			"Content-ID": {"<" + cids[i] + ">"},
		}))
	}

	productsHtml := ""
	for i, product := range products {
		productsHtml += fmt.Sprintf(`
		<table style="width: 100%%; margin-bottom: 20px; border: 1px solid #ddd; border-radius: 8px; padding: 10px;">
			<tr>
				<td style="width: 150px; padding: 10px; text-align: left; vertical-align: top;">
					<img src="cid:%s" alt="Product Image" style="max-width: 140px; border-radius: 8px; border: 2px solid #ccc; padding: 4px;"/>
				</td>
				<td style="padding: 10px; vertical-align: top; font-family: Arial, sans-serif;">
					<p><strong>Name:</strong> %s</p>
					<p><strong>Count:</strong> %d</p>
					<p><strong>Price:</strong> %.2f</p>
					<p><strong>Description:</strong> %s</p>
				</td>
			</tr>
		</table>
		`, cids[i], product.Product.Name, product.Count, product.Price, product.Product.Description)
	}

	ownerMsg := fmt.Sprintf(`
	<div style="font-family: Arial, sans-serif; max-width: 600px; border: 1px solid #ddd; border-radius: 8px; padding: 20px; position: relative;">
		<div style="position: absolute; top: 10px; right: 20px; font-size: 14px; color: #888;">%s</div>
		<h2 style="color: #4CAF50;">🛒 New Order Received</h2>
		<p><strong>Order Number:</strong> %d</p>
		<p><strong>Customer Email:</strong> %s</p>
		<p><strong>Shipping Address:</strong> %s</p>
		<p><strong>Order Number:</strong> %0.2f</p>
		<hr style="margin: 20px 0;">
		%s
	</div>
	`, time.Now().Format("02 Jan 2006 15:04"), order.ID, order.CustomerEmail, order.Address, finalPrice, productsHtml)

	owner.SetBody("text/html", ownerMsg)

	return owner, tmpFiles, nil
}
