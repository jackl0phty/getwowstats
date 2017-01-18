# getwowstats
Ruby cmd-line utility that hits Blizzard's pubic API for WoW character stats.

# Install Prerequisites
If you use RVM you can do the following:
```
yoda@thedeathstar$ rvm install 2.4.0
yoda@thedeathstar$ rvm gemset create getwowstats
yoda@thedeathstar$ rvm gemset use getwowstats
yoda@thedeathstar$ ruby --version
ruby 2.4.0p0 (2016-12-24 revision 57164) [x86_64-darwin16]
yoda@thedeathstar$ gem install bundler
yoda@thedeathstar$ bundle
Fetching gem metadata from https://rubygems.org/.........
Fetching version metadata from https://rubygems.org/..
Fetching dependency metadata from https://rubygems.org/.
Resolving dependencies...
Installing multipart-post 2.0.0
Using bundler 1.13.7
Installing faraday 0.11.0
Installing faraday_middleware 0.11.0
Bundle complete! 2 Gemfile dependencies, 4 gems now installed.
Use `bundle show [gemname]` to see where a bundled gem is installed.
yoda@thedeathstar$
```

