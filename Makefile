.PHONY: build
build:
	sam build
	
.PHONY: run
run:
	sam local invoke PickupNewsFunction -n local/env.json -e local/event.json