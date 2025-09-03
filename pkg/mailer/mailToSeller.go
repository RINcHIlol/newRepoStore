package mailer

//func MailToSeller(products []models.Product, order models.Order) (*mail.Message, error) {
//	cids := make([]string, len(products))
//	tmpFiles := make([]*os.File, len(products))
//
//	for i, product := range products {
//		tmpFile, err := os.CreateTemp("", fmt.Sprintf("product-%d-*.jpg", i))
//		if err != nil {
//			return nil, err
//		}
//		defer os.Remove(tmpFile.Name())
//		defer tmpFile.Close()
//
//		_, err = tmpFile.Write(product.Image)
//		if err != nil {
//			return nil, err
//		}
//
//		tmpFiles[i] = tmpFile
//		cids[i] = fmt.Sprintf("product-image-%d", i)
//	}
//
//	owner := mail.NewMessage()
//	owner.SetHeader("From", "galimatron229@gmail.com")
//	owner.SetHeader("To", "workemailsvarka123@gmail.com")
//	owner.SetHeader("Subject", "Shopping in Svarka_Shop")
//
//	for i, tmpFile := range tmpFiles {
//		owner.Embed(tmpFile.Name(), mail.SetHeader(map[string][]string{
//			"Content-ID": {"<" + cids[i] + ">"},
//		}))
//	}
//
//	productsHtml := ""
//	for i, product := range products {
//		productsHtml += fmt.Sprintf(`
//		<table style="width: 100%%; margin-bottom: 20px; border: 1px solid #ddd; border-radius: 8px; padding: 10px;">
//			<tr>
//				<td style="width: 150px; padding: 10px; text-align: left; vertical-align: top;">
//					<img src="cid:%s" alt="Product Image" style="max-width: 140px; border-radius: 8px; border: 2px solid #ccc; padding: 4px;"/>
//				</td>
//				<td style="padding: 10px; vertical-align: top; font-family: Arial, sans-serif;">
//					<p><strong>Name:</strong> %s</p>
//					<p><strong>Price:</strong> %.2f</p>
//					<p><strong>Description:</strong> %s</p>
//				</td>
//			</tr>
//		</table>
//		`, cids[i], product.Name, product.Price, product.Description)
//	}
//
//	ownerMsg := fmt.Sprintf(`
//	<div style="font-family: Arial, sans-serif; max-width: 600px; border: 1px solid #ddd; border-radius: 8px; padding: 20px; position: relative;">
//		<div style="position: absolute; top: 10px; right: 20px; font-size: 14px; color: #888;">%s</div>
//		<h2 style="color: #4CAF50;">🛒 New Order Received</h2>
//		<p><strong>Order Number:</strong> %d</p>
//		<p><strong>Customer Email:</strong> %s</p>
//		<p><strong>Shipping Address:</strong> %s</p>
//		<hr style="margin: 20px 0;">
//		%s
//	</div>
//	`, time.Now().Format("02 Jan 2006 15:04"), order.ID, order.CustomerEmail, order.Address, productsHtml)
//
//	owner.SetBody("text/html", ownerMsg)
//
//	//d := mail.NewDialer("smtp.gmail.com", 587, "galimatron229@gmail.com", "your_password_here")
//	//
//	//if err := d.DialAndSend(owner); err != nil {
//	//	return nil, err
//	//}
//
//	return owner, nil
//}
