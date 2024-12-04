#!/bin/sh

# insert new blank import into main.go
sed -i 's/\t\/\/ SIP/\t_ \"hasmanytrees.com\/aoc-2024\/cmd\/'"$1"'\"\n	\/\/ SIP/' main.go

# copy the dayX template folder into the user specified day
cp -r ./cmd/dayX ./cmd/$1

# navigate to the new folder
cd ./cmd/$1

# rename the day file
mv dayX.go $1.go

# replace the template day name with the user specified day name
sed -i 's/dayX/'"$1"'/g' *
