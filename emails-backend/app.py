import json
import smtplib 
from email.message import EmailMessage
from kafka import KafkaConsumer

ORDER_CONFIRMED_TOPIC = "order_confirmed"

consumer = KafkaConsumer(
  ORDER_CONFIRMED_TOPIC,
  bootstrap_servers="host.docker.internal:9094",
)

print("Waiting for messages...")

for message in consumer:
  print("Received message: ", message.value)
  order = json.loads(message.value)

  msg = EmailMessage()

  msg["Subject"] = "Order confirmed"

  msg["From"] = "food@order.co"

  msg["To"] = order["client_email"]

  msg.set_content(f"Your order {order['document_id']} has been confirmed!")

  with smtplib.SMTP("smtp.mailtrap.io", 2525) as smtp:
    smtp.login("5ba372509783c9", "05263c6ca312f9")
    smtp.send_message(msg)