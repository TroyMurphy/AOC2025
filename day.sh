#!/bin/bash
source .env
THIS_DIR="$(pwd)"

# Find all dayN folders, extract N, get max
NEXT_DAY=$(find . -maxdepth 1 -type d -name 'day[0-9]*' \
  | sed 's|./day||' \
  | grep -E '^[0-9]+$' \
  | sort -n \
  | tail -1)

if [[ -z "$NEXT_DAY" ]]; then
  NEXT_DAY=1
else
  NEXT_DAY=$((NEXT_DAY + 1))
fi

DAY_FOLDER="day${NEXT_DAY}"

# Run go run . NEXT_DAY
go run ./new $NEXT_DAY
echo "go run ./new $NEXT_DAY"

# Terminate if $DAY_FOLDER directory does not exist
if [ ! -d "$DAY_FOLDER" ]; then
  echo "Error: Directory $DAY_FOLDER does not exist."
  exit 1
fi

# Initialize go module, copy template
cd $DAY_FOLDER
touch sample.txt
cp ../template/*.go .
go mod init "${DAY_FOLDER}"

# replace the string "DAYFOLDER" on line 41 of main.go with day${NEXT_DAY}
sed -i "s/DAYFOLDER/${DAY_FOLDER}/" main.go

cd ..
go work use $DAY_FOLDER
echo "Initialized $DAY_FOLDER with template.go"



