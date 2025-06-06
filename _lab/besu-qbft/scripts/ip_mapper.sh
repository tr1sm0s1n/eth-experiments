#!/bin/bash

ENV_FILE=".env.local"

# Update the IP address
chain_url="http://$(hostname -I | awk '{print $1}'):8545"

# Check if the .env file exists
if [ -f "$ENV_FILE" ]; then

  echo -e "\033[92m[INFO]: Reading $ENV_FILE file.\033[0m"

  # Read the .env file line by line
  while IFS= read -r line; do
    # Skip comments and empty lines
    if [[ "$line" =~ ^\s*#.*$ || -z "$line" ]]; then
      continue
    fi

    # Split the line to get key
    key=$(echo "$line" | cut -d '=' -f 1)

    # Replace CHAIN_URL with current one
    if [ "$key" == "CHAIN_URL" ]; then
        sed -i "s#$line#$key=$chain_url#" "$ENV_FILE"
        # Print the previous value
        echo -e "    \033[91m- : $line\033[0m"
        # Print the new value
        echo -e "    \033[92m+ : $key=$chain_url\033[0m"
        break
    fi

  done < "$ENV_FILE"
  echo -e "\033[92m[DONE]: Updated $ENV_FILE file.\033[0m"
else
  echo -e "\033[91m[ERROR]: $ENV_FILE not found.\033[0m"
fi
