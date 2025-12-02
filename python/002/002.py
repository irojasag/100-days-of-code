print("What a nice dinner! Time to pay the bill.")
bill_amount = float(input("Enter the total bill amount: $"))
recommended_tip = bill_amount * 0.15
print(f"A recommended tip (15%) would be: ${recommended_tip:.2f}")
tip_amount = float(input("Enter the tip amount you would like to give: $"))
total_amount = bill_amount + tip_amount
print(f"The total amount to be paid is: ${total_amount:.2f}")