module github.com/dev-yakuza/study-golang/module/hello

go 1.17

replace github.com/dev-yakuza/study-golang/module/greeting => ../greeting

require github.com/dev-yakuza/study-golang/module/greeting v0.0.0-00010101000000-000000000000
