glide:
	@sed -i.bk '/aximemory\|axiprotocol\|sdaccel/d' main.go
	@sed -i.bk '/xcl/d' cmd/test-mcmc/main.go
	@sed -i.bk '/xcl\|reco-sdaccel/d' cmd/bench-mcmc/main.go
	@mv main.go.bk ../
	@mv cmd/test-mcmc/main.go.bk ../cmd.go.bk
	@mv cmd/bench-mcmc/main.go.bk ../bench.go.bk
	@glide update
	@mv ../main.go.bk main.go
	@mv ../cmd.go.bk cmd/test-mcmc/main.go
	@mv ../bench.go.bk cmd/bench-mcmc/main.go

test:
	@reco project set mcmc
	reco test run test-mcmc

bench:
	@reco project set mcmc
	reco test run bench-mcmc
