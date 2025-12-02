#!/bin/bash
echo "What a nice dinner! Time to pay the bill."
read -p "Enter the total bill amount: $" bill_amount
recommended_tip=$(echo "scale=2; $bill_amount * 0.15" | bc)
echo "A recommended tip (15%) would be: \$$recommended_tip"
read -p "Enter the tip amount you would like to give: $" tip_amount
total_amount=$(echo "scale=2; $bill_amount + $tip_amount" | bc)
echo "The total amount to be paid is: \$$total_amount"