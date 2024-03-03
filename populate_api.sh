#!/bin/bash

bb=$(tput bold)
nn=$(tput sgr0)
bold=$(tput bold)
green=$(tput setaf 2)
yellow=$(tput setaf 3)

function make_request() {
    local title="$1"
    local description="$2"

    local response=$(curl -X POST http://localhost:8080/seasons \
         -H "Content-Type: application/json" \
         -d '{"title": "'"$title"'", "description": "'"$description"'"}')
}

echo ""
echo -ne "${bold}"
echo "┌────────────────────────────┐"
echo "  Populate Go-Go-Power-Rangers"
echo "└────────────────────────────┘"
echo -ne "${nn}"
echo ""

make_request "Mighty Morphin" "The original Power Rangers series."
make_request "Zeo" "Power Rangers Zeo introduced a new team and storyline."
make_request "In Space" "The Space Rangers defend the universe against evil forces."
make_request "Dino Thunder" "Power Rangers with dinosaur-themed powers."
make_request "Lost Galaxy" "Power Rangers in outer space."
make_request "Time Force" "Time-traveling Power Rangers."
make_request "Ninja Storm" "Ninja-themed Power Rangers."
make_request "Jungle Fury" "Power Rangers with animal spirits."
make_request "Lightspeed Rescue" "Emergency-response-themed Power Rangers."
make_request "RPM" "Post-apocalyptic Power Rangers."
make_request "Dino Charge" "Power Rangers with dinosaur-themed powers and energems."
make_request "Beast Morphers" "Power Rangers with animal-themed powers and advanced technology."

echo ""
echo -n "${green}Seasons populated successfully!${norm}"
