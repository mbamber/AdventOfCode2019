.PHONY: clean help solve

# Remove any built binaries
clean:
	@if [[ -f ./aoc ]]; then rm ./aoc; fi

# Print some help to the screen
help:
	@echo "Use 'make solve DAY=X PART=Y' to solve the puzzle for the given day and part"
	@echo "Optionally supply 'INPUT=path_to_input' to override the default path to the input file"

# Solve a puzzle for a day and part
solve: clean
	@go build
ifdef INPUT
	@./aoc -d $(DAY) -p $(PART) -i $(INPUT)
else
	@./aoc -d $(DAY) -p $(PART)
endif
