export LONGEVITY_WAITTIME="20s"
export LONGEVITY_ITERATION_COUNT="2000"
export TF_ACC=1 
go test ./... -v -count=1 -run=TestAccLongevity -timeout 72h
# -run=TestAccLongevity -benchmem -bench=TestAccLongevity
date