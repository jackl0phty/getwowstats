# GETWOWSTATS
GOLang cmd-line utility that hits Blizzard's pubic API for WoW character stats.

# Install Prerequisites
A recent version of GOlang ( tested with version go1.13.8 linux/amd64 on Linux Mint )

Library to read cmd line options
```
go get github.com/ogier/pflag
```

# HARD REQUIREMENTS YOU MUST DO TO USE GETWOWSTATS
You MUST have a valid WoW token for Blizzard's API or you can't use this tool. No worries it's free.

# How/Where do I get a WoW API Token?
Great question and Blizz has your back. Read their doc [here](https://develop.battle.net/documentation/guides/getting-started).

# OK I Have a WoW API Token What next?
You MUST export ENV VAR WOW_API_TOKEN with the value of a valid token.
```
export WOW_API_TOKEN=your_secret_api_wow_token_goes_here
```

# Sample Usage & Output

## Get profile summary stats for a given character
```
./getwowstats.go -r aggramar -l us -c pownage -m GetProfileSummary
Name: Pownage
Gender: Male
Faction: Horde
Race: Orc
Class: Warlock
Realm: Aggramar
Level: 54
Active Spec: Affliction
Experience: 246258
Achieveent Points: 1310
Last Login: 1.622774087e+12
Avg. Item Lvl.: 94
Equipped Item Lvl.: 88
```
## Get version of getwowstats
```
./getwowstats -veresion
getwowstats: Version 1.0
```
